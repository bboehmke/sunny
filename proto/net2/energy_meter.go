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
)

// EnergyMeterPacketProtocolID protocol ID used for EnergyMeter sub packets
const EnergyMeterPacketProtocolID uint16 = 0x6069

// OBISIdentifier for values
type OBISIdentifier struct {
	Channel          uint8
	MeasurementValue uint8
	MeasurementType  uint8
	Tariff           uint8
}

// Bytes returns binary data
func (o *OBISIdentifier) Bytes() []byte {
	return []byte{
		o.Channel,
		o.MeasurementValue,
		o.MeasurementType,
		o.Tariff,
	}
}

// Read binary representation
func (o *OBISIdentifier) Read(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("invalid OBISIdentifier - length %d", len(data))
	}

	o.Channel = data[0]
	o.MeasurementValue = data[1]
	o.MeasurementType = data[2]
	o.Tariff = data[3]
	return nil
}

// String representation of identifier
func (o OBISIdentifier) String() string {
	return fmt.Sprintf("%d:%d.%d.%d",
		o.Channel, o.MeasurementValue, o.MeasurementType, o.Tariff)
}

// MeasuredData received from energy meter
type MeasuredData struct {
	OBIS  OBISIdentifier
	Value interface{}
}

// Bytes returns binary data
func (e *MeasuredData) Bytes() []byte {
	var data []byte
	if e.OBIS.MeasurementType == 8 {
		data = make([]byte, 12)
		binary.BigEndian.PutUint64(data[4:], e.Value.(uint64))
	} else {
		data = make([]byte, 8)
		binary.BigEndian.PutUint32(data[4:], e.Value.(uint32))
	}

	copy(data[:4], e.OBIS.Bytes())
	return data
}

// Read binary representation
func (e *MeasuredData) Read(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("invalid MeasuredData - length %d", len(data))
	}

	err := e.OBIS.Read(data)
	if err != nil {
		return err
	}

	if e.OBIS.MeasurementType == 8 {
		if len(data) < 12 {
			return fmt.Errorf("invalid large MeasuredData - length %d", len(data))
		}
		e.Value = binary.BigEndian.Uint64(data[4:])
	} else {
		e.Value = binary.BigEndian.Uint32(data[4:])
	}
	return nil
}

// EnergyMeterPacket contains response of an energy meter
type EnergyMeterPacket struct {
	// energy meter identifier
	Id DeviceId

	// ticker measuring time in ms (with overflow)
	Ticker uint32

	Values []*MeasuredData
}

// ProtocolID identifies packet type
func (e *EnergyMeterPacket) ProtocolID() uint16 {
	return EnergyMeterPacketProtocolID
}

// Bytes returns binary data
func (e *EnergyMeterPacket) Bytes() []byte {
	data := make([]byte, 10)

	copy(data, e.Id.Bytes(binary.BigEndian))
	binary.BigEndian.PutUint32(data[6:], e.Ticker)

	for _, v := range e.Values {
		data = append(data, v.Bytes()...)
	}
	return data
}

// Read binary representation
func (e *EnergyMeterPacket) Read(data []byte) error {
	if len(data) < 10 {
		return fmt.Errorf("invalid EnergyMeterPacket - length %d", len(data))
	}

	err := e.Id.Read(data, binary.BigEndian)
	if err != nil {
		return err
	}

	e.Ticker = binary.BigEndian.Uint32(data[6:])

	e.Values = make([]*MeasuredData, 0)

	dataLen := len(data)
	index := 10
	for dataLen-index >= 8 {
		val := new(MeasuredData)

		err := val.Read(data[index:])
		if err != nil {
			return err
		}

		if val.OBIS.MeasurementType == 8 {
			index += 12
		} else {
			index += 8
		}
		e.Values = append(e.Values, val)
	}

	return nil
}

// GetValues as list from packet
func (e *EnergyMeterPacket) GetValues() map[string]interface{} {
	values := make(map[string]interface{}, len(e.Values))
	for _, v := range e.Values {
		values[v.OBIS.String()] = v.Value
	}
	return values
}
