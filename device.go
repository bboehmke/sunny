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

	// connection instance for communication
	conn *connection

	// device information
	energyMeter bool
	id          net2.DeviceId
}

// NewDevice creates a new device instance
func NewDevice(address, password string) (*Device, error) {
	device := Device{
		password: password,
	}

	var err error
	device.address, err = net.ResolveUDPAddr("udp", address+":9522")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve udp address: %v", err)
	}

	// get connection instance
	device.conn, err = getConnection()
	if err != nil {
		return nil, err
	}

	// clear message buffer
	conn.clearReceived(device.address)

	// send ping
	pingData := net2.NewDeviceData(0xa0)
	pingData.AddParameter(0)
	pingData.AddParameter(0)
	err = device.sendDeviceData(pingData)
	if err != nil {
		return nil, err
	}

	// wait for receive
	net2Entry, err := device.readNet2(time.Second)
	if err != nil {
		return nil, err
	}

	switch c := net2Entry.Content.(type) {
	case *net2.EnergyMeterPacket:
		device.energyMeter = true
		device.id = c.Id

	case *net2.DeviceData:
		device.id = c.Source

	default:
		return nil, fmt.Errorf("received unknown net2 packet from %s", address)
	}

	return &device, nil
}

// SerialNumber returns the serial number of the device
func (d *Device) SerialNumber() uint32 {
	return d.id.SerialNumber
}

// GetDeviceClass returns the class identifier of the device (1 = energy meter)
func (d *Device) GetDeviceClass() (uint32, error) {
	if d.energyMeter {
		return 1, nil
	}

	err := d.login()
	if err != nil {
		return 0, err
	}

	values, err := d.requestValues(getRequest("device_class"))
	if err != nil {
		return 0, err
	}

	// clear queue -> get fresh data
	d.conn.clearReceived(d.address)

	d.logout()
	if val, ok := values["device_class"]; ok {
		return val.(uint32), nil
	}
	return 0, fmt.Errorf("")
}

// GetValues from device
func (d *Device) GetValues() (map[string]interface{}, error) {
	// clear queue -> get fresh data
	d.conn.clearReceived(d.address)

	if d.energyMeter {
		net2Entry, err := d.readNet2(time.Second * 2)
		if err != nil {
			return nil, err
		}

		packet, ok := net2Entry.Content.(*net2.EnergyMeterPacket)
		if !ok {
			// TODO retry !!!
			return nil, fmt.Errorf("invalid packet received")
		}

		return packet.GetValues(), nil
	}

	// login to device
	err := d.login()
	if err != nil {
		return nil, err
	}

	// request all values and join to one map
	valuesMap := make(map[string]interface{})
	for _, def := range getAllRequests() {
		values, err := d.requestValues(def)
		if err != nil {
			return nil, err
		}
		if values == nil {
			continue
		}
		for key, value := range values {
			valuesMap[key] = value
		}
	}

	// logout
	d.logout()

	return valuesMap, nil
}

// login to device
func (d *Device) login() error {
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

	response, err := d.sendDeviceDataResponse(loginData, time.Second)
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	if response.Status != 0 {
		return fmt.Errorf("login failed")
	}
	return nil
}

// logout to device
func (d *Device) logout() {
	request := net2.NewDeviceData(0xa0)
	request.Command = 0x0e
	request.Object = 0xfffd
	request.JobNumber = 0x03

	request.AddParameter(0xFFFFFFFF)

	_ = d.sendDeviceData(request)
}

// requestValues from given definition
func (d *Device) requestValues(def valDef) (map[string]interface{}, error) {
	request := net2.NewDeviceData(0xa0)
	request.Object = def.Object
	request.AddParameter(def.Start)
	request.AddParameter(def.End)

	response, err := d.sendDeviceDataResponse(request, time.Second*10)
	if err != nil {
		return nil, err
	}

	if response.Status == 0x15 {
		return nil, nil
	}
	if response.Status != 0 {
		return nil, fmt.Errorf("failed to get values")
	}

	return parseValues(response.ResponseValues), nil
}

// sendDeviceDataResponse sends the package and wait for response
func (d *Device) sendDeviceDataResponse(data *net2.DeviceData,
	timeout time.Duration) (*net2.DeviceData, error) {

	err := d.sendDeviceData(data)
	if err != nil {
		return nil, err
	}

	startTime := time.Now()
	for time.Since(startTime) < timeout {
		entry, err := d.readNet2(timeout)
		if err != nil {
			continue // no valid packet
		}

		responseData, ok := entry.Content.(*net2.DeviceData)
		if !ok {
			continue // no device data packet
		}

		if responseData.PacketID != data.PacketID {
			continue // no response to request
		}

		return responseData, nil
	}

	return nil, fmt.Errorf("no packet received in timeout")
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

	return conn.sendPacket(d.address, &pack)
}

// readNet2 read package from connection
func (d *Device) readNet2(timeout time.Duration) (*proto.SmaNet2PacketEntry, error) {
	packet := d.conn.readPacket(d.address, timeout)
	if packet == nil {
		return nil, fmt.Errorf("device does not respond at %s", d.address.IP.String())
	}

	entry := packet.GetEntry(proto.SmaNet2PacketEntryTag)
	if entry == nil {
		return nil, fmt.Errorf("received invalid response from %s", d.address.IP.String())
	}

	return entry.(*proto.SmaNet2PacketEntry), nil
}
