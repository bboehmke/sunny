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

// ValueID identifies a value read from a device
type ValueID string

const (
	// ActivePowerMax Maximum active power (AC)
	ActivePowerMax = ValueID("active_power_max")
	// ActivePowerMinus Active power - (AC)
	ActivePowerMinus = ValueID("active_power_minus")
	// ActivePowerMinusL1 Active power - L1 (AC)
	ActivePowerMinusL1 = ValueID("active_power_minus_l1")
	// ActivePowerMinusL2 Active power - L2 (AC)
	ActivePowerMinusL2 = ValueID("active_power_minus_l2")
	// ActivePowerMinusL3 Active power - L3 (AC)
	ActivePowerMinusL3 = ValueID("active_power_minus_l3")
	// ActivePowerPlus Active power + (AC)
	ActivePowerPlus = ValueID("active_power_plus")
	// ActivePowerPlusL1 Active power + L1 (AC)
	ActivePowerPlusL1 = ValueID("active_power_plus_l1")
	// ActivePowerPlusL2 Active power + L2 (AC)
	ActivePowerPlusL2 = ValueID("active_power_plus_l2")
	// ActivePowerPlusL3 Active power + L3 (AC)
	ActivePowerPlusL3 = ValueID("active_power_plus_l3")
	// ApparentPowerMinus Apparent power - (AC)
	ApparentPowerMinus = ValueID("apparent_power_minus")
	// ApparentPowerMinusL1 Apparent power - L1 (AC)
	ApparentPowerMinusL1 = ValueID("apparent_power_minus_l1")
	// ApparentPowerMinusL2 Apparent power - L2 (AC)
	ApparentPowerMinusL2 = ValueID("apparent_power_minus_l2")
	// ApparentPowerMinusL3 Apparent power - L3 (AC)
	ApparentPowerMinusL3 = ValueID("apparent_power_minus_l3")
	// ApparentPowerPlus Apparent power + (AC)
	ApparentPowerPlus = ValueID("apparent_power_plus")
	// ApparentPowerPlusL1 Apparent power + L1 (AC)
	ApparentPowerPlusL1 = ValueID("apparent_power_plus_l1")
	// ApparentPowerPlusL2 Apparent power + L2 (AC)
	ApparentPowerPlusL2 = ValueID("apparent_power_plus_l2")
	// ApparentPowerPlusL3 Apparent power + L3 (AC)
	ApparentPowerPlusL3 = ValueID("apparent_power_plus_l3")
	// ReactivePowerMinus Reactive power - (AC)
	ReactivePowerMinus = ValueID("reactive_power_minus")
	// ReactivePowerMinusL1 Reactive power - L1 (AC)
	ReactivePowerMinusL1 = ValueID("reactive_power_minus_l1")
	// ReactivePowerMinusL2 Reactive power - L2 (AC)
	ReactivePowerMinusL2 = ValueID("reactive_power_minus_l2")
	// ReactivePowerMinusL3 Reactive power - L3 (AC)
	ReactivePowerMinusL3 = ValueID("reactive_power_minus_l3")
	// ReactivePowerPlus Reactive power + (AC)
	ReactivePowerPlus = ValueID("reactive_power_plus")
	// ReactivePowerPlusL1 Reactive power + L1 (AC)
	ReactivePowerPlusL1 = ValueID("reactive_power_plus_l1")
	// ReactivePowerPlusL2 Reactive power + L2 (AC)
	ReactivePowerPlusL2 = ValueID("reactive_power_plus_l2")
	// ReactivePowerPlusL3 Reactive power + L3 (AC)
	ReactivePowerPlusL3 = ValueID("reactive_power_plus_l3")
	// PowerS1 Power String 1 (DC)
	PowerS1 = ValueID("power_s1")
	// PowerS2 Power String 2 (DC)
	PowerS2 = ValueID("power_s2")
	// PowerFactor Power Factor (AC)
	PowerFactor = ValueID("power_factor")
	// PowerFactorL1 Power Factor L1 (AC)
	PowerFactorL1 = ValueID("power_factor_l1")
	// PowerFactorL2 Power Factor L2 (AC)
	PowerFactorL2 = ValueID("power_factor_l2")
	// PowerFactorL3 Power Factor L3 (AC)
	PowerFactorL3 = ValueID("power_factor_l3")

	// ActiveEnergyMinus Active Energy - (AC)
	ActiveEnergyMinus = ValueID("active_energy_minus")
	// ActiveEnergyMinusL1 Active Energy - L1 (AC)
	ActiveEnergyMinusL1 = ValueID("active_energy_minus_l1")
	// ActiveEnergyMinusL2 Active Energy - L2 (AC)
	ActiveEnergyMinusL2 = ValueID("active_energy_minus_l2")
	// ActiveEnergyMinusL3 Active Energy - L3 (AC)
	ActiveEnergyMinusL3 = ValueID("active_energy_minus_l3")
	// ActiveEnergyPlus Active Energy + (AC)
	ActiveEnergyPlus = ValueID("active_energy_plus")
	// ActiveEnergyPlusL1 Active Energy + L1 (AC)
	ActiveEnergyPlusL1 = ValueID("active_energy_plus_l1")
	// ActiveEnergyPlusL2 Active Energy + L2 (AC)
	ActiveEnergyPlusL2 = ValueID("active_energy_plus_l2")
	// ActiveEnergyPlusL3 Active Energy + L3 (AC)
	ActiveEnergyPlusL3 = ValueID("active_energy_plus_l3")
	// ActiveEnergyPlusToday Active Energy + today (AC)
	ActiveEnergyPlusToday = ValueID("active_energy_plus_today")
	// ApparentEnergyMinus Apparent Energy - (AC)
	ApparentEnergyMinus = ValueID("apparent_energy_minus")
	// ApparentEnergyMinusL1 Apparent Energy - L1 (AC)
	ApparentEnergyMinusL1 = ValueID("apparent_energy_minus_l1")
	// ApparentEnergyMinusL2 Apparent Energy - L2 (AC)
	ApparentEnergyMinusL2 = ValueID("apparent_energy_minus_l2")
	// ApparentEnergyMinusL3 Apparent Energy - L3 (AC)
	ApparentEnergyMinusL3 = ValueID("apparent_energy_minus_l3")
	// ApparentEnergyPlus Apparent Energy + (AC)
	ApparentEnergyPlus = ValueID("apparent_energy_plus")
	// ApparentEnergyPlusL1 Apparent Energy + L1 (AC)
	ApparentEnergyPlusL1 = ValueID("apparent_energy_plus_l1")
	// ApparentEnergyPlusL2 Apparent Energy + L2 (AC)
	ApparentEnergyPlusL2 = ValueID("apparent_energy_plus_l2")
	// ApparentEnergyPlusL3 Apparent Energy + L3 (AC)
	ApparentEnergyPlusL3 = ValueID("apparent_energy_plus_l3")
	// ReactiveEnergyMinus Reactive Energy - (AC)
	ReactiveEnergyMinus = ValueID("reactive_energy_minus")
	// ReactiveEnergyMinusL1 Reactive Energy - L1 (AC)
	ReactiveEnergyMinusL1 = ValueID("reactive_energy_minus_l1")
	// ReactiveEnergyMinusL2 Reactive Energy - L2 (AC)
	ReactiveEnergyMinusL2 = ValueID("reactive_energy_minus_l2")
	// ReactiveEnergyMinusL3 Reactive Energy - L3 (AC)
	ReactiveEnergyMinusL3 = ValueID("reactive_energy_minus_l3")
	// ReactiveEnergyPlus Reactive Energy + (AC)
	ReactiveEnergyPlus = ValueID("reactive_energy_plus")
	// ReactiveEnergyPlusL1 Reactive Energy + L1 (AC)
	ReactiveEnergyPlusL1 = ValueID("reactive_energy_plus_l1")
	// ReactiveEnergyPlusL2 Reactive Energy + L2 (AC)
	ReactiveEnergyPlusL2 = ValueID("reactive_energy_plus_l2")
	// ReactiveEnergyPlusL3 Reactive Energy + L3 (AC)
	ReactiveEnergyPlusL3 = ValueID("reactive_energy_plus_l3")

	// CurrentL1 Current L1 (AC)
	CurrentL1 = ValueID("current_l1")
	// CurrentL2 Current L2 (AC)
	CurrentL2 = ValueID("current_l2")
	// CurrentL3 Current L3 (AC)
	CurrentL3 = ValueID("current_l3")
	// CurrentS1 Current String 1 (DC)
	CurrentS1 = ValueID("current_s1")
	// CurrentS2 Current String 2 (DC)
	CurrentS2 = ValueID("current_s2")

	// VoltageL1 Voltage L1 (AC)
	VoltageL1 = ValueID("voltage_l1")
	// VoltageL2 Voltage L2 (AC)
	VoltageL2 = ValueID("voltage_l2")
	// VoltageL3 Voltage L3 (AC)
	VoltageL3 = ValueID("voltage_l3")
	// VoltageS1 Voltage String 1 (DC)
	VoltageS1 = ValueID("voltage_s1")
	// VoltageS2 Voltage String 2 (DC)
	VoltageS2 = ValueID("voltage_s2")

	// TimeFeed Feed in time
	TimeFeed = ValueID("time_feed")
	// TimeOperating Operation time
	TimeOperating = ValueID("time_operating")
	// UtilityFrequency Utility frequency
	UtilityFrequency = ValueID("utility_frequency")

	// BatteryCharge Charge state of battery
	BatteryCharge = ValueID("battery_charge")
	// BatteryTemperature Temperature of battery
	BatteryTemperature = ValueID("battery_temperature")

	// DeviceClass ID of device class
	DeviceClass = ValueID("device_class")
	// DeviceGridRelay Status of grid relay
	DeviceGridRelay = ValueID("device_grid_relay")
	// DeviceName Name of device
	DeviceName = ValueID("device_name")
	// DeviceStatus Status of device
	DeviceStatus = ValueID("device_status")
	// DeviceTemperature Temperature of device
	DeviceTemperature = ValueID("device_temperature")
	// DeviceType ID of device type
	DeviceType = ValueID("device_type")
	// SoftwareVersion Software version of device
	SoftwareVersion = ValueID("software_version")
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
