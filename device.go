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
	}

	return nil, nil
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
