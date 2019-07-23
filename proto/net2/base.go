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

package net2

import (
	"encoding/binary"
	"fmt"
	"net"
)

// DeviceId provides information about a device
type DeviceId struct {
	SusyID       uint16
	SerialNumber uint32
}

// Bytes returns binary data
func (d *DeviceId) Bytes(byteOrder binary.ByteOrder) []byte {
	data := make([]byte, 6)
	byteOrder.PutUint16(data, d.SusyID)
	byteOrder.PutUint32(data[2:], d.SerialNumber)
	return data
}

// Read binary representation
func (d *DeviceId) Read(data []byte, byteOrder binary.ByteOrder) error {
	if len(data) < 6 {
		return fmt.Errorf("invalid DeviceId - length %d", len(data))
	}

	d.SusyID = byteOrder.Uint16(data)
	d.SerialNumber = byteOrder.Uint32(data[2:])
	return nil
}

// LocalDeviceId returns the device ID of the local system
// It is created from the first IPv4 address of an interface
func LocalDeviceId() *DeviceId {
	id := &DeviceId{
		SusyID:       120,
		SerialNumber: 123456789,
	}

	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		// fallback to static one
		return id
	}

	for _, inf := range interfaces {
		network, ok := inf.(*net.IPNet)

		// skip loopback and interfaces without IPv4
		if ok && !network.IP.IsLoopback() && network.IP.To4() != nil {
			return &DeviceId{
				SusyID:       120,
				SerialNumber: binary.BigEndian.Uint32(network.IP.To4()),
			}
		}
	}

	// fallback to static one
	return id
}
