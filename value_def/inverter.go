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

package value_def

import "gitlab.com/bboehmke/sunny/proto/net2"

// InverterValuesDef defines a value of an inverter device
type InverterValuesDef struct {
	Object uint16
	Start  uint32
	End    uint32

	Class uint8 // 0 -> ignore class
	Code  uint16

	ID     ValueID
	Factor float64
}

// inverterValues contains all values that can be read from inverters
var inverterValues = []InverterValuesDef{
	{0x5100, 0x00263F00, 0x00263FFF, 0x00, 0x263F, ActivePowerPlus, 0},
	{0x5100, 0x00295A00, 0x00295AFF, 0x00, 0x295A, BatteryCharge, 0},
	{0x5100, 0x00411E00, 0x004120FF, 0x00, 0x411E, ActivePowerMax, 0},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4640, ActivePowerPlusL1, 0},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4641, ActivePowerPlusL2, 0},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4642, ActivePowerPlusL3, 0},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4648, VoltageL1, 0.01},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4649, VoltageL2, 0.01},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x464a, VoltageL3, 0.01},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4653, CurrentL1, 0.001},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4654, CurrentL2, 0.001},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4655, CurrentL3, 0.001},
	{0x5100, 0x00465700, 0x004657FF, 0x00, 0x4657, UtilityFrequency, 0.01},
	{0x5100, 0x00491E00, 0x00495DFF, 0x00, 0x495B, BatteryTemperature, 0.1},

	// TODO more decoding for device_status & device_grid_relay
	{0x5180, 0x00214800, 0x002148FF, 0x00, 0x2148, DeviceStatus, 0},
	{0x5180, 0x00416400, 0x004164FF, 0x00, 0x4164, DeviceGridRelay, 0},

	{0x5200, 0x00237700, 0x002377FF, 0x00, 0x2377, DeviceTemperature, 0.01},

	{0x5380, 0x00251E00, 0x00251EFF, 0x01, 0x251E, PowerS1, 0},
	{0x5380, 0x00251E00, 0x00251EFF, 0x02, 0x251E, PowerS2, 0},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x451F, VoltageS1, 0.01},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x451F, VoltageS2, 0.01},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x4521, CurrentS1, 0.001},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x4521, CurrentS2, 0.001},

	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2601, ActiveEnergyPlus, 3600},
	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2622, ActiveEnergyPlusToday, 3600},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462E, TimeOperating, 0},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462F, TimeFeed, 0},

	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821E, DeviceName, 0},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821F, DeviceClass, 0},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x8220, DeviceType, 0},
}

// cache for responses and requests
var (
	// inverterResponseValues map response codes to ValueID
	inverterResponseValues map[uint32]ValueID
	// inverterAllRequests contains a list of InverterValuesDef that is used to get all values
	inverterAllRequests []InverterValuesDef
	// inverterValueMap maps ValueID to InverterValuesDef
	inverterValueMap map[ValueID]InverterValuesDef
)

func init() {
	inverterResponseValues = make(map[uint32]ValueID, len(inverterValues))
	for _, def := range inverterValues {
		inverterResponseValues[uint32(def.Code)<<16+uint32(def.Class)] = def.ID
	}

	inverterAllRequests = GetInverterRequests(inverterValues)

	inverterValueMap = make(map[ValueID]InverterValuesDef, len(inverterValues))
	for _, def := range inverterValues {
		inverterValueMap[def.ID] = def
	}
}

// checkInverterValue checks if response is a known value
func checkInverterValue(value *net2.ResponseValue) ValueID {
	if def, ok := inverterResponseValues[uint32(value.Code)<<16+uint32(value.Class)]; ok {
		return def
	}
	if def, ok := inverterResponseValues[uint32(value.Code)<<16]; ok {
		return def
	}
	return ""
}

// GetAllInverterRequests to receive all values
func GetAllInverterRequests() []InverterValuesDef {
	return inverterAllRequests
}

// GetInverterRequest for given ID
func GetInverterRequest(id ValueID) InverterValuesDef {
	return inverterValueMap[id]
}

// GetInverterRequests to receive all of the given values (reduce request amount)
func GetInverterRequests(values []InverterValuesDef) []InverterValuesDef {
	defs := make([]InverterValuesDef, 0, len(values))

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

// ParseInverterValues from response
func ParseInverterValues(values []*net2.ResponseValue) map[ValueID]interface{} {
	data := make(map[ValueID]interface{}, len(values))

	for _, val := range values {
		if len(val.Values) == 0 {
			continue
		}

		if id := checkInverterValue(val); id != "" {
			value := val.Values[0]
			// handle correction factor
			if inverterValueMap[id].Factor != 0 {
				if v, ok := value.(uint64); ok {
					value = float64(v) * inverterValueMap[id].Factor
				} else if v, ok := value.(uint32); ok {
					value = float64(v) * inverterValueMap[id].Factor
				} else if v, ok := value.(int64); ok {
					value = float64(v) * inverterValueMap[id].Factor
				} else if v, ok := value.(int32); ok {
					value = float64(v) * inverterValueMap[id].Factor
				}
			}
			data[id] = value
		}
	}
	return data
}
