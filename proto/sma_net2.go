package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var smaNet2subPackets = []SmaNet2SubPacket{
	&EnergyMeterPacket{},
	&DeviceData{},
}

type SmaNet2SubPacket interface {
	ProtocolID() uint16
	Bytes() []byte
	Read(data []byte) (SmaNet2SubPacket, error)
}

const SmaNet2PacketEntryTag = 0x0010

type SmaNet2PacketEntry struct {
	Content SmaNet2SubPacket
}

func (e SmaNet2PacketEntry) Tag() uint16 {
	return SmaNet2PacketEntryTag
}
func (e *SmaNet2PacketEntry) Bytes() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, e.Content.ProtocolID())
	return append(b, e.Content.Bytes()...)
}
func (e *SmaNet2PacketEntry) Read(data []byte) (PacketEntry, error) {
	err := checkLen(data, 2)
	if err != nil {
		return nil, err
	}

	protoID := binary.BigEndian.Uint16(data[0:2])

	for _, p := range smaNet2subPackets {
		if p.ProtocolID() == protoID {
			packet, err := p.Read(data[2:])
			if err != nil {
				return nil, err
			}

			return &SmaNet2PacketEntry{
				Content: packet,
			}, nil
		}
	}
	return &SmaNet2PacketEntry{
		Content: nil,
	}, nil
	//return nil, fmt.Errorf("invalid protocol id: %d", protoID)
}

type OBISIdentifier struct {
	Channel          uint8
	MeasurementValue uint8
	MeasurementType  uint8
	Tariff           uint8
}

func (o *OBISIdentifier) Bytes() []byte {
	b := make([]byte, 4)
	b[0] = o.Channel
	b[1] = o.MeasurementValue
	b[2] = o.MeasurementType
	b[3] = o.Tariff
	return b
}
func (o OBISIdentifier) String() string {
	return fmt.Sprintf("%d:%d.%d.%d",
		o.Channel, o.MeasurementValue, o.MeasurementType, o.Tariff)
}

type MeasuredData struct {
	OBIS  OBISIdentifier
	Value uint64
}

func (e *MeasuredData) Bytes() []byte {
	var b []byte
	if e.OBIS.MeasurementType == 8 {
		b = make([]byte, 8)
		binary.BigEndian.PutUint64(b, e.Value)
	} else {
		b = make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(e.Value))
	}

	return append(e.OBIS.Bytes(), b...)
}

type EnergyMeterPacket struct {
	// energy meter identifier
	SusyID uint16
	SerNo  uint32

	// ticker measuring time in ms (with overflow)
	Ticker uint32

	Values []MeasuredData
}

func (e EnergyMeterPacket) ProtocolID() uint16 {
	return 0x6069
}
func (e *EnergyMeterPacket) Bytes() []byte {
	var buffer bytes.Buffer
	b := make([]byte, 4)
	binary.BigEndian.PutUint16(b, e.SusyID)
	buffer.Write(b[:2])
	binary.BigEndian.PutUint32(b, e.SerNo)
	buffer.Write(b[:4])

	binary.BigEndian.PutUint32(b, e.Ticker)
	buffer.Write(b[:4])

	for _, v := range e.Values {
		buffer.Write(v.Bytes())
	}
	return buffer.Bytes()
}
func (e *EnergyMeterPacket) Read(data []byte) (SmaNet2SubPacket, error) {
	err := checkLen(data, 18)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(data)
	p := EnergyMeterPacket{
		SusyID: binary.BigEndian.Uint16(buffer.Next(2)),
		SerNo:  binary.BigEndian.Uint32(buffer.Next(4)),
		Ticker: binary.BigEndian.Uint32(buffer.Next(4)),
		Values: make([]MeasuredData, 0),
	}

	for buffer.Len() >= 8 {
		var data MeasuredData
		obis := buffer.Next(4)

		data.OBIS.Channel = obis[0]
		data.OBIS.MeasurementValue = obis[1]
		data.OBIS.MeasurementType = obis[2]
		data.OBIS.Tariff = obis[3]

		if data.OBIS.MeasurementType == 8 {
			data.Value = binary.BigEndian.Uint64(buffer.Next(8))
		} else {
			data.Value = uint64(binary.BigEndian.Uint32(buffer.Next(4)))
		}
		p.Values = append(p.Values, data)
	}

	return &p, nil
}
func (e *EnergyMeterPacket) GetValues() map[string]interface{} {
	values := make(map[string]interface{}, len(e.Values))
	for _, v := range e.Values {
		values[v.OBIS.String()] = v.Value
	}
	return values
}

type ResponseValue struct {
	Class     uint8
	Code      uint16
	Type      uint8
	Timestamp uint32

	Values []interface{}
}

var packetIdCounter uint8

func NewDeviceData(control uint8) *DeviceData {
	packetIdCounter++

	return &DeviceData{
		Control: control,

		// sunny explorer values
		SrcSusyID:       120,
		SrcSerialNumber: 954830016,

		PacketID: uint16(packetIdCounter),

		Parameters:     make([]uint32, 0),
		ResponseValues: make([]ResponseValue, 0),
	}
}

type DeviceData struct {
	Control uint8

	DstSusyID       uint16
	DstSerialNumber uint32
	JobNumber       uint8
	SrcSusyID       uint16
	SrcSerialNumber uint32

	Status      uint16
	PacketCount uint16
	PacketID    uint16

	Command uint8
	Object  uint16

	Parameters []uint32

	ResponseValues []ResponseValue
	Data           []byte
}

func (d DeviceData) ProtocolID() uint16 {
	return 0x6065
}
func (d *DeviceData) Bytes() []byte {
	var buffer bytes.Buffer
	b := make([]byte, 4)

	if d.Data == nil {
		buffer.WriteByte(uint8((28 + len(d.Parameters)*4) / 4))
	} else {
		buffer.WriteByte(uint8((28 + len(d.Parameters)*4 + len(d.Data)) / 4))
	}

	buffer.WriteByte(d.Control)

	binary.LittleEndian.PutUint16(b, d.DstSusyID)
	buffer.Write(b[:2])
	binary.LittleEndian.PutUint32(b, d.DstSerialNumber)
	buffer.Write(b[:4])
	buffer.WriteByte(0)
	buffer.WriteByte(d.JobNumber)
	binary.LittleEndian.PutUint16(b, d.SrcSusyID)
	buffer.Write(b[:2])
	binary.LittleEndian.PutUint32(b, d.SrcSerialNumber)
	buffer.Write(b[:4])
	buffer.WriteByte(0)
	buffer.WriteByte(d.JobNumber)

	binary.LittleEndian.PutUint16(b, d.Status)
	buffer.Write(b[:2])
	binary.LittleEndian.PutUint16(b, d.PacketCount)
	buffer.Write(b[:2])
	binary.LittleEndian.PutUint16(b, d.PacketID|0x8000)
	buffer.Write(b[:2])

	buffer.WriteByte(d.Command)
	buffer.WriteByte(uint8(len(d.Parameters)))
	binary.LittleEndian.PutUint16(b, d.Object)
	buffer.Write(b[:2])

	for _, param := range d.Parameters {
		binary.LittleEndian.PutUint32(b, param)
		buffer.Write(b[:4])
	}

	if d.Data != nil {
		buffer.Write(d.Data)
	}

	return buffer.Bytes()
}
func (d *DeviceData) Read(data []byte) (SmaNet2SubPacket, error) {
	err := checkLen(data, 30)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(data)

	// validate data size
	length := int(buffer.Next(1)[0]) * 4
	if len(data) != length {
		return nil, fmt.Errorf(
			"invalid sma net2 data size. expected %d - get %d",
			length, len(data))
	}

	p := DeviceData{
		Parameters:     make([]uint32, 0),
		ResponseValues: make([]ResponseValue, 0),
	}

	p.Control = uint8(buffer.Next(1)[0])

	p.DstSusyID = binary.LittleEndian.Uint16(buffer.Next(2))
	p.DstSerialNumber = binary.LittleEndian.Uint32(buffer.Next(4))
	buffer.Next(1) // skip unknown
	p.JobNumber = uint8(buffer.Next(1)[0])
	p.SrcSusyID = binary.LittleEndian.Uint16(buffer.Next(2))
	p.SrcSerialNumber = binary.LittleEndian.Uint32(buffer.Next(4))
	buffer.Next(1) // skip unknown
	buffer.Next(1) // skip JobNumber

	p.Status = binary.LittleEndian.Uint16(buffer.Next(2))
	p.PacketCount = binary.LittleEndian.Uint16(buffer.Next(2))
	p.PacketID = binary.LittleEndian.Uint16(buffer.Next(2)) & ^uint16(0x8000)

	p.Command = uint8(buffer.Next(1)[0])
	parameterCount := int(buffer.Next(1)[0])
	p.Object = binary.LittleEndian.Uint16(buffer.Next(2))

	for i := 0; i < parameterCount; i++ {
		p.Parameters = append(p.Parameters,
			binary.LittleEndian.Uint32(buffer.Next(4)))
	}

	// no data or response
	if buffer.Len() == 0 || p.Command != 0x01 {
		return &p, nil
	}

	for buffer.Len() > 8 {
		responseValue := ResponseValue{
			Class:     uint8(buffer.Next(1)[0]),
			Code:      binary.LittleEndian.Uint16(buffer.Next(2)),
			Type:      uint8(buffer.Next(1)[0]),
			Timestamp: binary.LittleEndian.Uint32(buffer.Next(4)),
		}

		if responseValue.Type == 0x10 {
			responseValue.Values = append(responseValue.Values,
				string(buffer.Next(32)))

		} else if responseValue.Type == 0x08 {
			for i := 0; i < 8; i++ {
				if buffer.Len() < 4 {
					break
				}

				val := binary.LittleEndian.Uint32(buffer.Next(4))

				if val == 0xfffffe {
					break
				}
				if val>>24 == 1 {
					responseValue.Values = append(responseValue.Values, val&0xffffff)
				}
			}

		} else if p.Object == 0x5400 {
			if buffer.Len() < 8 {
				break
			}
			responseValue.Values = append(responseValue.Values,
				binary.LittleEndian.Uint64(buffer.Next(8)))

		} else if responseValue.Type == 0x00 {
			for i := 0; i < 5; i++ {
				if buffer.Len() < 4 {
					break
				}

				val := binary.LittleEndian.Uint32(buffer.Next(4))

				if val == 0xffffffff {
					break
				}

				responseValue.Values = append(responseValue.Values, val)
			}

		} else if responseValue.Type == 0x40 {
			for i := 0; i < 5; i++ {
				if buffer.Len() < 4 {
					break
				}

				val := int32(binary.LittleEndian.Uint32(buffer.Next(4)))

				if val == -0x80000000 {
					break
				}

				responseValue.Values = append(responseValue.Values, val)
			}
		}
		p.ResponseValues = append(p.ResponseValues, responseValue)
	}

	return &p, nil
}
func (d *DeviceData) AddParameter(param uint32) {
	d.Parameters = append(d.Parameters, param)
}
