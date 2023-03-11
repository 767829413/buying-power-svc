package static

import "gitlab.com/eAuction/enumer/pkg/state"

type OdometerUnit int32

const (
	OdometerUnitOther      OdometerUnit = 1
	OdometerUnitKilometers OdometerUnit = 2
	OdometerUnitMiles      OdometerUnit = 3
)

var OdometerUnitNames = map[int32]string{
	1: "OTHER",
	2: "KILOMETERS",
	3: "MILES",
}

var OdometerUnitValues = map[string]int32{
	"OTHER":      1,
	"KILOMETERS": 2,
	"MILES":      3,
}

// OdometerUnitAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var OdometerUnitAliases = []state.Alias{
	{Alias: "^(other|unknown|null)$", Value: 1},
	{Alias: "(km|kms|kilometers)$", Value: 2},
	{Alias: "(mile|miles)$", Value: 3},
}

// OdometerUnitOrderedValues is a list of sorted values
// of OdometerUnit. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var OdometerUnitOrderedValues = []int32{
	1,
	2,
	3,
}
