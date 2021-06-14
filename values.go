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

const ValueActivePowerMax = "active_power_max"
const ValueActivePowerMinus = "active_power_minus"
const ValueActivePowerMinusL1 = "active_power_minus_l1"
const ValueActivePowerMinusL2 = "active_power_minus_l2"
const ValueActivePowerMinusL3 = "active_power_minus_l3"
const ValueActivePowerPlus = "active_power_plus"
const ValueActivePowerPlusL1 = "active_power_plus_l1"
const ValueActivePowerPlusL2 = "active_power_plus_l2"
const ValueActivePowerPlusL3 = "active_power_plus_l3"
const ValueApparentPowerMinus = "apparent_power_minus"
const ValueApparentPowerMinusL1 = "apparent_power_minus_l1"
const ValueApparentPowerMinusL2 = "apparent_power_minus_l2"
const ValueApparentPowerMinusL3 = "apparent_power_minus_l3"
const ValueApparentPowerPlus = "apparent_power_plus"
const ValueApparentPowerPlusL1 = "apparent_power_plus_l1"
const ValueApparentPowerPlusL2 = "apparent_power_plus_l2"
const ValueApparentPowerPlusL3 = "apparent_power_plus_l3"
const ValueReactivePowerMinus = "reactive_power_minus"
const ValueReactivePowerMinusL1 = "reactive_power_minus_l1"
const ValueReactivePowerMinusL2 = "reactive_power_minus_l2"
const ValueReactivePowerMinusL3 = "reactive_power_minus_l3"
const ValueReactivePowerPlus = "reactive_power_plus"
const ValueReactivePowerPlusL1 = "reactive_power_plus_l1"
const ValueReactivePowerPlusL2 = "reactive_power_plus_l2"
const ValueReactivePowerPlusL3 = "reactive_power_plus_l3"
const ValuePowerS1 = "power_s1"
const ValuePowerS2 = "power_s2"
const ValuePowerFactor = "power_factor"
const ValuePowerFactorL1 = "power_factor_l1"
const ValuePowerFactorL2 = "power_factor_l2"
const ValuePowerFactorL3 = "power_factor_l3"

const ValueActiveEnergyMinus = "active_energy_minus"
const ValueActiveEnergyMinusL1 = "active_energy_minus_l1"
const ValueActiveEnergyMinusL2 = "active_energy_minus_l2"
const ValueActiveEnergyMinusL3 = "active_energy_minus_l3"
const ValueActiveEnergyPlus = "active_energy_plus"
const ValueActiveEnergyPlusL1 = "active_energy_plus_l1"
const ValueActiveEnergyPlusL2 = "active_energy_plus_l2"
const ValueActiveEnergyPlusL3 = "active_energy_plus_l3"
const ValueActiveEnergyPlusToday = "active_energy_plus_today"
const ValueApparentEnergyMinus = "apparent_energy_minus"
const ValueApparentEnergyMinusL1 = "apparent_energy_minus_l1"
const ValueApparentEnergyMinusL2 = "apparent_energy_minus_l2"
const ValueApparentEnergyMinusL3 = "apparent_energy_minus_l3"
const ValueApparentEnergyPlus = "apparent_energy_plus"
const ValueApparentEnergyPlusL1 = "apparent_energy_plus_l1"
const ValueApparentEnergyPlusL2 = "apparent_energy_plus_l2"
const ValueApparentEnergyPlusL3 = "apparent_energy_plus_l3"
const ValueReactiveEnergyMinus = "reactive_energy_minus"
const ValueReactiveEnergyMinusL1 = "reactive_energy_minus_l1"
const ValueReactiveEnergyMinusL2 = "reactive_energy_minus_l2"
const ValueReactiveEnergyMinusL3 = "reactive_energy_minus_l3"
const ValueReactiveEnergyPlus = "reactive_energy_plus"
const ValueReactiveEnergyPlusL1 = "reactive_energy_plus_l1"
const ValueReactiveEnergyPlusL2 = "reactive_energy_plus_l2"
const ValueReactiveEnergyPlusL3 = "reactive_energy_plus_l3"

const ValueCurrentL1 = "current_l1"
const ValueCurrentL2 = "current_l2"
const ValueCurrentL3 = "current_l3"
const ValueCurrentS1 = "current_s1"
const ValueCurrentS2 = "current_s2"

const ValueVoltageL1 = "voltage_l1"
const ValueVoltageL2 = "voltage_l2"
const ValueVoltageL3 = "voltage_l3"
const ValueVoltageS1 = "voltage_s1"
const ValueVoltageS2 = "voltage_s2"

const ValueTimeFeed = "time_feed"
const ValueTimeOperating = "time_operating"
const ValueUtilityFrequency = "utility_frequency"

const ValueBatteryCharge = "battery_charge"
const ValueBatteryTemperature = "battery_temperature"

const ValueDeviceClass = "device_class"
const ValueDeviceGridRelay = "device_grid_relay"
const ValueDeviceName = "device_name"
const ValueDeviceStatus = "device_status"
const ValueDeviceTemperature = "device_temperature"
const ValueDeviceType = "device_type"

const ValueSoftwareVersion = "software_version"

// ValueInfo describes a values
type ValueInfo struct {
	Description string
	Unit        string
	Type        string
}

// valDef defines a value of an inverter device
type valDef struct {
	Object uint16
	Start  uint32
	End    uint32

	Class uint8 // 0 -> ignore class
	Code  uint16

	Key    string
	Factor float64

	Info ValueInfo
}

var valuesDef = []valDef{
	{0x5100, 0x00263F00, 0x00263FFF, 0x00, 0x263F, ValueActivePowerPlus, 0, ValueInfo{"Total power on AC", "W", "power"}},
	{0x5100, 0x00295A00, 0x00295AFF, 0x00, 0x295A, ValueBatteryCharge, 0, ValueInfo{"Charge state of battery", "%", ""}},
	{0x5100, 0x00411E00, 0x004120FF, 0x00, 0x411E, ValueActivePowerMax, 0, ValueInfo{"Maximum possible power", "W", "power"}},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4640, ValueActivePowerPlusL1, 0, ValueInfo{"Power on AC L1", "W", "power"}},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4641, ValueActivePowerPlusL2, 0, ValueInfo{"Power on AC L2", "W", "power"}},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4642, ValueActivePowerPlusL3, 0, ValueInfo{"Power on AC L3", "W", "power"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4648, ValueVoltageL1, 0.01, ValueInfo{"Voltage on AC L1", "V", "voltage"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4649, ValueVoltageL2, 0.01, ValueInfo{"Voltage on AC L2", "V", "voltage"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x464a, ValueVoltageL3, 0.01, ValueInfo{"Voltage on AC L3", "V", "voltage"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4653, ValueCurrentL1, 0.001, ValueInfo{"Current on AC L1", "mA", "current"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4654, ValueCurrentL2, 0.001, ValueInfo{"Current on AC L2", "mA", "current"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4655, ValueCurrentL3, 0.001, ValueInfo{"Current on AC L3", "mA", "current"}},
	{0x5100, 0x00465700, 0x004657FF, 0x00, 0x4657, ValueUtilityFrequency, 0.01, ValueInfo{"Utility frequency", "Hz", ""}},
	{0x5100, 0x00491E00, 0x00495DFF, 0x00, 0x495B, ValueBatteryTemperature, 0.1, ValueInfo{"Temperature of battery", "°C", "temperature"}},

	// TODO more decoding for device_status & device_grid_relay
	{0x5180, 0x00214800, 0x002148FF, 0x00, 0x2148, ValueDeviceStatus, 0, ValueInfo{"Status of device", "", ""}},
	{0x5180, 0x00416400, 0x004164FF, 0x00, 0x4164, ValueDeviceGridRelay, 0, ValueInfo{"Status of grid relay", "", ""}},

	{0x5200, 0x00237700, 0x002377FF, 0x00, 0x2377, ValueDeviceTemperature, 0.01, ValueInfo{"Temperature of device", "°C", "temperature"}},

	{0x5380, 0x00251E00, 0x00251EFF, 0x01, 0x251E, ValuePowerS1, 0, ValueInfo{"Power on DC Line 1", "W", "power"}},
	{0x5380, 0x00251E00, 0x00251EFF, 0x02, 0x251E, ValuePowerS2, 0, ValueInfo{"Power on DC Line 2", "W", "power"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x451F, ValueVoltageS1, 0.01, ValueInfo{"Voltage on DC Line 1", "V", "voltage"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x451F, ValueVoltageS2, 0.01, ValueInfo{"Voltage on DC Line 2", "V", "voltage"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x4521, ValueCurrentS1, 0.001, ValueInfo{"Current on DC Line 1", "mA", "current"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x4521, ValueCurrentS2, 0.001, ValueInfo{"Current on DC Line 2", "mA", "current"}},

	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2601, ValueActiveEnergyPlus, 0, ValueInfo{"Energy produced since installation", "Wh", "energy"}},
	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2622, ValueActiveEnergyPlusToday, 0, ValueInfo{"Energy produced today", "Wh", "energy"}},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462E, ValueTimeOperating, 0, ValueInfo{"Operation time", "s", ""}},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462F, ValueTimeFeed, 0, ValueInfo{"Feed in time", "s", ""}},

	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821E, ValueDeviceName, 0, ValueInfo{"Name of device", "", ""}},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821F, ValueDeviceClass, 0, ValueInfo{"ID of device class", "", ""}},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x8220, ValueDeviceType, 0, ValueInfo{"ID of device type", "", ""}},
}

// emValDef defines information of an energy meter value
type emValDef struct {
	OBIS   string
	Key    string
	Factor float64

	Info ValueInfo
}

var emValuesDef = []emValDef{
	{"0:1.4.0", ValueActivePowerPlus, 0.1, ValueInfo{"Active Power +", "W", "power"}},
	{"0:1.8.0", ValueActiveEnergyPlus, 0, ValueInfo{"Active Energy +", "Ws", "energy"}},
	{"0:2.4.0", ValueActivePowerMinus, 0.1, ValueInfo{"Active Power -", "W", "power"}},
	{"0:2.8.0", ValueActiveEnergyMinus, 0, ValueInfo{"Active Energy -", "Ws", "energy"}},
	{"0:3.4.0", ValueReactivePowerPlus, 0.1, ValueInfo{"Reactive Power +", "var", "power"}},
	{"0:3.8.0", ValueReactiveEnergyPlus, 0, ValueInfo{"Reactive Energy +", "vars", "energy"}},
	{"0:4.4.0", ValueReactivePowerMinus, 0.1, ValueInfo{"Reactive Power -", "var", "power"}},
	{"0:4.8.0", ValueReactiveEnergyMinus, 0, ValueInfo{"Reactive Energy -", "vars", "energy"}},
	{"0:9.4.0", ValueApparentPowerPlus, 0.1, ValueInfo{"Apparent Power +", "VA", "power"}},
	{"0:9.8.0", ValueApparentEnergyPlus, 0, ValueInfo{"Apparent Energy +", "VAs", "energy"}},
	{"0:10.4.0", ValueApparentPowerMinus, 0.1, ValueInfo{"Apparent Power -", "VA", "power"}},
	{"0:10.8.0", ValueApparentEnergyMinus, 0, ValueInfo{"Apparent Energy -", "VAs", "energy"}},
	{"0:13.4.0", ValuePowerFactor, 0.001, ValueInfo{"Power Factor", "", ""}},
	{"0:14.4.0", ValueUtilityFrequency, 0.001, ValueInfo{"utility Frequency", "Hz", ""}},

	{"0:21.4.0", ValueActivePowerPlusL1, 0.1, ValueInfo{"L1 Active Power +", "W", "power"}},
	{"0:21.8.0", ValueActiveEnergyPlusL1, 0, ValueInfo{"L1 Active Energy +", "Ws", "energy"}},
	{"0:22.4.0", ValueActivePowerMinusL1, 0.1, ValueInfo{"L1 Active Power -", "W", "power"}},
	{"0:22.8.0", ValueActiveEnergyMinusL1, 0, ValueInfo{"L1 Active Energy -", "Ws", "energy"}},
	{"0:23.4.0", ValueReactivePowerPlusL1, 0.1, ValueInfo{"L1 Reactive Power +", "var", "power"}},
	{"0:23.8.0", ValueReactiveEnergyPlusL1, 0, ValueInfo{"L1 Reactive Energy +", "vars", "energy"}},
	{"0:24.4.0", ValueReactivePowerMinusL1, 0.1, ValueInfo{"L1 Reactive Power -", "var", "power"}},
	{"0:24.8.0", ValueReactiveEnergyMinusL1, 0, ValueInfo{"L1 Reactive Energy -", "vars", "energy"}},
	{"0:29.4.0", ValueApparentPowerPlusL1, 0.1, ValueInfo{"L1 Apparent Power +", "VA", "power"}},
	{"0:29.8.0", ValueApparentEnergyPlusL1, 0, ValueInfo{"L1 Apparent Energy +", "VAs", "energy"}},
	{"0:30.4.0", ValueApparentPowerMinusL1, 0.1, ValueInfo{"L1 Apparent Power -", "VA", "power"}},
	{"0:30.8.0", ValueApparentEnergyMinusL1, 0, ValueInfo{"L1 Apparent Energy -", "VAs", "energy"}},
	{"0:31.4.0", ValueCurrentL1, 0.001, ValueInfo{"L1 Current", "A", "current"}},
	{"0:32.4.0", ValueVoltageL1, 0.001, ValueInfo{"L1 Voltage", "V", "voltage"}},
	{"0:33.4.0", ValuePowerFactorL1, 0.001, ValueInfo{"L1 Power Factor", "", ""}},

	{"0:41.4.0", ValueActivePowerPlusL2, 0.1, ValueInfo{"L2 Active Power +", "W", "power"}},
	{"0:41.8.0", ValueActiveEnergyPlusL2, 0, ValueInfo{"L2 Active Energy +", "Ws", "energy"}},
	{"0:42.4.0", ValueActivePowerMinusL2, 0.1, ValueInfo{"L2 Active Power -", "W", "power"}},
	{"0:42.8.0", ValueActiveEnergyMinusL2, 0, ValueInfo{"L2 Active Energy -", "Ws", "energy"}},
	{"0:43.4.0", ValueReactivePowerPlusL2, 0.1, ValueInfo{"L2 Reactive Power +", "var", "power"}},
	{"0:43.8.0", ValueReactiveEnergyPlusL2, 0, ValueInfo{"L2 Reactive Energy +", "vars", "energy"}},
	{"0:44.4.0", ValueReactivePowerMinusL2, 0.1, ValueInfo{"L2 Reactive Power -", "var", "power"}},
	{"0:44.8.0", ValueReactiveEnergyMinusL2, 0, ValueInfo{"L2 Reactive Energy -", "vars", "energy"}},
	{"0:49.4.0", ValueApparentPowerPlusL2, 0.1, ValueInfo{"L2 Apparent Power +", "VA", "power"}},
	{"0:49.8.0", ValueApparentEnergyPlusL2, 0, ValueInfo{"L2 Apparent Energy +", "VAs", "energy"}},
	{"0:50.4.0", ValueApparentPowerMinusL2, 0.1, ValueInfo{"L2 Apparent Power -", "VA", "power"}},
	{"0:50.8.0", ValueApparentEnergyMinusL2, 0, ValueInfo{"L2 Apparent Energy -", "VAs", "energy"}},
	{"0:51.4.0", ValueCurrentL2, 0.001, ValueInfo{"L2 Current", "A", "current"}},
	{"0:52.4.0", ValueVoltageL2, 0.001, ValueInfo{"L2 Voltage", "V", "voltage"}},
	{"0:53.4.0", ValuePowerFactorL2, 0.001, ValueInfo{"L2 Power Factor", "", ""}},

	{"0:61.4.0", ValueActivePowerPlusL3, 0.1, ValueInfo{"L3 Active Power +", "W", "power"}},
	{"0:61.8.0", ValueActiveEnergyPlusL3, 0, ValueInfo{"L3 Active Energy +", "Ws", "energy"}},
	{"0:62.4.0", ValueActivePowerMinusL3, 0.1, ValueInfo{"L3 Active Power -", "W", "power"}},
	{"0:62.8.0", ValueActiveEnergyMinusL3, 0, ValueInfo{"L3 Active Energy -", "Ws", "energy"}},
	{"0:63.4.0", ValueReactivePowerPlusL3, 0.1, ValueInfo{"L3 Reactive Power +", "var", "power"}},
	{"0:63.8.0", ValueReactiveEnergyPlusL3, 0, ValueInfo{"L3 Reactive Energy +", "vars", "energy"}},
	{"0:64.4.0", ValueReactivePowerMinusL3, 0.1, ValueInfo{"L3 Reactive Power -", "var", "power"}},
	{"0:64.8.0", ValueReactiveEnergyMinusL3, 0, ValueInfo{"L3 Reactive Energy -", "vars", "energy"}},
	{"0:69.4.0", ValueApparentPowerPlusL3, 0.1, ValueInfo{"L3 Apparent Power +", "VA", "power"}},
	{"0:69.8.0", ValueApparentEnergyPlusL3, 0, ValueInfo{"L3 Apparent Energy +", "VAs", "energy"}},
	{"0:70.4.0", ValueApparentPowerMinusL3, 0.1, ValueInfo{"L3 Apparent Power -", "VA", "power"}},
	{"0:70.8.0", ValueApparentEnergyMinusL3, 0, ValueInfo{"L3 Apparent Energy -", "VAs", "energy"}},
	{"0:71.4.0", ValueCurrentL3, 0.001, ValueInfo{"L3 Current", "A", "current"}},
	{"0:72.4.0", ValueVoltageL3, 0.001, ValueInfo{"L3 Voltage", "V", "voltage"}},
	{"0:73.4.0", ValuePowerFactorL3, 0.001, ValueInfo{"L3 Power Factor", "", ""}},

	{"144:0.0.0", ValueSoftwareVersion, 0, ValueInfo{"Software version of energy meter", "", ""}},
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

	Log.Printf("valuesDef: %d", len(valuesDef))
	Log.Printf("responseValues: %d", len(valuesDef))
	Log.Printf("allRequests: %d", len(allRequests))
	Log.Printf("emValuesDef: %d", len(emValuesDef))
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
			value := val.Values[0]
			// handle correction factor
			if valueMap[key].Factor != 0 {
				if v, ok := value.(uint64); ok {
					value = float64(v) * valueMap[key].Factor
				} else if v, ok := value.(uint32); ok {
					value = float64(v) * valueMap[key].Factor
				} else if v, ok := value.(int64); ok {
					value = float64(v) * valueMap[key].Factor
				} else if v, ok := value.(int32); ok {
					value = float64(v) * valueMap[key].Factor
				}
			}
			data[key] = value
		}
	}
	return data
}

// emKeyValues convert energy meter OBIS to key
func emKeyValues(values map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{}, len(values))
	for obis, value := range values {
		if def, ok := emObisMap[obis]; ok {
			// handle correction factor
			if def.Factor != 0 {
				if v, ok := value.(uint64); ok {
					value = float64(v) * def.Factor
				} else if v, ok := value.(uint32); ok {
					value = float64(v) * def.Factor
				}
			}
			data[def.Key] = value
		} else {
			data[obis] = value
		}
	}
	return data
}
