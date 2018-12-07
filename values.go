package sunny

import "gitlab.com/bboehmke/sunny/proto"

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
	{0x5380, 0x00451F00, 0x004521FF, 0x01, 0x4521, "current_dc1"},
	{0x5380, 0x00451F00, 0x004521FF, 0x02, 0x451F, "voltage_dc2"},
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
var _responseValues map[uint32]string
var _allRequests []valDef

// defInit initialize cache
func defInit() {
	if _responseValues != nil {
		return
	}

	_responseValues = make(map[uint32]string, len(valuesDef))
	for _, def := range valuesDef {
		_responseValues[uint32(def.Code)<<16+uint32(def.Class)] = def.Key
	}

	_allRequests = getRequests(valuesDef)
}

// _checkValue checks if response is a known value (without cache initialization)
func _checkValue(value proto.ResponseValue) string {
	if def, ok := _responseValues[uint32(value.Code)<<16+uint32(value.Class)]; ok {
		return def
	}
	if def, ok := _responseValues[uint32(value.Code)<<16]; ok {
		return def
	}
	return ""
}

// getAllRequests to receive all values
func getAllRequests() []valDef {
	defInit()

	return _allRequests
}

// getRequest for given key
func getRequest(key string) valDef {
	for _, def := range valuesDef {
		if def.Key == key {
			return def
		}
	}
	return valDef{}
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
func parseValues(values []proto.ResponseValue) map[string]interface{} {
	defInit()

	data := make(map[string]interface{}, len(values))

	for _, val := range values {
		if len(val.Values) == 0 {
			continue
		}

		if key := _checkValue(val); key != "" {
			data[key] = val.Values[0]
		}
	}

	return data
}
