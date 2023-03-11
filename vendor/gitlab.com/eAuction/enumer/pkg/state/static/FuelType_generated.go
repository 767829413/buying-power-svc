package static

import "gitlab.com/eAuction/enumer/pkg/state"

type FuelType int32

const (
	FuelTypeOther        FuelType = 1
	FuelTypeCng          FuelType = 2
	FuelTypeDiesel       FuelType = 3
	FuelTypeElectric     FuelType = 4
	FuelTypeFlexibleFuel FuelType = 5
	FuelTypeGas          FuelType = 6
	FuelTypeHybrid       FuelType = 7
	FuelTypeHydrogen     FuelType = 8
	FuelTypeLpg          FuelType = 9
	FuelTypePluginHybrid FuelType = 10
	FuelTypeEthanol      FuelType = 11
)

var FuelTypeNames = map[int32]string{
	1:  "OTHER",
	2:  "CNG",
	3:  "DIESEL",
	4:  "ELECTRIC",
	5:  "FLEXIBLE_FUEL",
	6:  "GAS",
	7:  "HYBRID",
	8:  "HYDROGEN",
	9:  "LPG",
	10: "PLUGIN_HYBRID",
	11: "ETHANOL",
}

var FuelTypeValues = map[string]int32{
	"OTHER":         1,
	"CNG":           2,
	"DIESEL":        3,
	"ELECTRIC":      4,
	"FLEXIBLE_FUEL": 5,
	"GAS":           6,
	"HYBRID":        7,
	"HYDROGEN":      8,
	"LPG":           9,
	"PLUGIN_HYBRID": 10,
	"ETHANOL":       11,
}

// FuelTypeAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var FuelTypeAliases = []state.Alias{
	{Alias: "^(other|unknown|null|convertibletogaseouspowered|lng|towngas)$", Value: 1},
	{Alias: "^(cng|compressednaturalgas|compressednaturalgascng)$", Value: 2},
	{Alias: "^(diesel)$", Value: 3},
	{Alias: "^(electric)$", Value: 4},
	{Alias: "^(flexiblefuel|flexfuel)$", Value: 5},
	{Alias: "^(gas|gasoline|petrol|petrolgas)$", Value: 6},
	{Alias: "^(hybrid|hybridengine|dieselelect|dieselelectrechargeable|gasolineelect|gasolineelectrechargeable)$", Value: 7},
	{Alias: "^(hydrogen|hydrohen|hydrogenfuelcell)$", Value: 8},
	{Alias: "^(lpg|liquidpetroleumgas)$", Value: 9},
	{Alias: "^(pluginhybrid)$", Value: 10},
	{Alias: "^(ethanol)$", Value: 11},
}

// FuelTypeOrderedValues is a list of sorted values
// of FuelType. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var FuelTypeOrderedValues = []int32{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
}
