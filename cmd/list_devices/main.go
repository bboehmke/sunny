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

package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"

	"gitlab.com/bboehmke/sunny"
)

var inf = flag.String("inf", "", "Interface devices are connected to")

func main() {
	flag.Parse()
	//sunny.Log = log.Default()

	var wg sync.WaitGroup
	wg.Add(1)
	devices := make(chan *sunny.Device, 10)

	go func() {
		for device := range devices {
			fmt.Printf("==================================================\n")
			fmt.Printf("IP:             %s\n", device.Address())
			fmt.Printf("Serial:         %d\n", device.SerialNumber())
			fmt.Printf("Is EnergyMeter: %v\n", device.IsEnergyMeter())
			fmt.Printf("--------------------------------------------------\n")
			name, err := device.GetValue(sunny.DeviceName)
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
			} else {
				fmt.Printf("Name: %s\n", name)
			}
			fmt.Printf("--------------------------------------------------\n")
			values, err := device.GetValues()
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
			} else {
				for key, value := range values {
					switch value.(type) {
					case float64:
						fmt.Printf("%s: %f %s\n", key, value, sunny.GetValueInfo(key).Unit)
					default:
						fmt.Printf("%s: %v %s\n", key, value, sunny.GetValueInfo(key).Unit)
					}
				}
			}
			fmt.Printf("==================================================\n")
			fmt.Println()
		}
		wg.Done()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	connection, err := sunny.NewConnection(*inf)
	if err != nil {
		panic(err)
	}
	connection.DiscoverDevices(ctx, devices, "0000")
	cancel()

	close(devices)
	wg.Wait()
}
