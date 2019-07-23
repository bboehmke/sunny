package proto

import (
	"encoding/binary"
	"fmt"

	"gitlab.com/bboehmke/sunny/proto/net2"
)

// SmaNet2SubPacket inside net2 entry
type SmaNet2SubPacket interface {
	// ProtocolID identifies packet type
	ProtocolID() uint16
	// Bytes returns binary data
	Bytes() []byte
	// Read binary representation
	Read(data []byte) error
}

// SmaNet2PacketEntryTag identifier for net2 entries
const SmaNet2PacketEntryTag uint16 = 0x0010

// SmaNet2PacketEntry with a content packet
type SmaNet2PacketEntry struct {
	Content SmaNet2SubPacket
}

// Tag returns entry identifier
func (e *SmaNet2PacketEntry) Tag() uint16 {
	return SmaNet2PacketEntryTag
}

// Bytes returns binary data
func (e *SmaNet2PacketEntry) Bytes() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, e.Content.ProtocolID())
	return append(b, e.Content.Bytes()...)
}

// Read packet from the given binary data
func (e *SmaNet2PacketEntry) Read(data []byte) (PacketEntry, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("invalid SmaNet2PacketEntry - length %d", len(data))
	}

	protoID := binary.BigEndian.Uint16(data[0:2])

	var packet SmaNet2SubPacket
	var err error

	switch protoID {
	case net2.EnergyMeterPacketProtocolID:
		packet = new(net2.EnergyMeterPacket)
		err = packet.Read(data[2:])

	case net2.DeviceDataProtocolID:
		packet = new(net2.DeviceData)
		err = packet.Read(data[2:])
	}

	if err != nil {
		return nil, err
	}
	return &SmaNet2PacketEntry{
		Content: packet,
	}, nil
}
