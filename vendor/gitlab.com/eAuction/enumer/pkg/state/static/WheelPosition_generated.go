package static

import "gitlab.com/eAuction/enumer/pkg/state"

type WheelPosition int32

const (
	WheelPositionOther WheelPosition = 1
	WheelPositionLeft  WheelPosition = 2
	WheelPositionRight WheelPosition = 3
)

var WheelPositionNames = map[int32]string{
	1: "OTHER",
	2: "LEFT",
	3: "RIGHT",
}

var WheelPositionValues = map[string]int32{
	"OTHER": 1,
	"LEFT":  2,
	"RIGHT": 3,
}

// WheelPositionAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var WheelPositionAliases = []state.Alias{
	{Alias: "^other|unknown|null$", Value: 1},
	{Alias: "^left$", Value: 2},
	{Alias: "^right$", Value: 3},
}

// WheelPositionOrderedValues is a list of sorted values
// of WheelPosition. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var WheelPositionOrderedValues = []int32{
	1,
	2,
	3,
}
