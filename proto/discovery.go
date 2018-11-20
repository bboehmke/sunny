package proto

import (
	"net"

	"github.com/sirupsen/logrus"
)

func NewDiscoveryRequest() *Packet {
	var pack Packet
	pack.AddEntry(&GroupPacketEntry{
		Group: 0xFFFFFFFF,
	})
	pack.AddEntry(&DiscoveryRequestPacketEntry{})
	return &pack
}

const DiscoveryRequestPacketEntryTag = 0x0020

type DiscoveryRequestPacketEntry struct {
}

func (e DiscoveryRequestPacketEntry) Tag() uint16 {
	return DiscoveryRequestPacketEntryTag
}
func (e *DiscoveryRequestPacketEntry) Bytes() []byte {
	return []byte{}
}
func (e *DiscoveryRequestPacketEntry) Read(data []byte) (PacketEntry, error) {
	return &DiscoveryRequestPacketEntry{}, nil
}

const DiscoveryIpPacketEntryTag = 0x0030

type DiscoveryIpPacketEntry struct {
	IP net.IP
}

func (e DiscoveryIpPacketEntry) Tag() uint16 {
	return DiscoveryIpPacketEntryTag
}
func (e *DiscoveryIpPacketEntry) Bytes() []byte {
	return e.IP
}
func (e *DiscoveryIpPacketEntry) Read(data []byte) (PacketEntry, error) {
	logrus.Debug(data)
	return &DiscoveryIpPacketEntry{
		IP: data,
	}, nil
}
