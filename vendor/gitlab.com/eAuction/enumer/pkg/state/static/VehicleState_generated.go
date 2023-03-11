package static

import "gitlab.com/eAuction/enumer/pkg/state"

type VehicleState int32

const (
	VehicleStateOther              VehicleState = 1
	VehicleStateRunsAndDrives      VehicleState = 2
	VehicleStateEngineStartProgram VehicleState = 3
	VehicleStateEnhanced           VehicleState = 4
	VehicleStateNew                VehicleState = 5
)

var VehicleStateNames = map[int32]string{
	1: "OTHER",
	2: "RUNS_AND_DRIVES",
	3: "ENGINE_START_PROGRAM",
	4: "ENHANCED",
	5: "NEW",
}

var VehicleStateValues = map[string]int32{
	"OTHER":                1,
	"RUNS_AND_DRIVES":      2,
	"ENGINE_START_PROGRAM": 3,
	"ENHANCED":             4,
	"NEW":                  5,
}

// VehicleStateAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var VehicleStateAliases = []state.Alias{
	{Alias: "^other|unknown|null$", Value: 1},
	{Alias: "^runsanddrives$", Value: 2},
	{Alias: "^(enginestartprogram|runsnotdrives)$", Value: 3},
	{Alias: ".*enhanced.*|^forparts", Value: 4},
	{Alias: "^new$", Value: 5},
}

// VehicleStateOrderedValues is a list of sorted values
// of VehicleState. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var VehicleStateOrderedValues = []int32{
	1,
	2,
	3,
	4,
	5,
}
