package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
)

var packets = []ReadPacketEntry{
	&GroupPacketEntry{},

	&SmaNet2PacketEntry{},

	&DiscoveryRequestPacketEntry{},
	&DiscoveryIpPacketEntry{},
}

type PacketEntry interface {
	Tag() uint16
	Bytes() []byte
}
type ReadPacketEntry interface {
	PacketEntry
	Read(data []byte) (PacketEntry, error)
}

type Packet struct {
	entries []PacketEntry
}

func (p *Packet) String() string {
	names := make([]string, 0, len(p.entries))
	for _, e := range p.entries {
		names = append(names, reflect.TypeOf(e).Elem().String())
	}
	return strings.Join(names, ", ")
}

func (p *Packet) GetEntry(tag uint16) PacketEntry {
	for _, e := range p.entries {
		if e.Tag() == tag {
			return e
		}
	}
	return nil
}

func (p *Packet) AddEntry(entry PacketEntry) {
	if p.entries == nil {
		p.entries = make([]PacketEntry, 0)
	}
	p.entries = append(p.entries, entry)
}

func (p *Packet) Bytes() []byte {
	var buffer bytes.Buffer
	// packet header
	buffer.Write([]byte{'S', 'M', 'A', 0})

	for _, e := range p.entries {
		b := make([]byte, 2)

		binary.BigEndian.PutUint16(b, uint16(len(e.Bytes())))
		buffer.Write(b)

		binary.BigEndian.PutUint16(b, e.Tag())
		buffer.Write(b)

		buffer.Write(e.Bytes())
	}

	// add 4 empty bytes to the end
	buffer.Write(make([]byte, 4))

	return buffer.Bytes()
}
func (p *Packet) Read(data []byte) error {
	buffer := bytes.NewBuffer(data)

	if buffer.Len() < 20 {
		return fmt.Errorf("invalid packet - to small: %d", len(data))
	}

	head := buffer.Next(4)
	if head[0] != 'S' || head[1] != 'M' || head[2] != 'A' || head[3] != 0 {
		return fmt.Errorf("invalid packet - header: %s", string(head[:3]))
	}

	p.entries = make([]PacketEntry, 0)
	for buffer.Len() >= 4 {
		length := binary.BigEndian.Uint16(buffer.Next(2))
		tag := binary.BigEndian.Uint16(buffer.Next(2)) // including version

		if length == 0 {
			// last package
			break
		}

		found := false
		for _, packet := range packets {
			if packet.Tag() == tag {
				entry, err := packet.Read(buffer.Next(int(length)))
				if err != nil {
					return err
				}
				p.entries = append(p.entries, entry)
				found = true
				break
			}

		}
		if !found {
			p.entries = append(p.entries, &UnknownPacketEntry{
				Data: buffer.Next(int(length)),
				T:    tag,
			})
		}
	}
	return nil
}

func checkLen(data []byte, length int) error {
	if len(data) >= length {
		return nil
	}
	return fmt.Errorf("invalid length %d - required %d", len(data), length)
}

const GroupPacketEntryTag = 0x02A0

type GroupPacketEntry struct {
	Group uint32
}

func (e *GroupPacketEntry) Tag() uint16 {
	return GroupPacketEntryTag
}
func (e *GroupPacketEntry) Bytes() []byte {
	return []byte{
		byte(e.Group >> 24),
		byte(e.Group >> 16),
		byte(e.Group >> 8),
		byte(e.Group),
	}
}
func (e *GroupPacketEntry) Read(data []byte) (PacketEntry, error) {
	err := checkLen(data, 4)
	if err != nil {
		return nil, err
	}

	return &GroupPacketEntry{
		Group: binary.BigEndian.Uint32(data),
	}, nil
}

type UnknownPacketEntry struct {
	Data []byte
	T    uint16
}

func (e *UnknownPacketEntry) Tag() uint16 {
	return 0
}
func (e *UnknownPacketEntry) Bytes() []byte {
	return e.Data
}
