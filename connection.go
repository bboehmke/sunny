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
	"fmt"
	"net"
	"sync"
	"time"

	"gitlab.com/bboehmke/sunny/proto"
)

const listenAddress = "239.12.255.254:9522"

var listenInterface *net.Interface
var interfaceMutex sync.RWMutex

// SetMulticastInterface for communication with devices
func SetMulticastInterface(name string) (err error) {
	interfaceMutex.Lock()
	defer interfaceMutex.Unlock()

	if name == "" {
		listenInterface = nil
	} else {
		listenInterface, err = net.InterfaceByName(name)
	}
	return
}

var conn *connection

// DiscoverDevices in network
func DiscoverDevices(password string) ([]*Device, error) {
	c, err := getConnection()
	if err != nil {
		return nil, err
	}

	addresses, err := c.discover()
	if err != nil {
		return nil, err
	}

	devices := make([]*Device, 0, len(addresses))
	for _, ip := range addresses {
		device, err := NewDevice(ip, password)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

// getConnection instance and initialize it if first access
func getConnection() (*connection, error) {
	if conn == nil {
		c, err := newConnection()
		if err != nil {
			return nil, err
		}
		conn = c
	}
	return conn, nil
}

// Connection for communication with devices
type connection struct {
	mutex sync.RWMutex

	// multicast address
	address *net.UDPAddr
	// multicast socket
	socket *net.UDPConn

	// buffer for received packet
	receivedBuffer map[string]chan *proto.Packet
	// list of discovered devices
	discoveredDevices map[string]*proto.Packet
}

// newConnection creates a new connection object and starts listening
func newConnection() (*connection, error) {
	conn := connection{
		receivedBuffer:    make(map[string]chan *proto.Packet),
		discoveredDevices: make(map[string]*proto.Packet),
	}

	var err error
	conn.address, err = net.ResolveUDPAddr("udp", listenAddress)
	if err != nil {
		return nil, err
	}

	interfaceMutex.RLock()
	conn.socket, err = net.ListenMulticastUDP("udp", listenInterface, conn.address)
	interfaceMutex.RUnlock()
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %v", err)
	}

	err = conn.socket.SetReadBuffer(2048)
	if err != nil {
		return nil, err
	}

	go conn.listenLoop()

	return &conn, nil
}

// listenLoop for received packets
func (c *connection) listenLoop() {
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
			continue
		}

		// store discovery responses
		if pack.GetEntry(proto.DiscoveryRequestPacketEntryTag) != nil {
			c.mutex.Lock()
			c.discoveredDevices[src.IP.String()] = &pack
			c.mutex.Unlock()
		}

		select {
		case c.getRecvChannel(src) <- &pack:
		case <-time.After(time.Millisecond):
			// channel for received packets busy -> drop packet
		}
	}
}

// getRecvChannel returns the receive channel of this address
func (c *connection) getRecvChannel(address *net.UDPAddr) chan *proto.Packet {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.receivedBuffer[address.IP.String()]; !ok {
		c.receivedBuffer[address.IP.String()] = make(chan *proto.Packet, 5)
	}

	return c.receivedBuffer[address.IP.String()]
}

// clearReceived packages of the given address
func (c *connection) clearReceived(address *net.UDPAddr) {
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
func (c *connection) sendPacket(address *net.UDPAddr, packet *proto.Packet) error {
	_, err := c.socket.WriteToUDP(packet.Bytes(), address)
	if err != nil {
		return fmt.Errorf("failed to send packet: %v", err)
	}
	return nil
}

// readPacket from received channel
func (c *connection) readPacket(address *net.UDPAddr, timeout time.Duration) *proto.Packet {
	ch := c.getRecvChannel(address)

	select {
	case pack := <-ch:
		return pack
	case <-time.After(timeout):
		return nil
	}
}

// discover reachable devices
func (c *connection) discover() ([]string, error) {
	// send discover packet
	_, err := c.socket.WriteTo(proto.NewDiscoveryRequest().Bytes(), conn.address)
	if err != nil {
		return nil, fmt.Errorf("failed to send packet: %v", err)
	}

	// wait some time for responses
	time.Sleep(time.Second) // to much?

	c.mutex.RLock()
	defer c.mutex.RUnlock()
	deviceAddresses := make([]string, 0, len(c.discoveredDevices))
	for ip := range c.discoveredDevices {
		deviceAddresses = append(deviceAddresses, ip)
	}

	return deviceAddresses, nil
}
