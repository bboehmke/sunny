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
	"strings"
	"sync/atomic"
)

// DeviceDataProtocolID protocol ID used for DeviceData sub packets
const DeviceDataProtocolID uint16 = 0x6065

// ResponseValue of device data packet response
type ResponseValue struct {
	Class     uint8
	Code      uint16
	Type      uint8
	Timestamp uint32

	Values []interface{}
}

// Bytes returns binary data
func (v *ResponseValue) Bytes(object uint16) []byte {
	var data []byte

	// string value
	if v.Type == 0x10 {
		data = make([]byte, 40)

		// attributes
	} else if v.Type == 0x08 {
		data = make([]byte, 40)

		// large integer value
	} else if object == 0x5400 {
		data = make([]byte, 16)

		// signed and unsigned 32 bit integer value
	} else if v.Type == 0x00 || v.Type == 0x40 {
		data = make([]byte, 8+5*4)

		// unknown
	} else {
		data = make([]byte, 8)
	}

	data[0] = v.Class
	binary.LittleEndian.PutUint16(data[1:], v.Code)
	data[3] = v.Type
	binary.LittleEndian.PutUint32(data[4:], v.Timestamp)

	if len(v.Values) == 0 {
		return data
	}

	// string value
	if v.Type == 0x10 {
		s, ok := v.Values[0].(string)
		if ok {
			copy(data[8:], []byte(s))
		}

		// attributes
	} else if v.Type == 0x08 {
		index := 8
		for _, val := range v.Values {
			v, ok := val.(uint32)
			if !ok {
				break
			}

			binary.LittleEndian.PutUint32(data[index:], (v&0xffffff)|0x01000000)
			index += 4
		}
		if len(v.Values) < 8 {
			binary.LittleEndian.PutUint32(data[index:], 0xfffffe)
		}

		// large integer value
	} else if object == 0x5400 {
		val, ok := v.Values[0].(uint64)
		if ok {
			binary.LittleEndian.PutUint64(data[8:], val)
		}

		// unsigned 32 bit integer value
	} else if v.Type == 0x00 {
		index := 8
		for _, val := range v.Values {
			val, ok := val.(uint32)
			if !ok {
				break
			}
			binary.LittleEndian.PutUint32(data[index:], val)
			index += 4
		}
		if len(v.Values) < 5 {
			binary.LittleEndian.PutUint32(data[index:], 0xffffffff)
		}

		// signed 32 bit integer value
	} else if v.Type == 0x40 {
		index := 8
		for _, val := range v.Values {
			val, ok := val.(int32)
			if !ok {
				break
			}
			binary.LittleEndian.PutUint32(data[index:], uint32(val))
			index += 4
		}
		if len(v.Values) < 5 {
			v := -0x80000000
			binary.LittleEndian.PutUint32(data[index:], uint32(v))
		}
	}

	return data
}

// Read binary representation
func (v *ResponseValue) Read(data []byte, object uint16) (int, error) {
	if len(data) < 8 {
		return 0, fmt.Errorf("invalid ResponseValue - length %d", len(data))
	}

	v.Class = data[0]
	v.Code = binary.LittleEndian.Uint16(data[1:])
	v.Type = data[3]
	v.Timestamp = binary.LittleEndian.Uint32(data[4:])

	// string value
	if v.Type == 0x10 {
		v.Values = []interface{}{
			strings.Trim(string(data[8:40]), "\x00"),
		}
		return 40, nil

		// attributes
	} else if v.Type == 0x08 {
		dataLength := len(data)
		index := 8
		v.Values = make([]interface{}, 0, 8)
		for i := 0; i < 8; i++ {
			if dataLength-index < 4 {
				break
			}

			val := binary.LittleEndian.Uint32(data[index:])
			index += 4

			if val == 0xfffffe {
				break
			}
			if val>>24 == 1 {
				v.Values = append(v.Values, val&0xffffff)
			}
		}
		return 40, nil

		// large integer value
	} else if object == 0x5400 {
		if len(data) < 16 {
			return 8, fmt.Errorf("invalid ResponseValue - length %d", len(data))
		}
		v.Values = []interface{}{
			binary.LittleEndian.Uint64(data[8:16]),
		}
		return 16, nil

		// signed and unsigned 32 bit integer value
	} else if v.Type == 0x00 || v.Type == 0x40 {
		dataLength := len(data)
		index := 8

		v.Values = make([]interface{}, 0, 5)

		for i := 0; i < 5; i++ {
			if dataLength-index < 4 {
				break
			}

			val := binary.LittleEndian.Uint32(data[index:])
			index += 4

			if v.Type == 0x40 { // signed
				if int32(val) == -0x80000000 {
					break
				}
				v.Values = append(v.Values, int32(val))

			} else { // unsigned
				if val == 0xffffffff {
					break
				}

				v.Values = append(v.Values, val)
			}
		}
		return 8 + 5*4, nil
	}
	return 8, nil
}

// counter for packet id - increased on every packet
var packetIDCounter uint32

// more or less unique ID of the current system
var systemID *DeviceId

// NewDeviceData creates a device data request
func NewDeviceData(control uint8) *DeviceData {
	pkgId := atomic.AddUint32(&packetIDCounter, 1)

	// initialize system id on first call
	if systemID == nil {
		systemID = LocalDeviceId()
	}

	return &DeviceData{
		Control: control,
		Source:  *systemID,
		// count up only last byte
		PacketID: uint16(pkgId & 0xFF),
	}
}

// DeviceData sub packet
type DeviceData struct {
	Control uint8

	Destination DeviceId
	JobNumber   uint8
	Source      DeviceId

	Status      uint16
	PacketCount uint16
	PacketID    uint16

	Command uint8
	Object  uint16

	Parameters []uint32

	// used for responses
	ResponseValues []*ResponseValue

	// used for requests
	Data []byte
}

// ProtocolID identifies packet type
func (d *DeviceData) ProtocolID() uint16 {
	return DeviceDataProtocolID
}

// Bytes returns binary data
func (d *DeviceData) Bytes() []byte {
	// package length
	parameterCount := len(d.Parameters)
	var length int
	if d.Data == nil {
		length = 28 + parameterCount*4
	} else {
		length = 28 + parameterCount*4 + len(d.Data)
	}

	data := make([]byte, length)
	data[0] = uint8(length / 4)

	data[1] = d.Control
	copy(data[2:], d.Destination.Bytes(binary.LittleEndian))

	data[8] = 0 // unknown
	data[9] = d.JobNumber

	copy(data[10:], d.Source.Bytes(binary.LittleEndian))

	data[16] = 0 // unknown
	data[17] = d.JobNumber

	binary.LittleEndian.PutUint16(data[18:], d.Status)
	binary.LittleEndian.PutUint16(data[20:], d.PacketCount)
	binary.LittleEndian.PutUint16(data[22:], d.PacketID|0x8000)

	data[24] = d.Command
	data[25] = uint8(parameterCount)

	binary.LittleEndian.PutUint16(data[26:], d.Object)

	for i, param := range d.Parameters {
		binary.LittleEndian.PutUint32(data[28+4*i:], param)
	}

	if d.Data != nil {
		copy(data[28+4*parameterCount:], d.Data)
	}
	return data
}

// Read binary representation
func (d *DeviceData) Read(data []byte) error {
	if len(data) < 30 {
		return fmt.Errorf("invalid DeviceData - length %d", len(data))
	}

	// validate data size
	length := int(data[0]) * 4
	if len(data) != length {
		return fmt.Errorf(
			"invalid sma net2 data size. expected %d - get %d",
			length, len(data))
	}

	d.Control = uint8(data[1])

	err := d.Destination.Read(data[2:], binary.LittleEndian)
	if err != nil {
		return err
	}

	// data[8] - unknown

	d.JobNumber = data[9]

	err = d.Source.Read(data[10:], binary.LittleEndian)
	if err != nil {
		return err
	}

	// data[16] - unknown
	// data[17] - JobNumber (again ?)

	d.Status = binary.LittleEndian.Uint16(data[18:])
	d.PacketCount = binary.LittleEndian.Uint16(data[20:])
	d.PacketID = binary.LittleEndian.Uint16(data[22:]) & ^uint16(0x8000)

	d.Command = uint8(data[24])
	parameterCount := int(data[25])
	d.Object = binary.LittleEndian.Uint16(data[26:])

	d.Parameters = make([]uint32, parameterCount)
	index := 28
	// parse parameters
	for i := 0; i < parameterCount; i++ {
		d.Parameters[i] = binary.LittleEndian.Uint32(data[index:])
		index += 4
	}

	// no data or response
	dataLength := len(data)
	if dataLength-index <= 0 || d.Command != 0x01 {
		return nil
	}

	// load response variables
	d.ResponseValues = make([]*ResponseValue, 0)
	for dataLength-index > 8 {
		val := new(ResponseValue)
		n, err := val.Read(data[index:], d.Object)
		if err != nil {
			return err
		}

		index += n

		d.ResponseValues = append(d.ResponseValues, val)
	}

	return nil
}

// AddParameter to sub packet
func (d *DeviceData) AddParameter(param uint32) {
	if d.Parameters == nil {
		d.Parameters = make([]uint32, 0)
	}
	d.Parameters = append(d.Parameters, param)
}
