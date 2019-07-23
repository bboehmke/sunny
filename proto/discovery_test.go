package proto

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDiscoveryRequest(t *testing.T) {
	ass := assert.New(t)

	packet := NewDiscoveryRequest()
	ass.Len(packet.entries, 2)

	ass.NotNil(packet.GetEntry(0x02a0))
	ass.NotNil(packet.GetEntry(0x0020))
}

func TestDiscoveryRequestPacketEntry_Tag(t *testing.T) {
	ass := assert.New(t)

	ass.Equal(uint16(0x0020), new(DiscoveryRequestPacketEntry).Tag())
	ass.Equal(uint16(0x0020), DiscoveryRequestPacketEntryTag)
}

func TestDiscoveryRequestPacketEntry_Bytes(t *testing.T) {
	ass := assert.New(t)

	ass.Empty(new(DiscoveryRequestPacketEntry).Bytes())
}

func TestDiscoveryRequestPacketEntry_Read(t *testing.T) {
	ass := assert.New(t)

	entry, err := new(DiscoveryRequestPacketEntry).Read(nil)
	ass.Nil(err)
	ass.NotNil(entry)
}

func TestDiscoveryIPPacketEntry_Tag(t *testing.T) {
	ass := assert.New(t)

	ass.Equal(uint16(0x0030), new(DiscoveryIPPacketEntry).Tag())
	ass.Equal(uint16(0x0030), DiscoveryIPPacketEntryTag)
}

func TestDiscoveryIPPacketEntry_Bytes(t *testing.T) {
	ass := assert.New(t)

	entry := DiscoveryIPPacketEntry{
		IP: []byte{0x01, 0x02, 0x03, 0x04},
	}
	ass.Equal([]byte{0x01, 0x02, 0x03, 0x04}, entry.Bytes())
}

func TestDiscoveryIPPacketEntry_Read(t *testing.T) {
	ass := assert.New(t)

	entry, err := new(DiscoveryIPPacketEntry).Read([]byte{0x01, 0x02, 0x03, 0x04})
	ass.Nil(err)
	ass.NotNil(entry)
	ass.Equal(net.IP([]byte{0x01, 0x02, 0x03, 0x04}), entry.(*DiscoveryIPPacketEntry).IP)
}
