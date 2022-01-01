// Copyright 2021 Benjamin Böhmke <benjamin@boehmke.net>.
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
	"sync"
	"time"

	"gitlab.com/bboehmke/sunny/proto"
)

// SimpleDiscoverDevices in Connection with a simpler interface
func (c *Connection) SimpleDiscoverDevices(password string) []*Device {
	// add found devices to list
	var wg sync.WaitGroup
	wg.Add(1)
	devices := make(chan *Device, 10)

	var deviceList []*Device
	go func() {
		for device := range devices {
			deviceList = append(deviceList, device)
		}
		wg.Done()
	}()

	// search for devices
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	c.DiscoverDevices(ctx, devices, password)
	cancel()

	close(devices)
	wg.Wait()
	return deviceList
}

// DiscoverDevices in Connection
func (c *Connection) DiscoverDevices(ctx context.Context, devices chan *Device, password string) {
	var wg sync.WaitGroup
	knownIps := make(map[string]*Device)
	var knownMutex sync.Mutex
	ticker := time.NewTicker(time.Millisecond * 500)

	discoverCh := make(chan string)
	c.registerDiscoverer(discoverCh)
	defer c.unregisterDiscoverer(discoverCh)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop

		// handle received responses
		case ip := <-discoverCh:
			wg.Add(1)
			go func(ip string) {
				knownMutex.Lock()
				defer knownMutex.Unlock()

				if _, ok := knownIps[ip]; !ok {
					device, err := c.NewDevice(ip, password)
					if err != nil {
						Log.Printf("discover - skip ip %s: %v", ip, err)
					} else {
						Log.Printf("found device %d at %s", device.SerialNumber(), ip)
						knownIps[ip] = device
						devices <- device
					}
				}

				wg.Done()
			}(ip)

		// send discover packages
		case <-ticker.C:
			// send discover packet
			Log.Printf("send discover package")
			_, err := c.socket.WriteTo(proto.NewDiscoveryRequest().Bytes(), c.address)
			if err != nil {
				Log.Printf("failed to send packet: %w", err)
			}
		}
	}
	ticker.Stop()
	wg.Wait()
}
