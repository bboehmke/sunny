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

import "gitlab.com/bboehmke/sunny/proto/net2"

//go:generate go run github.com/dmarkham/enumer -type ValueID -output values_enumer.go

// ValueID identifies a value read from a device
type ValueID int

const (
	// ActivePowerMax Maximum active power (AC)
	ActivePowerMax ValueID = iota + 1
	// ActivePowerMinus Active power - (AC)
	ActivePowerMinus
	// ActivePowerMinusL1 Active power - L1 (AC)
	ActivePowerMinusL1
	// ActivePowerMinusL2 Active power - L2 (AC)
	ActivePowerMinusL2
	// ActivePowerMinusL3 Active power - L3 (AC)
	ActivePowerMinusL3
	// ActivePowerPlus Active power + (AC)
	ActivePowerPlus
	// ActivePowerPlusL1 Active power + L1 (AC)
	ActivePowerPlusL1
	// ActivePowerPlusL2 Active power + L2 (AC)
	ActivePowerPlusL2
	// ActivePowerPlusL3 Active power + L3 (AC)
	ActivePowerPlusL3
	// ApparentPowerMinus Apparent power - (AC)
	ApparentPowerMinus
	// ApparentPowerMinusL1 Apparent power - L1 (AC)
	ApparentPowerMinusL1
	// ApparentPowerMinusL2 Apparent power - L2 (AC)
	ApparentPowerMinusL2
	// ApparentPowerMinusL3 Apparent power - L3 (AC)
	ApparentPowerMinusL3
	// ApparentPowerPlus Apparent power + (AC)
	ApparentPowerPlus
	// ApparentPowerPlusL1 Apparent power + L1 (AC)
	ApparentPowerPlusL1
	// ApparentPowerPlusL2 Apparent power + L2 (AC)
	ApparentPowerPlusL2
	// ApparentPowerPlusL3 Apparent power + L3 (AC)
	ApparentPowerPlusL3
	// ReactivePowerMinus Reactive power - (AC)
	ReactivePowerMinus
	// ReactivePowerMinusL1 Reactive power - L1 (AC)
	ReactivePowerMinusL1
	// ReactivePowerMinusL2 Reactive power - L2 (AC)
	ReactivePowerMinusL2
	// ReactivePowerMinusL3 Reactive power - L3 (AC)
	ReactivePowerMinusL3
	// ReactivePowerPlus Reactive power + (AC)
	ReactivePowerPlus
	// ReactivePowerPlusL1 Reactive power + L1 (AC)
	ReactivePowerPlusL1
	// ReactivePowerPlusL2 Reactive power + L2 (AC)
	ReactivePowerPlusL2
	// ReactivePowerPlusL3 Reactive power + L3 (AC)
	ReactivePowerPlusL3
	// PowerS1 Power String 1 (DC)
	PowerS1
	// PowerS2 Power String 2 (DC)
	PowerS2
	// PowerFactor Power Factor (AC)
	PowerFactor
	// PowerFactorL1 Power Factor L1 (AC)
	PowerFactorL1
	// PowerFactorL2 Power Factor L2 (AC)
	PowerFactorL2
	// PowerFactorL3 Power Factor L3 (AC)
	PowerFactorL3

	// ActiveEnergyMinus Active Energy - (AC)
	ActiveEnergyMinus
	// ActiveEnergyMinusL1 Active Energy - L1 (AC)
	ActiveEnergyMinusL1
	// ActiveEnergyMinusL2 Active Energy - L2 (AC)
	ActiveEnergyMinusL2
	// ActiveEnergyMinusL3 Active Energy - L3 (AC)
	ActiveEnergyMinusL3
	// ActiveEnergyPlus Active Energy + (AC)
	ActiveEnergyPlus
	// ActiveEnergyPlusL1 Active Energy + L1 (AC)
	ActiveEnergyPlusL1
	// ActiveEnergyPlusL2 Active Energy + L2 (AC)
	ActiveEnergyPlusL2
	// ActiveEnergyPlusL3 Active Energy + L3 (AC)
	ActiveEnergyPlusL3
	// ActiveEnergyPlusToday Active Energy + today (AC)
	ActiveEnergyPlusToday
	// ApparentEnergyMinus Apparent Energy - (AC)
	ApparentEnergyMinus
	// ApparentEnergyMinusL1 Apparent Energy - L1 (AC)
	ApparentEnergyMinusL1
	// ApparentEnergyMinusL2 Apparent Energy - L2 (AC)
	ApparentEnergyMinusL2
	// ApparentEnergyMinusL3 Apparent Energy - L3 (AC)
	ApparentEnergyMinusL3
	// ApparentEnergyPlus Apparent Energy + (AC)
	ApparentEnergyPlus
	// ApparentEnergyPlusL1 Apparent Energy + L1 (AC)
	ApparentEnergyPlusL1
	// ApparentEnergyPlusL2 Apparent Energy + L2 (AC)
	ApparentEnergyPlusL2
	// ApparentEnergyPlusL3 Apparent Energy + L3 (AC)
	ApparentEnergyPlusL3
	// ReactiveEnergyMinus Reactive Energy - (AC)
	ReactiveEnergyMinus
	// ReactiveEnergyMinusL1 Reactive Energy - L1 (AC)
	ReactiveEnergyMinusL1
	// ReactiveEnergyMinusL2 Reactive Energy - L2 (AC)
	ReactiveEnergyMinusL2
	// ReactiveEnergyMinusL3 Reactive Energy - L3 (AC)
	ReactiveEnergyMinusL3
	// ReactiveEnergyPlus Reactive Energy + (AC)
	ReactiveEnergyPlus
	// ReactiveEnergyPlusL1 Reactive Energy + L1 (AC)
	ReactiveEnergyPlusL1
	// ReactiveEnergyPlusL2 Reactive Energy + L2 (AC)
	ReactiveEnergyPlusL2
	// ReactiveEnergyPlusL3 Reactive Energy + L3 (AC)
	ReactiveEnergyPlusL3

	// CurrentL1 Current L1 (AC)
	CurrentL1
	// CurrentL2 Current L2 (AC)
	CurrentL2
	// CurrentL3 Current L3 (AC)
	CurrentL3
	// CurrentS1 Current String 1 (DC)
	CurrentS1
	// CurrentS2 Current String 2 (DC)
	CurrentS2

	// VoltageL1 Voltage L1 (AC)
	VoltageL1
	// VoltageL2 Voltage L2 (AC)
	VoltageL2
	// VoltageL3 Voltage L3 (AC)
	VoltageL3
	// VoltageS1 Voltage String 1 (DC)
	VoltageS1
	// VoltageS2 Voltage String 2 (DC)
	VoltageS2

	// TimeFeed Feed in time
	TimeFeed
	// TimeOperating Operation time
	TimeOperating
	// UtilityFrequency Utility frequency
	UtilityFrequency

	// BatteryCharge Charge state of battery
	BatteryCharge
	// BatteryTemperature Temperature of battery
	BatteryTemperature

	// DeviceClass ID of device class
	DeviceClass
	// DeviceGridRelay Status of grid relay
	DeviceGridRelay
	// DeviceName Name of device
	DeviceName
	// DeviceStatus Status of device
	DeviceStatus
	// DeviceTemperature Temperature of device
	DeviceTemperature
	// DeviceType ID of device type
	DeviceType
	// SoftwareVersion Software version of device
	SoftwareVersion
)

// ValueDescription describes a value
type ValueDescription struct {
	Description string
	Unit        string
	Type        string
}

// valueDesc provides additional information
var valueDesc = map[ValueID]ValueDescription{
	ActivePowerMax:       {"Maximum active power (AC)", "W", "power"},
	ActivePowerMinus:     {"Active power - (AC)", "W", "power"},
	ActivePowerMinusL1:   {"Active power - L1 (AC)", "W", "power"},
	ActivePowerMinusL2:   {"Active power - L2 (AC)", "W", "power"},
	ActivePowerMinusL3:   {"Active power - L3 (AC)", "W", "power"},
	ActivePowerPlus:      {"Active power + (AC)", "W", "power"},
	ActivePowerPlusL1:    {"Active power + L1 (AC)", "W", "power"},
	ActivePowerPlusL2:    {"Active power + L2 (AC)", "W", "power"},
	ActivePowerPlusL3:    {"Active power + L3 (AC)", "W", "power"},
	ApparentPowerMinus:   {"Apparent power - (AC)", "VA", "power"},
	ApparentPowerMinusL1: {"Apparent power - L1 (AC)", "VA", "power"},
	ApparentPowerMinusL2: {"Apparent power - L2 (AC)", "VA", "power"},
	ApparentPowerMinusL3: {"Apparent power - L3 (AC)", "VA", "power"},
	ApparentPowerPlus:    {"Apparent power + (AC)", "VA", "power"},
	ApparentPowerPlusL1:  {"Apparent power + L1 (AC)", "VA", "power"},
	ApparentPowerPlusL2:  {"Apparent power + L2 (AC)", "VA", "power"},
	ApparentPowerPlusL3:  {"Apparent power + L3 (AC)", "VA", "power"},
	ReactivePowerMinus:   {"Reactive power - (AC)", "var", "power"},
	ReactivePowerMinusL1: {"Reactive power - L1 (AC)", "var", "power"},
	ReactivePowerMinusL2: {"Reactive power - L2 (AC)", "var", "power"},
	ReactivePowerMinusL3: {"Reactive power - L3 (AC)", "var", "power"},
	ReactivePowerPlus:    {"Reactive power + (AC)", "var", "power"},
	ReactivePowerPlusL1:  {"Reactive power + L1 (AC)", "var", "power"},
	ReactivePowerPlusL2:  {"Reactive power + L2 (AC)", "var", "power"},
	ReactivePowerPlusL3:  {"Reactive power + L3 (AC)", "var", "power"},
	PowerS1:              {"Power String 1 (DC)", "W", "power"},
	PowerS2:              {"Power String 2 (DC)", "W", "power"},
	PowerFactor:          {"Power Factor (AC)", "", ""},
	PowerFactorL1:        {"Power Factor L1 (AC)", "", ""},
	PowerFactorL2:        {"Power Factor L2 (AC)", "", ""},
	PowerFactorL3:        {"Power Factor L3 (AC)", "", ""},

	ActiveEnergyMinus:     {"Active Energy - (AC)", "Ws", "energy"},
	ActiveEnergyMinusL1:   {"Active Energy - L1 (AC)", "Ws", "energy"},
	ActiveEnergyMinusL2:   {"Active Energy - L2 (AC)", "Ws", "energy"},
	ActiveEnergyMinusL3:   {"Active Energy - L3 (AC)", "Ws", "energy"},
	ActiveEnergyPlus:      {"Active Energy + (AC)", "Ws", "energy"},
	ActiveEnergyPlusL1:    {"Active Energy + L1 (AC)", "Ws", "energy"},
	ActiveEnergyPlusL2:    {"Active Energy + L2 (AC)", "Ws", "energy"},
	ActiveEnergyPlusL3:    {"Active Energy + L3 (AC)", "Ws", "energy"},
	ActiveEnergyPlusToday: {"Active Energy + today (AC)", "Ws", "energy"},
	ApparentEnergyMinus:   {"Apparent Energy - (AC)", "VAs", "energy"},
	ApparentEnergyMinusL1: {"Apparent Energy - L1 (AC)", "VAs", "energy"},
	ApparentEnergyMinusL2: {"Apparent Energy - L2 (AC)", "VAs", "energy"},
	ApparentEnergyMinusL3: {"Apparent Energy - L3 (AC)", "VAs", "energy"},
	ApparentEnergyPlus:    {"Apparent Energy + (AC)", "VAs", "energy"},
	ApparentEnergyPlusL1:  {"Apparent Energy + L1 (AC)", "VAs", "energy"},
	ApparentEnergyPlusL2:  {"Apparent Energy + L2 (AC)", "VAs", "energy"},
	ApparentEnergyPlusL3:  {"Apparent Energy + L3 (AC)", "VAs", "energy"},
	ReactiveEnergyMinus:   {"Reactive Energy - (AC)", "vars", "energy"},
	ReactiveEnergyMinusL1: {"Reactive Energy - L1 (AC)", "vars", "energy"},
	ReactiveEnergyMinusL2: {"Reactive Energy - L2 (AC)", "vars", "energy"},
	ReactiveEnergyMinusL3: {"Reactive Energy - L3 (AC)", "vars", "energy"},
	ReactiveEnergyPlus:    {"Reactive Energy + (AC)", "vars", "energy"},
	ReactiveEnergyPlusL1:  {"Reactive Energy + L1 (AC)", "vars", "energy"},
	ReactiveEnergyPlusL2:  {"Reactive Energy + L2 (AC)", "vars", "energy"},
	ReactiveEnergyPlusL3:  {"Reactive Energy + L3 (AC)", "vars", "energy"},

	CurrentL1: {"Current L1 (AC)", "A", "current"},
	CurrentL2: {"Current L2 (AC)", "A", "current"},
	CurrentL3: {"Current L3 (AC)", "A", "current"},
	CurrentS1: {"Current String 1 (DC)", "A", "current"},
	CurrentS2: {"Current String 2 (DC)", "A", "current"},

	VoltageL1: {"Voltage L1 (AC)", "V", "voltage"},
	VoltageL2: {"Voltage L2 (AC)", "V", "voltage"},
	VoltageL3: {"Voltage L3 (AC)", "V", "voltage"},
	VoltageS1: {"Voltage String 1 (DC)", "V", "voltage"},
	VoltageS2: {"Voltage String 2 (DC)", "V", "voltage"},

	TimeFeed:         {"Feed in time", "s", ""},
	TimeOperating:    {"Operation time", "s", ""},
	UtilityFrequency: {"Utility frequency", "Hz", ""},

	BatteryCharge:      {"Charge state of battery", "%", ""},
	BatteryTemperature: {"Temperature of battery", "°C", "temperature"},

	DeviceClass:       {"ID of device class", "", ""},
	DeviceGridRelay:   {"Status of grid relay", "", ""},
	DeviceName:        {"Name of device", "", ""},
	DeviceStatus:      {"Status of device", "", ""},
	DeviceTemperature: {"Temperature of device", "°C", "temperature"},
	DeviceType:        {"ID of device type", "", ""},
	SoftwareVersion:   {"Software version of device", "", ""},
}

// GetValueDescription for value
func GetValueDescription(id ValueID) string {
	return valueDesc[id].Description
}

// GetValueInfo for value
func GetValueInfo(id ValueID) ValueDescription {
	return valueDesc[id]
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

// cache for em values
var (
	// emObisMap maps OBIS to energyMeterValuesDef
	emObisMap map[string]energyMeterValuesDef
	// emIDMap maps ValueID to energyMeterValuesDef
	emIDMap map[ValueID]energyMeterValuesDef
)

func init() {
	inverterResponseValues = make(map[uint32]ValueID, len(inverterValues))
	for _, def := range inverterValues {
		inverterResponseValues[uint32(def.Code)<<16+uint32(def.Class)] = def.ID
	}

	inverterAllRequests = getInverterRequests(inverterValues)

	inverterValueMap = make(map[ValueID]InverterValuesDef, len(inverterValues))
	for _, def := range inverterValues {
		inverterValueMap[def.ID] = def
	}

	emObisMap = make(map[string]energyMeterValuesDef, len(emValues))
	for _, def := range emValues {
		emObisMap[def.OBIS] = def
	}
	emIDMap = make(map[ValueID]energyMeterValuesDef, len(emValues))
	for _, def := range emValues {
		emIDMap[def.ID] = def
	}
}

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

// checkInverterValue checks if response is a known value
func checkInverterValue(value *net2.ResponseValue) ValueID {
	if def, ok := inverterResponseValues[uint32(value.Code)<<16+uint32(value.Class)]; ok {
		return def
	}
	if def, ok := inverterResponseValues[uint32(value.Code)<<16]; ok {
		return def
	}
	return 0
}

// getAllInverterRequests to receive all values
func getAllInverterRequests() []InverterValuesDef {
	return inverterAllRequests
}

// getInverterRequest for given ID
func getInverterRequest(id ValueID) InverterValuesDef {
	return inverterValueMap[id]
}

// getInverterRequests to receive all of the given values (reduce request amount)
func getInverterRequests(values []InverterValuesDef) []InverterValuesDef {
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

// parseInverterValues from response
func parseInverterValues(values []*net2.ResponseValue) map[ValueID]interface{} {
	data := make(map[ValueID]interface{}, len(values))

	for _, val := range values {
		if len(val.Values) == 0 {
			continue
		}

		if id := checkInverterValue(val); id != 0 {
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

// energyMeterValuesDef defines a value of an energy meter value
type energyMeterValuesDef struct {
	OBIS   string
	ID     ValueID
	Factor float64
}

// emValues contains all values that can be read from energy meter
var emValues = []energyMeterValuesDef{
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

// convertEnergyMeterValues from OBIS to ID based map
func convertEnergyMeterValues(values map[string]interface{}) map[ValueID]interface{} {
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
			Log.Printf("unknown obis value received: %s", obis)
		}
	}
	return data
}
