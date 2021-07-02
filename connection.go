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

package sunny

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"gitlab.com/bboehmke/sunny/proto"
)

const listenAddress = "239.12.255.254:9522"

var connectionMutex sync.Mutex
var connections = make(map[string]*Connection)

// Connection for communication with devices
type Connection struct {
	mutex sync.RWMutex

	// multicast address
	address *net.UDPAddr
	// multicast socket
	socket *net.UDPConn

	// buffer for received packet
	receivedBuffer map[string]chan *proto.Packet

	// interface for device discovery
	discoverMutex     sync.Mutex
	discoveredDevices chan string
}

// NewConnection creates a new Connection object and starts listening
func NewConnection(inf string) (*Connection, error) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()

	// connection already known
	if c, ok := connections[inf]; ok {
		return c, nil
	}

	conn := Connection{
		receivedBuffer:    make(map[string]chan *proto.Packet),
		discoveredDevices: make(chan string, 5),
	}

	var err error
	conn.address, err = net.ResolveUDPAddr("udp", listenAddress)
	if err != nil {
		return nil, err
	}

	// listen interface is optional
	var listenInterface *net.Interface
	if inf != "" {
		listenInterface, err = net.InterfaceByName(inf)
		if err != nil {
			return nil, err
		}
	}

	conn.socket, err = net.ListenMulticastUDP("udp", listenInterface, conn.address)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	err = conn.socket.SetReadBuffer(2048)
	if err != nil {
		return nil, err
	}

	go conn.listenLoop()

	connections[inf] = &conn
	return &conn, nil
}

// listenLoop for received packets
func (c *Connection) listenLoop() {
	b := make([]byte, 2048)

	for c.socket != nil {
		n, src, err := c.socket.ReadFromUDP(b)
		if err != nil {
			// failed to read from udp -> retry
			continue
		}

		var pack proto.Packet
		err = pack.Read(b[:n])
		if err != nil {
			// invalid packet received -> retry
			Log.Printf("received invalid package from %s: %v", src.IP.String(), err)
			continue
		}
		Log.Printf("received package from %s [%s]", src.IP.String(), pack)

		// forward discover packages
		select {
		case c.discoveredDevices <- src.IP.String():
		default:
		}

		select {
		case c.getRecvChannel(src) <- &pack:
		case <-time.After(time.Millisecond):
			// channel for received packets busy -> drop packet
		}
	}
}

// getRecvChannel returns the receive channel of this address
func (c *Connection) getRecvChannel(address *net.UDPAddr) chan *proto.Packet {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	srcIP := address.IP.String()
	if _, ok := c.receivedBuffer[srcIP]; !ok {
		c.receivedBuffer[srcIP] = make(chan *proto.Packet, 5)
	}

	return c.receivedBuffer[srcIP]
}

// clearReceived packages of the given address
func (c *Connection) clearReceived(address *net.UDPAddr) {
	ch := c.getRecvChannel(address)
	for {
		select {
		case <-ch:
		default:
			return
		}
	}
}

// sendPacket to the given address
func (c *Connection) sendPacket(address *net.UDPAddr, packet *proto.Packet) error {
	Log.Printf("send package to %s: [%s]", address.IP.String(), packet)
	_, err := c.socket.WriteToUDP(packet.Bytes(), address)
	if err != nil {
		return fmt.Errorf("failed to send packet: %w", err)
	}
	return nil
}

// readPacket from received channel
func (c *Connection) readPacket(address *net.UDPAddr, ctx context.Context) *proto.Packet {
	ch := c.getRecvChannel(address)

	select {
	case pack := <-ch:
		return pack
	case <-ctx.Done():
		return nil
	}
}
