package sunny

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/bboehmke/sunny/proto"
)

const listenAddress = "239.12.255.254:9522"

var conn *connection

func DiscoverDevices(password string) ([]*Device, error) {
	c, err := getConnection(nil)
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

func getConnection(address *net.UDPAddr) (*connection, error) {
	if conn == nil {
		c, err := newConnection()
		if err != nil {
			return nil, err
		}
		conn = c
	}
	if address != nil {
		srcIP := address.IP.String()
		if _, ok := conn.receivedBuffer[srcIP]; !ok {
			conn.Lock()
			conn.receivedBuffer[srcIP] = make(chan *proto.Packet, 5)
			conn.Unlock()
		}
	}
	return conn, nil
}

type connection struct {
	sync.RWMutex
	socket *net.UDPConn

	receivedBuffer    map[string]chan *proto.Packet
	discoveredDevices map[string]*proto.Packet
}

func newConnection() (*connection, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", listenAddress)
	if err != nil {
		return nil, err
	}

	conn := connection{
		receivedBuffer:    make(map[string]chan *proto.Packet),
		discoveredDevices: make(map[string]*proto.Packet),
	}
	conn.socket, err = net.ListenMulticastUDP("udp", nil, udpAddr)
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

func (c *connection) listenLoop() {
	b := make([]byte, 2048)

	for c.socket != nil {
		n, src, err := c.socket.ReadFromUDP(b)
		if err != nil {
			logrus.Debug("ReadFromUDP failed:", err)
			continue
		}

		var pack proto.Packet
		err = pack.Read(b[:n])
		if err != nil {
			logrus.Debug("invalid package:", err)
			continue
		}

		srcIP := src.IP.String()
		if _, ok := c.receivedBuffer[srcIP]; !ok {
			c.Lock()
			c.receivedBuffer[srcIP] = make(chan *proto.Packet, 5)
			c.Unlock()
		}

		// store discovery responses
		if pack.GetEntry(proto.DiscoveryRequestPacketEntryTag) != nil {
			c.Lock()
			c.discoveredDevices[srcIP] = &pack
			c.Unlock()
		}

		select {
		case c.receivedBuffer[srcIP] <- &pack:
		case <-time.After(time.Millisecond):
			logrus.Debug("Drop package:", pack.String())
		}
	}
}

func (c *connection) getRecvChannel(address *net.UDPAddr) chan *proto.Packet {
	c.RLock()
	defer c.RUnlock()
	return c.receivedBuffer[address.IP.String()]
}

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

func (c *connection) sendPacket(address *net.UDPAddr, packet *proto.Packet) error {
	_, err := c.socket.WriteToUDP(packet.Bytes(), address)
	if err != nil {
		return fmt.Errorf("failed to send packet: %v", err)
	}
	return nil
}

func (c *connection) readPacket(address *net.UDPAddr, timeout time.Duration) *proto.Packet {
	ch := c.getRecvChannel(address)

	select {
	case pack := <-ch:
		return pack
	case <-time.After(timeout):
		return nil
	}
}

func (c *connection) discover() ([]string, error) {
	address, err := net.ResolveUDPAddr("udp", "239.12.255.254:9522")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve udp address: %v", err)
	}

	_, err = c.socket.WriteTo(proto.NewDiscoveryRequest().Bytes(), address)
	if err != nil {
		return nil, fmt.Errorf("failed to send packet: %v", err)
	}

	time.Sleep(time.Second)

	deviceAddresses := make([]string, 0, len(c.discoveredDevices))
	for ip := range c.discoveredDevices {
		deviceAddresses = append(deviceAddresses, ip)
	}

	return deviceAddresses, nil
}
