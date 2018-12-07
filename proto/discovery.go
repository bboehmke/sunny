package proto

import (
	"net"
)

// NewDiscoveryRequest creates a new discovery request packet
func NewDiscoveryRequest() *Packet {
	var pack Packet
	pack.AddEntry(&GroupPacketEntry{
		Group: 0xFFFFFFFF,
	})
	pack.AddEntry(&DiscoveryRequestPacketEntry{})
	return &pack
}

// DiscoveryRequestPacketEntryTag identifier for discovery request entries
const DiscoveryRequestPacketEntryTag = 0x0020

// DiscoveryRequestPacketEntry empty packet
type DiscoveryRequestPacketEntry struct {
}

// Tag returns entry identifier
func (e DiscoveryRequestPacketEntry) Tag() uint16 {
	return DiscoveryRequestPacketEntryTag
}

// Bytes returns binary data
func (e *DiscoveryRequestPacketEntry) Bytes() []byte {
	return []byte{}
}

// Read packet from the given binary data
func (e *DiscoveryRequestPacketEntry) Read(data []byte) (PacketEntry, error) {
	return &DiscoveryRequestPacketEntry{}, nil
}

// DiscoveryIPPacketEntryTag identifier for discovery IP entries
const DiscoveryIPPacketEntryTag = 0x0030

// DiscoveryIPPacketEntry with IP address of device
type DiscoveryIPPacketEntry struct {
	IP net.IP
}

// Tag returns entry identifier
func (e DiscoveryIPPacketEntry) Tag() uint16 {
	return DiscoveryIPPacketEntryTag
}

// Bytes returns binary data
func (e *DiscoveryIPPacketEntry) Bytes() []byte {
	return e.IP
}

// Read packet from the given binary data
func (e *DiscoveryIPPacketEntry) Read(data []byte) (PacketEntry, error) {
	return &DiscoveryIPPacketEntry{
		IP: data,
	}, nil
}
