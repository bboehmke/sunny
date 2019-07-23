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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOBISIdentifier_Bytes(t *testing.T) {
	ass := assert.New(t)

	obis := OBISIdentifier{
		Channel:          0x12,
		MeasurementValue: 0x34,
		MeasurementType:  0x56,
		Tariff:           0x78,
	}

	ass.Equal([]byte{
		0x12, 0x34, 0x56, 0x78,
	}, obis.Bytes())
}

func TestOBISIdentifier_Read(t *testing.T) {
	ass := assert.New(t)

	obis := new(OBISIdentifier)
	ass.EqualError(obis.Read([]byte{0x12, 0x34}),
		"invalid OBISIdentifier - length 2")

	ass.NoError(obis.Read([]byte{0x12, 0x34, 0x56, 0x78}))

	ass.Equal(uint8(0x12), obis.Channel)
	ass.Equal(uint8(0x34), obis.MeasurementValue)
	ass.Equal(uint8(0x56), obis.MeasurementType)
	ass.Equal(uint8(0x78), obis.Tariff)
}

func TestOBISIdentifier_String(t *testing.T) {
	ass := assert.New(t)

	obis := OBISIdentifier{
		Channel:          0x01,
		MeasurementValue: 0x02,
		MeasurementType:  0x03,
		Tariff:           0x04,
	}

	ass.Equal("1:2.3.4", obis.String())
}

func TestMeasuredData_Bytes(t *testing.T) {
	ass := assert.New(t)

	data := MeasuredData{
		OBIS: OBISIdentifier{
			Channel:          0x01,
			MeasurementValue: 0x02,
			MeasurementType:  0x03,
			Tariff:           0x04,
		},
		Value: uint32(0x12345678),
	}

	ass.Equal([]byte{
		0x01, 0x02, 0x03, 0x04,
		0x12, 0x34, 0x56, 0x78,
	}, data.Bytes())

	data = MeasuredData{
		OBIS: OBISIdentifier{
			Channel:          0x01,
			MeasurementValue: 0x02,
			MeasurementType:  0x08,
			Tariff:           0x04,
		},
		Value: uint64(0x1234567812345678),
	}

	ass.Equal([]byte{
		0x01, 0x02, 0x08, 0x04,
		0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x56, 0x78,
	}, data.Bytes())
}

func TestMeasuredData_Read(t *testing.T) {
	ass := assert.New(t)

	data := new(MeasuredData)
	ass.EqualError(data.Read([]byte{0x12, 0x34}),
		"invalid MeasuredData - length 2")

	ass.NoError(data.Read([]byte{
		0x01, 0x02, 0x03, 0x04,
		0x12, 0x34, 0x56, 0x78,
	}))

	ass.Equal(uint8(0x01), data.OBIS.Channel)
	ass.Equal(uint8(0x02), data.OBIS.MeasurementValue)
	ass.Equal(uint8(0x03), data.OBIS.MeasurementType)
	ass.Equal(uint8(0x04), data.OBIS.Tariff)
	ass.Equal(uint32(0x12345678), data.Value)

	ass.EqualError(data.Read([]byte{
		0x01, 0x02, 0x08, 0x04,
		0x12, 0x34, 0x56, 0x78,
	}), "invalid large MeasuredData - length 8")

	ass.NoError(data.Read([]byte{
		0x01, 0x02, 0x08, 0x04,
		0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x56, 0x78,
	}))

	ass.Equal(uint8(0x01), data.OBIS.Channel)
	ass.Equal(uint8(0x02), data.OBIS.MeasurementValue)
	ass.Equal(uint8(0x08), data.OBIS.MeasurementType)
	ass.Equal(uint8(0x04), data.OBIS.Tariff)
	ass.Equal(uint64(0x1234567812345678), data.Value)
}

func TestEnergyMeterPacket_ProtocolID(t *testing.T) {
	ass := assert.New(t)

	packet := new(EnergyMeterPacket)
	ass.Equal(uint16(0x6069), packet.ProtocolID())
	ass.Equal(uint16(0x6069), EnergyMeterPacketProtocolID)
}

func TestEnergyMeterPacket_Bytes(t *testing.T) {
	ass := assert.New(t)

	packet := EnergyMeterPacket{
		Id: DeviceId{
			SusyID:       0x1234,
			SerialNumber: 0x12345678,
		},
		Ticker: 0x21436587,
		Values: []*MeasuredData{{
			OBIS: OBISIdentifier{
				Channel:          0x01,
				MeasurementValue: 0x02,
				MeasurementType:  0x03,
				Tariff:           0x04,
			},
			Value: uint32(0x12345678),
		}, {
			OBIS: OBISIdentifier{
				Channel:          0x01,
				MeasurementValue: 0x02,
				MeasurementType:  0x08,
				Tariff:           0x04,
			},
			Value: uint64(0x1234567812345678),
		}},
	}

	ass.Equal([]byte{
		0x12, 0x34, 0x12, 0x34, 0x56, 0x78, // id
		0x21, 0x43, 0x65, 0x87, // ticker
		0x01, 0x02, 0x03, 0x04, 0x12, 0x34, 0x56, 0x78, // value 1
		0x01, 0x02, 0x08, 0x04, 0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x56, 0x78, // value 2
	}, packet.Bytes())
}

func TestEnergyMeterPacket_Read(t *testing.T) {
	ass := assert.New(t)

	packet := new(EnergyMeterPacket)

	ass.EqualError(packet.Read([]byte{0x12, 0x34}),
		"invalid EnergyMeterPacket - length 2")

	ass.NoError(packet.Read([]byte{
		0x12, 0x34, 0x12, 0x34, 0x56, 0x78, // id
		0x21, 0x43, 0x65, 0x87, // ticker
		0x01, 0x02, 0x03, 0x04, 0x12, 0x34, 0x56, 0x78, // value 1
		0x01, 0x02, 0x08, 0x04, 0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x56, 0x78, // value 2
	}))

	ass.Equal(uint16(0x1234), packet.Id.SusyID)
	ass.Equal(uint32(0x12345678), packet.Id.SerialNumber)

	ass.Equal(uint32(0x21436587), packet.Ticker)

	ass.Len(packet.Values, 2)

	ass.Equal(uint8(0x01), packet.Values[0].OBIS.Channel)
	ass.Equal(uint8(0x02), packet.Values[0].OBIS.MeasurementValue)
	ass.Equal(uint8(0x03), packet.Values[0].OBIS.MeasurementType)
	ass.Equal(uint8(0x04), packet.Values[0].OBIS.Tariff)
	ass.Equal(uint32(0x12345678), packet.Values[0].Value)

	ass.Equal(uint8(0x01), packet.Values[1].OBIS.Channel)
	ass.Equal(uint8(0x02), packet.Values[1].OBIS.MeasurementValue)
	ass.Equal(uint8(0x08), packet.Values[1].OBIS.MeasurementType)
	ass.Equal(uint8(0x04), packet.Values[1].OBIS.Tariff)
	ass.Equal(uint64(0x1234567812345678), packet.Values[1].Value)
}

func TestEnergyMeterPacket_GetValues(t *testing.T) {
	ass := assert.New(t)

	packet := EnergyMeterPacket{
		Values: []*MeasuredData{{
			OBIS: OBISIdentifier{
				Channel:          0x01,
				MeasurementValue: 0x02,
				MeasurementType:  0x03,
				Tariff:           0x04,
			},
			Value: uint32(0x12345678),
		}, {
			OBIS: OBISIdentifier{
				Channel:          0x01,
				MeasurementValue: 0x02,
				MeasurementType:  0x08,
				Tariff:           0x04,
			},
			Value: uint64(0x1234567812345678),
		}},
	}

	values := packet.GetValues()

	ass.Len(values, 2)
	ass.Contains(values, "1:2.3.4")
	ass.Contains(values, "1:2.8.4")

	ass.Equal(uint32(0x12345678), values["1:2.3.4"])
	ass.Equal(uint64(0x1234567812345678), values["1:2.8.4"])
}
