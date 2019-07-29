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
	"gitlab.com/bboehmke/sunny/proto/net2"
)

// valDef defines a value of an inverter device
type valDef struct {
	Object uint16
	Start  uint32
	End    uint32

	Class uint8 // 0 -> ignore class
	Code  uint16

	Key string
}

var valuesDef = []valDef{
	{0x5100, 0x00263F00, 0x00263FFF, 0x00, 0x263F, "power_ac_total"},
	{0x5100, 0x00295A00, 0x00295AFF, 0x00, 0x295A, "battery_charge"},
	{0x5100, 0x00411E00, 0x004120FF, 0x00, 0x411E, "power_max"},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4640, "power_ac1"},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4641, "power_ac2"},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4642, "power_ac3"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4648, "voltage_ac1"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4649, "voltage_ac2"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x464a, "voltage_ac3"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4653, "current_ac1"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4654, "current_ac2"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4655, "current_ac3"},
	{0x5100, 0x00465700, 0x004657FF, 0x00, 0x4657, "frequency_ac"},
	{0x5100, 0x00491E00, 0x00495DFF, 0x00, 0x495B, "battery_temperature"},

	{0x5180, 0x00214800, 0x002148FF, 0x00, 0x2148, "device_status"},
	{0x5180, 0x00416400, 0x004164FF, 0x00, 0x4164, "device_grid_relay"},

	{0x5200, 0x00237700, 0x002377FF, 0x00, 0x2377, "device_temperature"},

	{0x5380, 0x00251E00, 0x00251EFF, 0x01, 0x251E, "power_dc1"},
	{0x5380, 0x00251E00, 0x00251EFF, 0x02, 0x251E, "power_dc2"},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x451F, "voltage_dc1"},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x451F, "voltage_dc2"},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x4521, "current_dc1"},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x4521, "current_dc2"},

	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2601, "energy_total"},
	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2622, "energy_today"},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462E, "time_operating"},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462F, "time_feed"},

	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821E, "device_name"},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821F, "device_class"},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x8220, "device_type"},
}

// cache for responses and requests
var responseValues map[uint32]string
var allRequests []valDef
var requestMap map[string]valDef

// init cache
func init() {
	if responseValues != nil {
		return
	}

	responseValues = make(map[uint32]string, len(valuesDef))
	for _, def := range valuesDef {
		responseValues[uint32(def.Code)<<16+uint32(def.Class)] = def.Key
	}

	allRequests = getRequests(valuesDef)

	requestMap = make(map[string]valDef)
	for _, def := range valuesDef {
		requestMap[def.Key] = def
	}
}

// checkValue checks if response is a known value
func checkValue(value *net2.ResponseValue) string {
	if def, ok := responseValues[uint32(value.Code)<<16+uint32(value.Class)]; ok {
		return def
	}
	if def, ok := responseValues[uint32(value.Code)<<16]; ok {
		return def
	}
	return ""
}

// getAllRequests to receive all values
func getAllRequests() []valDef {
	return allRequests
}

// getRequest for given key
func getRequest(key string) valDef {
	return requestMap[key]
}

// getRequests to receive all of the given values (reduce request amount)
func getRequests(values []valDef) []valDef {
	defs := make([]valDef, 0, len(values))

	for _, value := range values {

		found := false
		for _, def := range defs {
			if value.Object == def.Object &&
				value.Start == def.Start &&
				value.End == def.End {
				found = true
				break
			}
		}

		if !found {
			defs = append(defs, value)
		}
	}
	return defs
}

// parseValues from response
func parseValues(values []*net2.ResponseValue) map[string]interface{} {
	data := make(map[string]interface{}, len(values))

	for _, val := range values {
		if len(val.Values) == 0 {
			continue
		}

		if key := checkValue(val); key != "" {
			data[key] = val.Values[0]
		}
	}
	return data
}
