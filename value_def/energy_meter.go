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

// EnergyMeterValuesDef defines a value of an energy meter value
type EnergyMeterValuesDef struct {
	OBIS   string
	ID     ValueID
	Factor float64
}

// emValues contains all values that can be read from energy meter
var emValues = []EnergyMeterValuesDef{
	{"0:1.4.0", ActivePowerPlus, 0.1},
	{"0:1.8.0", ActiveEnergyPlus, 0},
	{"0:2.4.0", ActivePowerMinus, 0.1},
	{"0:2.8.0", ActiveEnergyMinus, 0},
	{"0:3.4.0", ReactivePowerPlus, 0.1},
	{"0:3.8.0", ReactiveEnergyPlus, 0},
	{"0:4.4.0", ReactivePowerMinus, 0.1},
	{"0:4.8.0", ReactiveEnergyMinus, 0},
	{"0:9.4.0", ApparentPowerPlus, 0.1},
	{"0:9.8.0", ApparentEnergyPlus, 0},
	{"0:10.4.0", ApparentPowerMinus, 0.1},
	{"0:10.8.0", ApparentEnergyMinus, 0},
	{"0:13.4.0", PowerFactor, 0.001},
	{"0:14.4.0", UtilityFrequency, 0.001},

	{"0:21.4.0", ActivePowerPlusL1, 0.1},
	{"0:21.8.0", ActiveEnergyPlusL1, 0},
	{"0:22.4.0", ActivePowerMinusL1, 0.1},
	{"0:22.8.0", ActiveEnergyMinusL1, 0},
	{"0:23.4.0", ReactivePowerPlusL1, 0.1},
	{"0:23.8.0", ReactiveEnergyPlusL1, 0},
	{"0:24.4.0", ReactivePowerMinusL1, 0.1},
	{"0:24.8.0", ReactiveEnergyMinusL1, 0},
	{"0:29.4.0", ApparentPowerPlusL1, 0.1},
	{"0:29.8.0", ApparentEnergyPlusL1, 0},
	{"0:30.4.0", ApparentPowerMinusL1, 0.1},
	{"0:30.8.0", ApparentEnergyMinusL1, 0},
	{"0:31.4.0", CurrentL1, 0.001},
	{"0:32.4.0", VoltageL1, 0.001},
	{"0:33.4.0", PowerFactorL1, 0.001},

	{"0:41.4.0", ActivePowerPlusL2, 0.1},
	{"0:41.8.0", ActiveEnergyPlusL2, 0},
	{"0:42.4.0", ActivePowerMinusL2, 0.1},
	{"0:42.8.0", ActiveEnergyMinusL2, 0},
	{"0:43.4.0", ReactivePowerPlusL2, 0.1},
	{"0:43.8.0", ReactiveEnergyPlusL2, 0},
	{"0:44.4.0", ReactivePowerMinusL2, 0.1},
	{"0:44.8.0", ReactiveEnergyMinusL2, 0},
	{"0:49.4.0", ApparentPowerPlusL2, 0.1},
	{"0:49.8.0", ApparentEnergyPlusL2, 0},
	{"0:50.4.0", ApparentPowerMinusL2, 0.1},
	{"0:50.8.0", ApparentEnergyMinusL2, 0},
	{"0:51.4.0", CurrentL2, 0.001},
	{"0:52.4.0", VoltageL2, 0.001},
	{"0:53.4.0", PowerFactorL2, 0.001},

	{"0:61.4.0", ActivePowerPlusL3, 0.1},
	{"0:61.8.0", ActiveEnergyPlusL3, 0},
	{"0:62.4.0", ActivePowerMinusL3, 0.1},
	{"0:62.8.0", ActiveEnergyMinusL3, 0},
	{"0:63.4.0", ReactivePowerPlusL3, 0.1},
	{"0:63.8.0", ReactiveEnergyPlusL3, 0},
	{"0:64.4.0", ReactivePowerMinusL3, 0.1},
	{"0:64.8.0", ReactiveEnergyMinusL3, 0},
	{"0:69.4.0", ApparentPowerPlusL3, 0.1},
	{"0:69.8.0", ApparentEnergyPlusL3, 0},
	{"0:70.4.0", ApparentPowerMinusL3, 0.1},
	{"0:70.8.0", ApparentEnergyMinusL3, 0},
	{"0:71.4.0", CurrentL3, 0.001},
	{"0:72.4.0", VoltageL3, 0.001},
	{"0:73.4.0", PowerFactorL3, 0.001},

	{"144:0.0.0", SoftwareVersion, 0},
}

// cache for em values
var (
	// emObisMap maps OBIS to EnergyMeterValuesDef
	emObisMap map[string]EnergyMeterValuesDef
	// emIdMap maps ValueID to EnergyMeterValuesDef
	emIdMap map[ValueID]EnergyMeterValuesDef
)

func init() {
	emObisMap = make(map[string]EnergyMeterValuesDef, len(emValues))
	for _, def := range emValues {
		emObisMap[def.OBIS] = def
	}
	emIdMap = make(map[ValueID]EnergyMeterValuesDef, len(emValues))
	for _, def := range emValues {
		emIdMap[def.ID] = def
	}
}

// ConvertEnergyMeterValues from OBIS to ID based map
func ConvertEnergyMeterValues(values map[string]interface{}) map[ValueID]interface{} {
	data := make(map[ValueID]interface{}, len(values))
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
			data[def.ID] = value
		} else {
			data[ValueID(obis)] = value
		}
	}
	return data
}
