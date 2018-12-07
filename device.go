package sunny

import (
	"fmt"
	"net"
	"time"

	"gitlab.com/bboehmke/sunny/proto"
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
	energyMeter  bool
	susyID       uint16
	serialNumber uint32
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
	pingData := proto.NewDeviceData(0xa0)
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

		packet, ok := net2Entry.Content.(*proto.EnergyMeterPacket)
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

// logout to device
func (d *Device) logout() {
	request := proto.NewDeviceData(0xa0)
	request.Command = 0x0e
	request.Object = 0xfffd
	request.JobNumber = 0x03

	request.AddParameter(0xFFFFFFFF)

	_ = d.sendDeviceData(request)
}

// requestValues from given definition
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

// sendDeviceDataResponse sends the package and wait for response
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

// sendDeviceData sends the package
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
