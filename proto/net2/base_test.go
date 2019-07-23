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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceId_Write(t *testing.T) {
	ass := assert.New(t)

	id := DeviceId{
		SusyID:       0x1234,
		SerialNumber: 0x12345678,
	}

	ass.Equal([]byte{
		0x12, 0x34,
		0x12, 0x34, 0x56, 0x78,
	}, id.Bytes(binary.BigEndian))

	ass.Equal([]byte{
		0x34, 0x12,
		0x78, 0x56, 0x34, 0x12,
	}, id.Bytes(binary.LittleEndian))
}

func TestDeviceId_Read(t *testing.T) {
	ass := assert.New(t)

	id := new(DeviceId)

	ass.EqualError(id.Read([]byte{
		0x12, 0x34,
	}, nil), "invalid DeviceId - length 2")

	ass.NoError(id.Read([]byte{
		0x12, 0x34,
		0x12, 0x34, 0x56, 0x78,
	}, binary.BigEndian))
	ass.Equal(uint16(0x1234), id.SusyID)
	ass.Equal(uint32(0x12345678), id.SerialNumber)

	ass.NoError(id.Read([]byte{
		0x12, 0x34,
		0x12, 0x34, 0x56, 0x78,
	}, binary.LittleEndian))
	ass.Equal(uint16(0x3412), id.SusyID)
	ass.Equal(uint32(0x78563412), id.SerialNumber)
}

func TestLocalDeviceId(t *testing.T) {
	ass := assert.New(t)

	id := LocalDeviceId()
	ass.Equal(uint16(120), id.SusyID)

	// we only test if the value are set
	ass.True(id.SerialNumber != 0)
}
