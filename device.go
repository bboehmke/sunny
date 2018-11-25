package sunny

import (
	"fmt"
	"net"
	"time"

	"gitlab.com/bboehmke/sunny/proto"
)

type Device struct {
	address  *net.UDPAddr
	password string

	conn *connection

	// device information
	energyMeter  bool
	susyID       uint16
	serialNumber uint32
}

func NewDevice(address, password string) (*Device, error) {
	device := Device{
		password: password,
	}

	var err error
	device.address, err = net.ResolveUDPAddr("udp", address+":9522")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve udp address: %v", err)
	}

	device.conn, err = getConnection(device.address)
	if err != nil {
		return nil, err
	}

	conn.clearReceived(device.address)

	pingData := proto.NewDeviceData(0xa0)
	pingData.AddParameter(0)
	pingData.AddParameter(0)
	err = device.sendDeviceData(pingData)
	if err != nil {
		return nil, err
	}

	net2Entry, err := device.readNet2(time.Second)
	if err != nil {
		return nil, err
	}

	switch c := net2Entry.Content.(type) {
	case *proto.EnergyMeterPacket:
		device.energyMeter = true
		device.susyID = c.SusyID
		device.serialNumber = c.SerNo

	case *proto.DeviceData:
		device.susyID = c.SrcSusyID
		device.serialNumber = c.SrcSerialNumber

	default:
		return nil, fmt.Errorf("received unknown net2 packet from %s", address)

	}

	return &device, nil
}

func (d *Device) GetValues() (map[string]interface{}, error) {
	// clear queue -> get fresh data
	d.conn.clearReceived(d.address)

	if d.energyMeter {
		net2Entry, err := d.readNet2(time.Second * 2)
		if err != nil {
			return nil, err
		}

		packet, ok := net2Entry.Content.(*proto.EnergyMeterPacket)
		if !ok {
			// TODO retry !!!
			return nil, fmt.Errorf("invalid packet received")
		}

		return packet.GetValues(), nil
	} else {
		err := d.login()
		if err != nil {
			return nil, err
		}

		values := make(map[string]interface{})
		for _, def := range getAllRequests() {
			vals, err := d.requestValues(def)
			if err != nil {
				return nil, err
			}
			if vals == nil {
				continue
			}
			for key, value := range vals {
				values[key] = value
			}
		}

		_ = d.logout()

		return values, nil
	}
}

func (d *Device) login() error {
	loginData := proto.NewDeviceData(0xa0)
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
func (d *Device) logout() error {
	request := proto.NewDeviceData(0xa0)
	request.Command = 0x0e
	request.Object = 0xfffd
	request.JobNumber = 0x03

	request.AddParameter(0xFFFFFFFF)

	return d.sendDeviceData(request)
}

func (d *Device) requestValues(def valDef) (map[string]interface{}, error) {
	request := proto.NewDeviceData(0xa0)
	request.Object = def.Object
	request.AddParameter(def.Start)
	request.AddParameter(def.End)

	response, err := d.sendDeviceDataResponse(request, time.Second)
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

func (d *Device) sendDeviceDataResponse(data *proto.DeviceData,
	timeout time.Duration) (*proto.DeviceData, error) {

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

		responseData, ok := entry.Content.(*proto.DeviceData)
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

func (d *Device) sendDeviceData(data *proto.DeviceData) error {
	if d.susyID == 0 && d.serialNumber == 0 {
		data.DstSusyID = 0xFFFF
		data.DstSerialNumber = 0xFFFFFFFF
	} else {
		data.DstSusyID = d.susyID
		data.DstSerialNumber = d.serialNumber
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
