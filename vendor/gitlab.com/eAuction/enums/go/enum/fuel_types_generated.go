// Generated. DO NOT EDIT.
package enum

import (
	"gitlab.com/eAuction/enums/go/internal/alias"
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

var fuelTypes = []EnumValue{
	{Name: "CNG", ReadableName: "CNG"},
	{Name: "DIESEL", ReadableName: "DIESEL"},
	{Name: "ELECTRIC", ReadableName: "ELECTRIC"},
	{Name: "ETHANOL", ReadableName: "ETHANOL"},
	{Name: "FLEXIBLE FUEL", ReadableName: "FLEXIBLE FUEL"},
	{Name: "GAS", ReadableName: "GASOLINE"},
	{Name: "HYBRID", ReadableName: "HYBRID"},
	{Name: "HYDROGEN", ReadableName: "HYDROGEN"},
	{Name: "LPG", ReadableName: "LPG"},
	{Name: "OTHER", ReadableName: "OTHER"},
	{Name: "PLUGIN HYBRID", ReadableName: "PLUGIN HYBRID"},
}

// FuelType parses and returns value of Fuel_Type enum.
func FuelType(s string) Fuel_Type {
	norm, ok := alias.FuelTypes[alphanum.Lower(s)]
	if !ok {
		return Fuel_TYPE_UNKNOWN
	}
	return mapAliasToFuelTypes[norm]
}

// Name returns string name of the enum value.
func (x Fuel_Type) Name() string {
	v, ok := mapFuelTypesToEnumValue[x]
	if !ok {
		return "UNKNOWN"
	}
	return v.Name
}

var mapAliasToFuelTypes = map[string]Fuel_Type{
	"CNG":           Fuel_TYPE_CNG,
	"DIESEL":        Fuel_TYPE_DIESEL,
	"ELECTRIC":      Fuel_TYPE_ELECTRIC,
	"ETHANOL":       Fuel_TYPE_ETHANOL,
	"FLEXIBLE FUEL": Fuel_TYPE_FLEXIBLE,
	"GAS":           Fuel_TYPE_GAS,
	"HYBRID":        Fuel_TYPE_HYBRID,
	"HYDROGEN":      Fuel_TYPE_HYDROGEN,
	"LPG":           Fuel_TYPE_LPG,
	"OTHER":         Fuel_TYPE_NULL,
	"PLUGIN HYBRID": Fuel_TYPE_PLUGIN_HYBRID,
}

var mapFuelTypesToEnumValue = map[Fuel_Type]EnumValue{
	Fuel_TYPE_CNG:           {Name: "CNG", ReadableName: "CNG"},
	Fuel_TYPE_DIESEL:        {Name: "DIESEL", ReadableName: "DIESEL"},
	Fuel_TYPE_ELECTRIC:      {Name: "ELECTRIC", ReadableName: "ELECTRIC"},
	Fuel_TYPE_ETHANOL:       {Name: "ETHANOL", ReadableName: "ETHANOL"},
	Fuel_TYPE_FLEXIBLE:      {Name: "FLEXIBLE FUEL", ReadableName: "FLEXIBLE FUEL"},
	Fuel_TYPE_GAS:           {Name: "GAS", ReadableName: "GASOLINE"},
	Fuel_TYPE_HYBRID:        {Name: "HYBRID", ReadableName: "HYBRID"},
	Fuel_TYPE_HYDROGEN:      {Name: "HYDROGEN", ReadableName: "HYDROGEN"},
	Fuel_TYPE_LPG:           {Name: "LPG", ReadableName: "LPG"},
	Fuel_TYPE_NULL:          {Name: "OTHER", ReadableName: "OTHER"},
	Fuel_TYPE_PLUGIN_HYBRID: {Name: "PLUGIN HYBRID", ReadableName: "PLUGIN HYBRID"},
}
