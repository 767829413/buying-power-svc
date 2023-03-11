package static

import "gitlab.com/eAuction/enumer/pkg/state"

type DamageLevel int32

const (
	DamageLevelOther      DamageLevel = 1
	DamageLevelLow        DamageLevel = 2
	DamageLevelLowMedium  DamageLevel = 3
	DamageLevelMedium     DamageLevel = 4
	DamageLevelMediumHigh DamageLevel = 5
	DamageLevelHigh       DamageLevel = 6
)

var DamageLevelNames = map[int32]string{
	1: "OTHER",
	2: "LOW",
	3: "LOW_MEDIUM",
	4: "MEDIUM",
	5: "MEDIUM_HIGH",
	6: "HIGH",
}

var DamageLevelValues = map[string]int32{
	"OTHER":       1,
	"LOW":         2,
	"LOW_MEDIUM":  3,
	"MEDIUM":      4,
	"MEDIUM_HIGH": 5,
	"HIGH":        6,
}

// DamageLevelAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var DamageLevelAliases = []state.Alias{
	{Alias: "^other|unknown|null$", Value: 1},
	{Alias: "^low$", Value: 2},
	{Alias: "^lowmedium$", Value: 3},
	{Alias: "^medium$", Value: 4},
	{Alias: "^mediumhigh$", Value: 5},
	{Alias: "^high$", Value: 6},
}

// DamageLevelOrderedValues is a list of sorted values
// of DamageLevel. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var DamageLevelOrderedValues = []int32{
	1,
	2,
	3,
	4,
	5,
	6,
}
