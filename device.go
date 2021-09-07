// Copyright 2019 Benjamin Böhmke <benjamin@boehmke.net>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sunny

import (
	"context"
	"fmt"
	"net"
	"time"

	"gitlab.com/bboehmke/sunny/proto"
	"gitlab.com/bboehmke/sunny/proto/net2"
)

// Device instance for communication with inverter and energy meter
type Device struct {
	// Address of inverter or energy meter
	address *net.UDPAddr
	// password for inverter communication
	password string

	// Connection instance for communication
	conn *Connection

	// device information
	energyMeter bool
	id          net2.DeviceId

	// receiver channel for received package with IP of this device
	receiver chan *proto.Packet
}

// NewDevice creates a new device instance
func (c *Connection) NewDevice(address, password string) (*Device, error) {
	device := Device{
		conn:     c,
		password: password,
		receiver: make(chan *proto.Packet, 2),
	}

	var err error
	device.address, err = net.ResolveUDPAddr("udp", address+":9522")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve udp address: %w", err)
	}
	// update address with resolved IP (in case of DNS)
	address = device.address.IP.String()

	// register receiver channel for this device
	c.registerReceiver(address, device.receiver)

	// send ping
	pingData := net2.NewDeviceData(0xa0)
	pingData.AddParameter(0)
	pingData.AddParameter(0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	for {
		// check for timeout
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("no ping response for %s", address)
		default:
		}

		err = device.sendDeviceData(pingData)
		if err != nil {
			Log.Printf("failed to send ping request for %s", address)
			return nil, err
		}

		// wait for receive
		receiveCtx, receiveCancel := context.WithTimeout(ctx, time.Millisecond*500)
		net2Entry, err := device.readNet2(receiveCtx)
		receiveCancel()
		if err != nil {
			continue
		}

		switch c := net2Entry.Content.(type) {
		case *net2.EnergyMeterPacket:
			Log.Printf("new energy meter at %s - Serial=%d", address, c.Id.SerialNumber)
			device.energyMeter = true
			device.id = c.Id
			return &device, nil

		case *net2.DeviceData:
			Log.Printf("new inverter at %s - Serial=%d", address, c.Source.SerialNumber)
			device.id = c.Source
			return &device, nil
		}
	}
}

// Close unregister receiver channel
func (d *Device) Close() {
	d.conn.unregisterReceiver(d.address.IP.String(), d.receiver)
}

// SetPassword for device communication
func (d *Device) SetPassword(pw string) {
	d.password = pw
}

// SerialNumber returns the serial number of the device
func (d *Device) SerialNumber() uint32 {
	return d.id.SerialNumber
}

// Address returns the address of the device
func (d *Device) Address() *net.UDPAddr {
	return d.address
}

// IsEnergyMeter returns true if devices is an energy meter
func (d *Device) IsEnergyMeter() bool {
	return d.energyMeter
}

// GetValue from inverter and returns nil if value does not exist
// Note: to request multiple values use GetValues
func (d *Device) GetValue(id ValueID) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return d.GetValueCtx(ctx, id)
}

// GetValueCtx from inverter and returns nil if value does not exist
// Note: to request multiple values use GetValues
func (d *Device) GetValueCtx(ctx context.Context, id ValueID) (interface{}, error) {
	if d.energyMeter {
		// handle some fixed energy meter values
		if id == DeviceClass {
			return 1, nil
		}
		if id == DeviceName {
			return "Energy Meter", nil
		}

		// no selective request for energy meter -> request all
		values, err := d.GetValuesCtx(ctx)
		if err != nil {
			return nil, err
		}
		return values[id], nil
	}

	// clear queue -> get fresh data
	d.clearReceiver()

	err := d.loginRetry(ctx, 3)
	if err != nil {
		return 0, err
	}

	values, err := d.requestValues(ctx, getInverterRequest(id))
	if err != nil {
		return 0, err
	}

	d.logout()
	return values[id], nil
}

// GetValues from device
func (d *Device) GetValues() (map[ValueID]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return d.GetValuesCtx(ctx)
}

// GetValuesCtx from device
func (d *Device) GetValuesCtx(ctx context.Context) (map[ValueID]interface{}, error) {
	// clear queue -> get fresh data
	d.clearReceiver()

	if d.energyMeter {
		for {
			// check for timeout
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("energy meter does not respond")
			default:
			}

			// wait for received packet
			net2Entry, err := d.readNet2(ctx)
			if err != nil {
				continue
			}

			packet, ok := net2Entry.Content.(*net2.EnergyMeterPacket)
			if !ok {
				continue
			}
			return convertEnergyMeterValues(packet.GetValues()), nil
		}
	}

	// login to device
	err := d.loginRetry(ctx, 3)
	if err != nil {
		return nil, err
	}

	// request all values and join to one map
	valuesMap := make(map[ValueID]interface{})
	for _, def := range getAllInverterRequests() {
		values, err := d.requestValues(ctx, def)
		if err != nil {
			Log.Printf("failed to get values for %s: %v", d.address, err)
			continue
		}
		if values == nil {
			continue
		}
		for id, value := range values {
			valuesMap[id] = value
		}
	}

	// logout
	d.logout()

	return valuesMap, nil
}

func (d *Device) loginRetry(ctx context.Context, trys int) (err error) {
	for i := 0; i < trys; i++ {
		if err = d.login(ctx); err == nil {
			return
		}
	}
	return
}

// login to device
func (d *Device) login(ctx context.Context) error {
	Log.Printf("login for %s", d.address)
	loginData := net2.NewDeviceData(0xa0)
	loginData.Command = 0x0c
	loginData.Object = 0xfffd
	loginData.JobNumber = 0x01

	loginData.AddParameter(7) // 10 for installer
	loginData.AddParameter(0x0384)
	loginData.AddParameter(uint32(time.Now().Unix()))
	loginData.AddParameter(0)

	// "encrypt" user password
	pass := []byte(d.password)
	encryptKey := byte(0x88) // 0xBB for installer

	passwordData := make([]byte, 12)
	for i := 0; i < 12; i++ {
		if i < len(pass) {
			passwordData[i] = pass[i] + encryptKey
		} else {
			passwordData[i] = encryptKey
		}
	}
	loginData.Data = passwordData

	response, err := d.sendDeviceDataResponse(loginData, time.Millisecond*500, ctx)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	if response.Status != 0 {
		return fmt.Errorf("login failed")
	}
	return nil
}

// logout to device
func (d *Device) logout() {
	Log.Printf("logout for %s", d.address)
	request := net2.NewDeviceData(0xa0)
	request.Command = 0x0e
	request.Object = 0xfffd
	request.JobNumber = 0x03

	request.AddParameter(0xFFFFFFFF)

	_ = d.sendDeviceData(request)
}

// requestValues from given definition
func (d *Device) requestValues(ctx context.Context, def InverterValuesDef) (map[ValueID]interface{}, error) {
	Log.Printf("requestValues for %s: 0x%X 0x%X 0x%X", d.address, def.Object, def.Start, def.End)
	request := net2.NewDeviceData(0xa0)
	request.Object = def.Object
	request.AddParameter(def.Start)
	request.AddParameter(def.End)

	response, err := d.sendDeviceDataResponse(request, time.Millisecond*500, ctx)
	if err != nil {
		return nil, err
	}

	if response.Status == 0x15 {
		return nil, nil
	}
	if response.Status != 0 {
		return nil, fmt.Errorf("failed to get values")
	}

	return parseInverterValues(response.ResponseValues), nil
}

// sendDeviceDataResponse sends the package and wait for response
func (d *Device) sendDeviceDataResponse(data *net2.DeviceData,
	resendInterval time.Duration, ctx context.Context) (*net2.DeviceData, error) {
	for {
		// stop after timeout
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("no packet received in timeout")
		default:
		}

		// send request
		err := d.sendDeviceData(data)
		if err != nil {
			return nil, err
		}

		// wait for response
		receiveCtx, cancel := context.WithTimeout(ctx, resendInterval)

		// wait for package until timeout
		for {
			select {
			case <-receiveCtx.Done():
				cancel()
				return nil, fmt.Errorf("no packet received in timeout")
			default:
			}

			responseData, err := d.readNet2DeviceData(receiveCtx, data.PacketID)
			if err != nil {
				continue // no valid packet
			}
			cancel()
			return responseData, nil
		}
	}
}

// sendDeviceData sends the package
func (d *Device) sendDeviceData(data *net2.DeviceData) error {
	if d.id.SusyID == 0 && d.id.SerialNumber == 0 {
		data.Destination.SusyID = 0xFFFF
		data.Destination.SerialNumber = 0xFFFFFFFF
	} else {
		data.Destination = d.id
	}

	var pack proto.Packet
	pack.AddEntry(&proto.GroupPacketEntry{
		Group: 0x00000001,
	})
	pack.AddEntry(&proto.SmaNet2PacketEntry{
		Content: data,
	})

	return d.conn.sendPacket(d.address, &pack)
}

// readNet2 read package from Connection
func (d *Device) readNet2(ctx context.Context) (*proto.SmaNet2PacketEntry, error) {
	var packet *proto.Packet
	select {
	case packet = <-d.receiver:
	case <-ctx.Done():
		return nil, fmt.Errorf("device does not respond at %s", d.address.IP.String())
	}

	entry := packet.GetEntry(proto.SmaNet2PacketEntryTag)
	if entry == nil {
		return nil, fmt.Errorf("received invalid response from %s", d.address.IP.String())
	}

	return entry.(*proto.SmaNet2PacketEntry), nil
}

// readNet2DeviceData read package from Connection
func (d *Device) readNet2DeviceData(ctx context.Context, pkgId uint16) (*net2.DeviceData, error) {
	entry, err := d.readNet2(ctx)
	if err != nil {
		return nil, err
	}

	responseData, ok := entry.Content.(*net2.DeviceData)
	if !ok {
		return nil, fmt.Errorf("invalid data received from %s", d.address.IP.String())
	}

	if responseData.PacketID != pkgId {
		return nil, fmt.Errorf("invalid package received from %s (%d)", d.address.IP.String(), responseData.PacketID)
	}

	return responseData, nil
}

// clearReceiver channel packages
func (d *Device) clearReceiver() {
	for {
		select {
		case <-d.receiver:
		default:
			return
		}
	}
}
