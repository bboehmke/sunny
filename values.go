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

	Key         string
	Description string
}

var valuesDef = []valDef{
	{0x5100, 0x00263F00, 0x00263FFF, 0x00, 0x263F, "power_ac_total", "Total power on AC (W)"},
	{0x5100, 0x00295A00, 0x00295AFF, 0x00, 0x295A, "battery_charge", "Charge state of battery (%)"},
	{0x5100, 0x00411E00, 0x004120FF, 0x00, 0x411E, "power_max", "Maximum possible power (W)"},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4640, "power_ac1", "Power on AC L1 (W)"},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4641, "power_ac2", "Power on AC L2 (W)"},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4642, "power_ac3", "Power on AC L3 (W)"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4648, "voltage_ac1", "Voltage on AC L1 (10 mV)"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4649, "voltage_ac2", "Voltage on AC L2 (10 mV)"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x464a, "voltage_ac3", "Voltage on AC L3 (10 mV)"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4653, "current_ac1", "Current on AC L1 (mA)"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4654, "current_ac2", "Current on AC L2 (mA)"},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4655, "current_ac3", "Current on AC L3 (mA)"},
	{0x5100, 0x00465700, 0x004657FF, 0x00, 0x4657, "frequency_ac", "AC frequency (1/100 Hz)"},
	{0x5100, 0x00491E00, 0x00495DFF, 0x00, 0x495B, "battery_temperature", "Temperature of battery (1/10 °C)"},

	// TODO more decoding for device_status & device_grid_relay
	{0x5180, 0x00214800, 0x002148FF, 0x00, 0x2148, "device_status", "Status of device"},
	{0x5180, 0x00416400, 0x004164FF, 0x00, 0x4164, "device_grid_relay", "Status of grid relay"},

	{0x5200, 0x00237700, 0x002377FF, 0x00, 0x2377, "device_temperature", "Temperature of device (1/10 °C)"},

	{0x5380, 0x00251E00, 0x00251EFF, 0x01, 0x251E, "power_dc1", "Power on DC Line 1 (W)"},
	{0x5380, 0x00251E00, 0x00251EFF, 0x02, 0x251E, "power_dc2", "Power on DC Line 2 (W)"},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x451F, "voltage_dc1", "Voltage on DC Line 1 (10 mV)"},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x451F, "voltage_dc2", "Voltage on DC Line 2 (10 mV)"},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x4521, "current_dc1", "Current on DC Line 1 (mA)"},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x4521, "current_dc2", "Current on DC Line 2 (mA)"},

	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2601, "energy_total", "Energy produced since installation (Wh)"},
	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2622, "energy_today", "Energy produced today (Wh)"},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462E, "time_operating", "Operation time (s)"},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462F, "time_feed", "Feed in time (s)"},

	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821E, "device_name", "Name of device"},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821F, "device_class", "ID of device class"},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x8220, "device_type", "ID of device type"},
}

// emValDef defines information of an energy meter value
type emValDef struct {
	OBIS        string
	Key         string
	Description string
}

var emValuesDef = []emValDef{
	{"0:1.4.0 ", "active_power_plus", "Active Power + (W)"},
	{"0:1.8.0 ", "active_energy_plus", "Active Energy + (Wh)"},
	{"0:2.4.0 ", "active_power_minus", "Active Power - (W)"},
	{"0:2.8.0 ", "active_energy_minus", "Active Energy - (Wh)"},
	{"0:3.4.0 ", "reactive_power_plus", "Reactive Power + (var)"},
	{"0:3.8.0 ", "reactive_energy_plus", "Reactive Energy + (varh)"},
	{"0:4.4.0 ", "reactive_power_minus", "Reactive Power - (var)"},
	{"0:4.8.0 ", "reactive_energy_minus", "Reactive Energy - (varh)"},
	{"0:9.4.0 ", "apparent_power_plus", "Apparent Power + (VA)"},
	{"0:9.8.0 ", "apparent_energy_plus", "Apparent Energy + (VAh)"},
	{"0:10.4.0", "apparent_power_minus", "Apparent Power - (VA)"},
	{"0:10.8.0", "apparent_energy_minus", "Apparent Energy - (VAh)"},
	{"0:13.4.0", "power_factor", "Power Factor (1/1000)"},

	{"0:21.4.0", "l1_active_power_plus", "L1 Active Power + (W)"},
	{"0:21.8.0", "l1_active_energy_plus", "L1 Active Energy + (Wh)"},
	{"0:22.4.0", "l1_active_power_minus", "L1 Active Power - (W)"},
	{"0:22.8.0", "l1_active_energy_minus", "L1 Active Energy - (Wh)"},
	{"0:23.4.0", "l1_reactive_power_plus", "L1 Reactive Power + (var)"},
	{"0:23.8.0", "l1_reactive_energy_plus", "L1 Reactive Energy + (varh)"},
	{"0:24.4.0", "l1_reactive_power_minus", "L1 Reactive Power - (var)"},
	{"0:24.8.0", "l1_reactive_energy_minus", "L1 Reactive Energy - (varh)"},
	{"0:29.4.0", "l1_apparent_power_plus", "L1 Apparent Power + (VA)"},
	{"0:29.8.0", "l1_apparent_energy_plus", "L1 Apparent Energy + (VAh)"},
	{"0:30.4.0", "l1_apparent_power_minus", "L1 Apparent Power - (VA)"},
	{"0:30.8.0", "l1_apparent_energy_minus", "L1 Apparent Energy - (VAh)"},
	{"0:31.4.0", "l1_current", "L1 Current (mA)"},
	{"0:32.4.0", "l1_voltage", "L1 Voltage (mV)"},
	{"0:33.4.0", "l1_power_factor", "L1 Power Factor (1/1000)"},

	{"0:41.4.0", "l2_active_power_plus", "L2 Active Power + (W)"},
	{"0:41.8.0", "l2_active_energy_plus", "L2 Active Energy + (Wh)"},
	{"0:42.4.0", "l2_active_power_minus", "L2 Active Power - (W)"},
	{"0:42.8.0", "l2_active_energy_minus", "L2 Active Energy - (Wh)"},
	{"0:43.4.0", "l2_reactive_power_plus", "L2 Reactive Power + (var)"},
	{"0:43.8.0", "l2_reactive_energy_plus", "L2 Reactive Energy + (varh)"},
	{"0:44.4.0", "l2_reactive_power_minus", "L2 Reactive Power - (var)"},
	{"0:44.8.0", "l2_reactive_energy_minus", "L2 Reactive Energy - (varh)"},
	{"0:49.4.0", "l2_apparent_power_plus", "L2 Apparent Power + (VA)"},
	{"0:49.8.0", "l2_apparent_energy_plus", "L2 Apparent Energy + (VAh)"},
	{"0:50.4.0", "l2_apparent_power_minus", "L2 Apparent Power - (VA)"},
	{"0:50.8.0", "l2_apparent_energy_minus", "L2 Apparent Energy - (VAh)"},
	{"0:51.4.0", "l2_current", "L2 Current (mA)"},
	{"0:52.4.0", "l2_voltage", "L2 Voltage (mV)"},
	{"0:53.4.0", "l2_power_factor", "L2 Power Factor (1/1000)"},

	{"0:61.4.0", "l3_active_power_plus", "L3 Active Power + (W)"},
	{"0:61.8.0", "l3_active_energy_plus", "L3 Active Energy + (Wh)"},
	{"0:62.4.0", "l3_active_power_minus", "L3 Active Power - (W)"},
	{"0:62.8.0", "l3_active_energy_minus", "L3 Active Energy - (Wh)"},
	{"0:63.4.0", "l3_reactive_power_plus", "L3 Reactive Power + (var)"},
	{"0:63.8.0", "l3_reactive_energy_plus", "L3 Reactive Energy + (varh)"},
	{"0:64.4.0", "l3_reactive_power_minus", "L3 Reactive Power - (var)"},
	{"0:64.8.0", "l3_reactive_energy_minus", "L3 Reactive Energy - (varh)"},
	{"0:69.4.0", "l3_apparent_power_plus", "L3 Apparent Power + (VA)"},
	{"0:69.8.0", "l3_apparent_energy_plus", "L3 Apparent Energy + (VAh)"},
	{"0:70.4.0", "l3_apparent_power_minus", "L3 Apparent Power - (VA)"},
	{"0:70.8.0", "l3_apparent_energy_minus", "L3 Apparent Energy - (VAh)"},
	{"0:71.4.0", "l3_current", "L3 Current (mA)"},
	{"0:72.4.0", "l3_voltage", "L3 Voltage (mV)"},
	{"0:73.4.0", "l3_power_factor", "L3 Power Factor (1/1000)"},

	{"144:0.0.0", "software_version", "Software version of energy meter"},
}

// cache for responses and requests
var responseValues map[uint32]string
var allRequests []valDef
var valueMap map[string]valDef

// cache for em values
var emObisMap map[string]emValDef
var emKeyMap map[string]emValDef

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

	valueMap = make(map[string]valDef, len(valuesDef))
	for _, def := range valuesDef {
		valueMap[def.Key] = def
	}

	emObisMap = make(map[string]emValDef, len(emValuesDef))
	for _, def := range emValuesDef {
		emObisMap[def.OBIS] = def
	}
	emKeyMap = make(map[string]emValDef, len(emValuesDef))
	for _, def := range emValuesDef {
		emKeyMap[def.Key] = def
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
	return valueMap[key]
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

// emKeyValues convert energy meter OBIS to key
func emKeyValues(values map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{}, len(values))
	for obis, value := range values {
		def, ok := emObisMap[obis]
		if ok {
			data[def.Key] = value
		} else {
			data[obis] = value
		}
	}
	return data
}
