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
