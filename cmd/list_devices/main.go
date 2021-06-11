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
	"fmt"

	"gitlab.com/bboehmke/sunny"
)

func main() {
	// sunny.Log = log.Default()

	connection, err := sunny.NewConnection("")
	if err != nil {
		panic(err)
	}

	devices, err := connection.DiscoverDevices("0000")
	if err != nil {
		panic(err)
	}

	for _, device := range devices {
		fmt.Printf("==================================================\n")
		fmt.Printf("IP:             %s\n", device.Address())
		fmt.Printf("Serial:         %d\n", device.SerialNumber())
		fmt.Printf("Is EnergyMeter: %v\n", device.IsEnergyMeter())
		fmt.Printf("--------------------------------------------------\n")
		name, err := device.GetDeviceName()
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
				fmt.Printf("%s: %v %s\n", key, value, device.GetValueInfo(key).Unit)
			}
		}
		fmt.Printf("==================================================\n")
		fmt.Println()
	}
}
