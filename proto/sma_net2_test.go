package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/bboehmke/sunny/proto/net2"
)

func TestSmaNet2PacketEntry_Tag(t *testing.T) {
	ass := assert.New(t)

	ass.Equal(uint16(0x0010), new(SmaNet2PacketEntry).Tag())
	ass.Equal(uint16(0x0010), SmaNet2PacketEntryTag)
}

func TestSmaNet2PacketEntry_Bytes(t *testing.T) {
	ass := assert.New(t)

	entry := SmaNet2PacketEntry{
		Content: new(net2.EnergyMeterPacket),
	}
	data := entry.Bytes()
	ass.Equal([]byte{0x60, 0x69}, data[0:2])
}

func TestSmaNet2PacketEntry_Read(t *testing.T) {
	ass := assert.New(t)

	loadEntry := new(SmaNet2PacketEntry)

	_, err := loadEntry.Read([]byte{0x60, 0x69})
	ass.EqualError(err, "invalid SmaNet2PacketEntry - length 2")

	entry, err := loadEntry.Read([]byte{0x12, 0x34, 0x56, 0x78})
	ass.NoError(err)
	ass.NotNil(entry)
	ass.Nil(entry.(*SmaNet2PacketEntry).Content)

	// TODO check real sub entries
}
