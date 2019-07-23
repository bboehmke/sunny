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
const DiscoveryRequestPacketEntryTag uint16 = 0x0020

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
const DiscoveryIPPacketEntryTag uint16 = 0x0030

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
