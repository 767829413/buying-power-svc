package static

import "gitlab.com/eAuction/enumer/pkg/state"

type Condition int32

const (
	ConditionGood        Condition = 2
	ConditionMinorDamage Condition = 3
	ConditionDamage      Condition = 4
	ConditionOther       Condition = 1
)

var ConditionNames = map[int32]string{
	2: "GOOD",
	3: "MINOR_DAMAGE",
	4: "DAMAGE",
	1: "OTHER",
}

var ConditionValues = map[string]int32{
	"GOOD":         2,
	"MINOR_DAMAGE": 3,
	"DAMAGE":       4,
	"OTHER":        1,
}

// ConditionAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var ConditionAliases = []state.Alias{
	{Alias: "^good$", Value: 2},
	{Alias: "^minordamage$", Value: 3},
	{Alias: "^damage$", Value: 4},
	{Alias: "^other|unknown$", Value: 1},
}

// ConditionOrderedValues is a list of sorted values
// of Condition. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var ConditionOrderedValues = []int32{
	2,
	3,
	4,
	1,
}
