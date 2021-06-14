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

type ValueKey string

const ActivePowerMax = ValueKey("active_power_max")
const ActivePowerMinus = ValueKey("active_power_minus")
const ActivePowerMinusL1 = ValueKey("active_power_minus_l1")
const ActivePowerMinusL2 = ValueKey("active_power_minus_l2")
const ActivePowerMinusL3 = ValueKey("active_power_minus_l3")
const ActivePowerPlus = ValueKey("active_power_plus")
const ActivePowerPlusL1 = ValueKey("active_power_plus_l1")
const ActivePowerPlusL2 = ValueKey("active_power_plus_l2")
const ActivePowerPlusL3 = ValueKey("active_power_plus_l3")
const ApparentPowerMinus = ValueKey("apparent_power_minus")
const ApparentPowerMinusL1 = ValueKey("apparent_power_minus_l1")
const ApparentPowerMinusL2 = ValueKey("apparent_power_minus_l2")
const ApparentPowerMinusL3 = ValueKey("apparent_power_minus_l3")
const ApparentPowerPlus = ValueKey("apparent_power_plus")
const ApparentPowerPlusL1 = ValueKey("apparent_power_plus_l1")
const ApparentPowerPlusL2 = ValueKey("apparent_power_plus_l2")
const ApparentPowerPlusL3 = ValueKey("apparent_power_plus_l3")
const ReactivePowerMinus = ValueKey("reactive_power_minus")
const ReactivePowerMinusL1 = ValueKey("reactive_power_minus_l1")
const ReactivePowerMinusL2 = ValueKey("reactive_power_minus_l2")
const ReactivePowerMinusL3 = ValueKey("reactive_power_minus_l3")
const ReactivePowerPlus = ValueKey("reactive_power_plus")
const ReactivePowerPlusL1 = ValueKey("reactive_power_plus_l1")
const ReactivePowerPlusL2 = ValueKey("reactive_power_plus_l2")
const ReactivePowerPlusL3 = ValueKey("reactive_power_plus_l3")
const PowerS1 = ValueKey("power_s1")
const PowerS2 = ValueKey("power_s2")
const PowerFactor = ValueKey("power_factor")
const PowerFactorL1 = ValueKey("power_factor_l1")
const PowerFactorL2 = ValueKey("power_factor_l2")
const PowerFactorL3 = ValueKey("power_factor_l3")

const ActiveEnergyMinus = ValueKey("active_energy_minus")
const ActiveEnergyMinusL1 = ValueKey("active_energy_minus_l1")
const ActiveEnergyMinusL2 = ValueKey("active_energy_minus_l2")
const ActiveEnergyMinusL3 = ValueKey("active_energy_minus_l3")
const ActiveEnergyPlus = ValueKey("active_energy_plus")
const ActiveEnergyPlusL1 = ValueKey("active_energy_plus_l1")
const ActiveEnergyPlusL2 = ValueKey("active_energy_plus_l2")
const ActiveEnergyPlusL3 = ValueKey("active_energy_plus_l3")
const ActiveEnergyPlusToday = ValueKey("active_energy_plus_today")
const ApparentEnergyMinus = ValueKey("apparent_energy_minus")
const ApparentEnergyMinusL1 = ValueKey("apparent_energy_minus_l1")
const ApparentEnergyMinusL2 = ValueKey("apparent_energy_minus_l2")
const ApparentEnergyMinusL3 = ValueKey("apparent_energy_minus_l3")
const ApparentEnergyPlus = ValueKey("apparent_energy_plus")
const ApparentEnergyPlusL1 = ValueKey("apparent_energy_plus_l1")
const ApparentEnergyPlusL2 = ValueKey("apparent_energy_plus_l2")
const ApparentEnergyPlusL3 = ValueKey("apparent_energy_plus_l3")
const ReactiveEnergyMinus = ValueKey("reactive_energy_minus")
const ReactiveEnergyMinusL1 = ValueKey("reactive_energy_minus_l1")
const ReactiveEnergyMinusL2 = ValueKey("reactive_energy_minus_l2")
const ReactiveEnergyMinusL3 = ValueKey("reactive_energy_minus_l3")
const ReactiveEnergyPlus = ValueKey("reactive_energy_plus")
const ReactiveEnergyPlusL1 = ValueKey("reactive_energy_plus_l1")
const ReactiveEnergyPlusL2 = ValueKey("reactive_energy_plus_l2")
const ReactiveEnergyPlusL3 = ValueKey("reactive_energy_plus_l3")

const CurrentL1 = ValueKey("current_l1")
const CurrentL2 = ValueKey("current_l2")
const CurrentL3 = ValueKey("current_l3")
const CurrentS1 = ValueKey("current_s1")
const CurrentS2 = ValueKey("current_s2")

const VoltageL1 = ValueKey("voltage_l1")
const VoltageL2 = ValueKey("voltage_l2")
const VoltageL3 = ValueKey("voltage_l3")
const VoltageS1 = ValueKey("voltage_s1")
const VoltageS2 = ValueKey("voltage_s2")

const TimeFeed = ValueKey("time_feed")
const TimeOperating = ValueKey("time_operating")
const UtilityFrequency = ValueKey("utility_frequency")

const BatteryCharge = ValueKey("battery_charge")
const BatteryTemperature = ValueKey("battery_temperature")

const DeviceClass = ValueKey("device_class")
const DeviceGridRelay = ValueKey("device_grid_relay")
const DeviceName = ValueKey("device_name")
const DeviceStatus = ValueKey("device_status")
const DeviceTemperature = ValueKey("device_temperature")
const DeviceType = ValueKey("device_type")

const SoftwareVersion = ValueKey("software_version")

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

	Key    ValueKey
	Factor float64

	Info ValueInfo
}

var valuesDef = []valDef{
	{0x5100, 0x00263F00, 0x00263FFF, 0x00, 0x263F, ActivePowerPlus, 0, ValueInfo{"Total power on AC", "W", "power"}},
	{0x5100, 0x00295A00, 0x00295AFF, 0x00, 0x295A, BatteryCharge, 0, ValueInfo{"Charge state of battery", "%", ""}},
	{0x5100, 0x00411E00, 0x004120FF, 0x00, 0x411E, ActivePowerMax, 0, ValueInfo{"Maximum possible power", "W", "power"}},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4640, ActivePowerPlusL1, 0, ValueInfo{"Power on AC L1", "W", "power"}},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4641, ActivePowerPlusL2, 0, ValueInfo{"Power on AC L2", "W", "power"}},
	{0x5100, 0x00464000, 0x004642FF, 0x00, 0x4642, ActivePowerPlusL3, 0, ValueInfo{"Power on AC L3", "W", "power"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4648, VoltageL1, 0.01, ValueInfo{"Voltage on AC L1", "V", "voltage"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4649, VoltageL2, 0.01, ValueInfo{"Voltage on AC L2", "V", "voltage"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x464a, VoltageL3, 0.01, ValueInfo{"Voltage on AC L3", "V", "voltage"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4653, CurrentL1, 0.001, ValueInfo{"Current on AC L1", "mA", "current"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4654, CurrentL2, 0.001, ValueInfo{"Current on AC L2", "mA", "current"}},
	{0x5100, 0x00464800, 0x004655FF, 0x00, 0x4655, CurrentL3, 0.001, ValueInfo{"Current on AC L3", "mA", "current"}},
	{0x5100, 0x00465700, 0x004657FF, 0x00, 0x4657, UtilityFrequency, 0.01, ValueInfo{"Utility frequency", "Hz", ""}},
	{0x5100, 0x00491E00, 0x00495DFF, 0x00, 0x495B, BatteryTemperature, 0.1, ValueInfo{"Temperature of battery", "°C", "temperature"}},

	// TODO more decoding for device_status & device_grid_relay
	{0x5180, 0x00214800, 0x002148FF, 0x00, 0x2148, DeviceStatus, 0, ValueInfo{"Status of device", "", ""}},
	{0x5180, 0x00416400, 0x004164FF, 0x00, 0x4164, DeviceGridRelay, 0, ValueInfo{"Status of grid relay", "", ""}},

	{0x5200, 0x00237700, 0x002377FF, 0x00, 0x2377, DeviceTemperature, 0.01, ValueInfo{"Temperature of device", "°C", "temperature"}},

	{0x5380, 0x00251E00, 0x00251EFF, 0x01, 0x251E, PowerS1, 0, ValueInfo{"Power on DC Line 1", "W", "power"}},
	{0x5380, 0x00251E00, 0x00251EFF, 0x02, 0x251E, PowerS2, 0, ValueInfo{"Power on DC Line 2", "W", "power"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x451F, VoltageS1, 0.01, ValueInfo{"Voltage on DC Line 1", "V", "voltage"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x451F, VoltageS2, 0.01, ValueInfo{"Voltage on DC Line 2", "V", "voltage"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x4521, CurrentS1, 0.001, ValueInfo{"Current on DC Line 1", "mA", "current"}},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x4521, CurrentS2, 0.001, ValueInfo{"Current on DC Line 2", "mA", "current"}},

	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2601, ActiveEnergyPlus, 0, ValueInfo{"Energy produced since installation", "Wh", "energy"}},
	{0x5400, 0x00260100, 0x002622FF, 0x00, 0x2622, ActiveEnergyPlusToday, 0, ValueInfo{"Energy produced today", "Wh", "energy"}},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462E, TimeOperating, 0, ValueInfo{"Operation time", "s", ""}},
	{0x5400, 0x00462E00, 0x00462FFF, 0x00, 0x462F, TimeFeed, 0, ValueInfo{"Feed in time", "s", ""}},

	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821E, DeviceName, 0, ValueInfo{"Name of device", "", ""}},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x821F, DeviceClass, 0, ValueInfo{"ID of device class", "", ""}},
	{0x5800, 0x00821E00, 0x008220FF, 0x00, 0x8220, DeviceType, 0, ValueInfo{"ID of device type", "", ""}},
}

// emValDef defines information of an energy meter value
type emValDef struct {
	OBIS   string
	Key    ValueKey
	Factor float64

	Info ValueInfo
}

var emValuesDef = []emValDef{
	{"0:1.4.0", ActivePowerPlus, 0.1, ValueInfo{"Active Power +", "W", "power"}},
	{"0:1.8.0", ActiveEnergyPlus, 0, ValueInfo{"Active Energy +", "Ws", "energy"}},
	{"0:2.4.0", ActivePowerMinus, 0.1, ValueInfo{"Active Power -", "W", "power"}},
	{"0:2.8.0", ActiveEnergyMinus, 0, ValueInfo{"Active Energy -", "Ws", "energy"}},
	{"0:3.4.0", ReactivePowerPlus, 0.1, ValueInfo{"Reactive Power +", "var", "power"}},
	{"0:3.8.0", ReactiveEnergyPlus, 0, ValueInfo{"Reactive Energy +", "vars", "energy"}},
	{"0:4.4.0", ReactivePowerMinus, 0.1, ValueInfo{"Reactive Power -", "var", "power"}},
	{"0:4.8.0", ReactiveEnergyMinus, 0, ValueInfo{"Reactive Energy -", "vars", "energy"}},
	{"0:9.4.0", ApparentPowerPlus, 0.1, ValueInfo{"Apparent Power +", "VA", "power"}},
	{"0:9.8.0", ApparentEnergyPlus, 0, ValueInfo{"Apparent Energy +", "VAs", "energy"}},
	{"0:10.4.0", ApparentPowerMinus, 0.1, ValueInfo{"Apparent Power -", "VA", "power"}},
	{"0:10.8.0", ApparentEnergyMinus, 0, ValueInfo{"Apparent Energy -", "VAs", "energy"}},
	{"0:13.4.0", PowerFactor, 0.001, ValueInfo{"Power Factor", "", ""}},
	{"0:14.4.0", UtilityFrequency, 0.001, ValueInfo{"utility Frequency", "Hz", ""}},

	{"0:21.4.0", ActivePowerPlusL1, 0.1, ValueInfo{"L1 Active Power +", "W", "power"}},
	{"0:21.8.0", ActiveEnergyPlusL1, 0, ValueInfo{"L1 Active Energy +", "Ws", "energy"}},
	{"0:22.4.0", ActivePowerMinusL1, 0.1, ValueInfo{"L1 Active Power -", "W", "power"}},
	{"0:22.8.0", ActiveEnergyMinusL1, 0, ValueInfo{"L1 Active Energy -", "Ws", "energy"}},
	{"0:23.4.0", ReactivePowerPlusL1, 0.1, ValueInfo{"L1 Reactive Power +", "var", "power"}},
	{"0:23.8.0", ReactiveEnergyPlusL1, 0, ValueInfo{"L1 Reactive Energy +", "vars", "energy"}},
	{"0:24.4.0", ReactivePowerMinusL1, 0.1, ValueInfo{"L1 Reactive Power -", "var", "power"}},
	{"0:24.8.0", ReactiveEnergyMinusL1, 0, ValueInfo{"L1 Reactive Energy -", "vars", "energy"}},
	{"0:29.4.0", ApparentPowerPlusL1, 0.1, ValueInfo{"L1 Apparent Power +", "VA", "power"}},
	{"0:29.8.0", ApparentEnergyPlusL1, 0, ValueInfo{"L1 Apparent Energy +", "VAs", "energy"}},
	{"0:30.4.0", ApparentPowerMinusL1, 0.1, ValueInfo{"L1 Apparent Power -", "VA", "power"}},
	{"0:30.8.0", ApparentEnergyMinusL1, 0, ValueInfo{"L1 Apparent Energy -", "VAs", "energy"}},
	{"0:31.4.0", CurrentL1, 0.001, ValueInfo{"L1 Current", "A", "current"}},
	{"0:32.4.0", VoltageL1, 0.001, ValueInfo{"L1 Voltage", "V", "voltage"}},
	{"0:33.4.0", PowerFactorL1, 0.001, ValueInfo{"L1 Power Factor", "", ""}},

	{"0:41.4.0", ActivePowerPlusL2, 0.1, ValueInfo{"L2 Active Power +", "W", "power"}},
	{"0:41.8.0", ActiveEnergyPlusL2, 0, ValueInfo{"L2 Active Energy +", "Ws", "energy"}},
	{"0:42.4.0", ActivePowerMinusL2, 0.1, ValueInfo{"L2 Active Power -", "W", "power"}},
	{"0:42.8.0", ActiveEnergyMinusL2, 0, ValueInfo{"L2 Active Energy -", "Ws", "energy"}},
	{"0:43.4.0", ReactivePowerPlusL2, 0.1, ValueInfo{"L2 Reactive Power +", "var", "power"}},
	{"0:43.8.0", ReactiveEnergyPlusL2, 0, ValueInfo{"L2 Reactive Energy +", "vars", "energy"}},
	{"0:44.4.0", ReactivePowerMinusL2, 0.1, ValueInfo{"L2 Reactive Power -", "var", "power"}},
	{"0:44.8.0", ReactiveEnergyMinusL2, 0, ValueInfo{"L2 Reactive Energy -", "vars", "energy"}},
	{"0:49.4.0", ApparentPowerPlusL2, 0.1, ValueInfo{"L2 Apparent Power +", "VA", "power"}},
	{"0:49.8.0", ApparentEnergyPlusL2, 0, ValueInfo{"L2 Apparent Energy +", "VAs", "energy"}},
	{"0:50.4.0", ApparentPowerMinusL2, 0.1, ValueInfo{"L2 Apparent Power -", "VA", "power"}},
	{"0:50.8.0", ApparentEnergyMinusL2, 0, ValueInfo{"L2 Apparent Energy -", "VAs", "energy"}},
	{"0:51.4.0", CurrentL2, 0.001, ValueInfo{"L2 Current", "A", "current"}},
	{"0:52.4.0", VoltageL2, 0.001, ValueInfo{"L2 Voltage", "V", "voltage"}},
	{"0:53.4.0", PowerFactorL2, 0.001, ValueInfo{"L2 Power Factor", "", ""}},

	{"0:61.4.0", ActivePowerPlusL3, 0.1, ValueInfo{"L3 Active Power +", "W", "power"}},
	{"0:61.8.0", ActiveEnergyPlusL3, 0, ValueInfo{"L3 Active Energy +", "Ws", "energy"}},
	{"0:62.4.0", ActivePowerMinusL3, 0.1, ValueInfo{"L3 Active Power -", "W", "power"}},
	{"0:62.8.0", ActiveEnergyMinusL3, 0, ValueInfo{"L3 Active Energy -", "Ws", "energy"}},
	{"0:63.4.0", ReactivePowerPlusL3, 0.1, ValueInfo{"L3 Reactive Power +", "var", "power"}},
	{"0:63.8.0", ReactiveEnergyPlusL3, 0, ValueInfo{"L3 Reactive Energy +", "vars", "energy"}},
	{"0:64.4.0", ReactivePowerMinusL3, 0.1, ValueInfo{"L3 Reactive Power -", "var", "power"}},
	{"0:64.8.0", ReactiveEnergyMinusL3, 0, ValueInfo{"L3 Reactive Energy -", "vars", "energy"}},
	{"0:69.4.0", ApparentPowerPlusL3, 0.1, ValueInfo{"L3 Apparent Power +", "VA", "power"}},
	{"0:69.8.0", ApparentEnergyPlusL3, 0, ValueInfo{"L3 Apparent Energy +", "VAs", "energy"}},
	{"0:70.4.0", ApparentPowerMinusL3, 0.1, ValueInfo{"L3 Apparent Power -", "VA", "power"}},
	{"0:70.8.0", ApparentEnergyMinusL3, 0, ValueInfo{"L3 Apparent Energy -", "VAs", "energy"}},
	{"0:71.4.0", CurrentL3, 0.001, ValueInfo{"L3 Current", "A", "current"}},
	{"0:72.4.0", VoltageL3, 0.001, ValueInfo{"L3 Voltage", "V", "voltage"}},
	{"0:73.4.0", PowerFactorL3, 0.001, ValueInfo{"L3 Power Factor", "", ""}},

	{"144:0.0.0", SoftwareVersion, 0, ValueInfo{"Software version of energy meter", "", ""}},
}

// cache for responses and requests
var responseValues map[uint32]ValueKey
var allRequests []valDef
var valueMap map[ValueKey]valDef

// cache for em values
var emObisMap map[string]emValDef
var emKeyMap map[ValueKey]emValDef

// init cache
func init() {
	responseValues = make(map[uint32]ValueKey, len(valuesDef))
	for _, def := range valuesDef {
		responseValues[uint32(def.Code)<<16+uint32(def.Class)] = def.Key
	}

	allRequests = getRequests(valuesDef)

	valueMap = make(map[ValueKey]valDef, len(valuesDef))
	for _, def := range valuesDef {
		valueMap[def.Key] = def
	}

	emObisMap = make(map[string]emValDef, len(emValuesDef))
	for _, def := range emValuesDef {
		emObisMap[def.OBIS] = def
	}
	emKeyMap = make(map[ValueKey]emValDef, len(emValuesDef))
	for _, def := range emValuesDef {
		emKeyMap[def.Key] = def
	}

	Log.Printf("valuesDef: %d", len(valuesDef))
	Log.Printf("responseValues: %d", len(valuesDef))
	Log.Printf("allRequests: %d", len(allRequests))
	Log.Printf("emValuesDef: %d", len(emValuesDef))
}

// checkValue checks if response is a known value
func checkValue(value *net2.ResponseValue) ValueKey {
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
func getRequest(key ValueKey) valDef {
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
func parseValues(values []*net2.ResponseValue) map[ValueKey]interface{} {
	data := make(map[ValueKey]interface{}, len(values))

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
func emKeyValues(values map[string]interface{}) map[ValueKey]interface{} {
	data := make(map[ValueKey]interface{}, len(values))
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
			data[ValueKey(obis)] = value
		}
	}
	return data
}
