package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEntry struct{}

func (e *TestEntry) Tag() uint16 {
	return 0x1234
}
func (e *TestEntry) Bytes() []byte {
	return []byte{
		0x12, 0x34, 0x56, 0x78,
		0x12, 0x34, 0x56, 0x78,
	}
}
func (e *TestEntry) Read(data []byte) (PacketEntry, error) {
	return new(TestEntry), nil
}

func TestPacket_String(t *testing.T) {
	ass := assert.New(t)

	packet := Packet{
		entries: []PacketEntry{
			new(TestEntry),
			new(TestEntry),
		},
	}

	ass.Equal("proto.TestEntry, proto.TestEntry", packet.String())
}

func TestPacket_GetEntry(t *testing.T) {
	ass := assert.New(t)

	entry := new(TestEntry)
	packet := Packet{
		entries: []PacketEntry{
			entry,
		},
	}

	ass.Nil(packet.GetEntry(0x12))
	ass.Equal(entry, packet.GetEntry(0x1234))
}

func TestPacket_AddEntry(t *testing.T) {
	ass := assert.New(t)

	entry := new(TestEntry)
	packet := new(Packet)
	packet.AddEntry(entry)

	ass.Equal([]PacketEntry{entry}, packet.entries)
}

func TestPacket_Bytes(t *testing.T) {
	ass := assert.New(t)

	entry := new(TestEntry)
	packet := Packet{
		entries: []PacketEntry{
			entry,
		},
	}

	ass.Equal([]byte{
		0x53, 0x4d, 0x41, 0x00, // header
		0x00, 0x08, // packet data length
		0x12, 0x34, // packet id
		0x12, 0x34, 0x56, 0x78, // packet data
		0x12, 0x34, 0x56, 0x78, // packet data
		0x00, 0x00, 0x00, 0x00, // packet end
	}, packet.Bytes())
}

func TestPacket_Read(t *testing.T) {
	ass := assert.New(t)

	packet := new(Packet)

	ass.EqualError(packet.Read([]byte{0x12, 0x34}),
		"invalid packet - to small: 2")

	ass.EqualError(packet.Read([]byte{
		0x52, 0x4d, 0x41, 0x00, // header
		0x00, 0x08, // packet data length
		0x12, 0x34, // packet id
		0x12, 0x34, 0x56, 0x78, // packet data
		0x12, 0x34, 0x56, 0x78, // packet data
		0x00, 0x00, 0x00, 0x00, // packet end
	}), "invalid packet - header: RMA")

	ass.NoError(packet.Read([]byte{
		0x53, 0x4d, 0x41, 0x00, // header

		0x00, 0x08, // packet data length
		0x12, 0x34, // packet id
		0x12, 0x34, 0x56, 0x78, // packet data
		0x12, 0x34, 0x56, 0x78, // packet data

		0x00, 0x04, // packet data length
		0x02, 0xa0, // packet id (GroupPacketEntry)
		0x12, 0x34, 0x56, 0x78, // packet data
		0x00, 0x00, 0x00, 0x00, // packet end
	}))

	ass.Len(packet.entries, 2)
	ass.NotNil(packet.GetEntry(0x1234))
	ass.NotNil(packet.GetEntry(0x02a0))
}

func TestGroupPacketEntry_Tag(t *testing.T) {
	ass := assert.New(t)

	ass.Equal(uint16(0x02A0), new(GroupPacketEntry).Tag())
	ass.Equal(uint16(0x02A0), GroupPacketEntryTag)
}

func TestGroupPacketEntry_Bytes(t *testing.T) {
	ass := assert.New(t)

	entry := GroupPacketEntry{
		Group: 0x12345678,
	}

	ass.Equal([]byte{
		0x12, 0x34, 0x56, 0x78,
	}, entry.Bytes())
}

func TestGroupPacketEntry_Read(t *testing.T) {
	ass := assert.New(t)

	loadEntry := new(GroupPacketEntry)
	_, err := loadEntry.Read([]byte{0x12, 0x34})
	ass.EqualError(err, "invalid GroupPacketEntry - length 2")

	entry, err := loadEntry.Read([]byte{0x12, 0x34, 0x56, 0x78})
	ass.NoError(err)
	ass.Equal(&GroupPacketEntry{
		Group: 0x12345678,
	}, entry)
}

func TestUnknownPacketEntry(t *testing.T) {
	ass := assert.New(t)

	entry := UnknownPacketEntry{
		Data: []byte{0x12, 0x34, 0x56, 0x78},
		T:    0x1234,
	}
	ass.Equal(uint16(0x1234), entry.Tag())
	ass.Equal([]byte{0x12, 0x34, 0x56, 0x78}, entry.Bytes())

	e, err := entry.Read([]byte{0x12, 0x34})
	ass.Nil(e)
	ass.Nil(err)
}
