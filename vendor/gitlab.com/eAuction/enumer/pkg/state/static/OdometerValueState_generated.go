package static

import "gitlab.com/eAuction/enumer/pkg/state"

type OdometerValueState int32

const (
	OdometerValueStateOther                   OdometerValueState = 1
	OdometerValueStateActual                  OdometerValueState = 2
	OdometerValueStateBroken                  OdometerValueState = 3
	OdometerValueStateBurnt                   OdometerValueState = 4
	OdometerValueStateExceedsMechanicalLimits OdometerValueState = 5
	OdometerValueStateExempt                  OdometerValueState = 6
	OdometerValueStateInoperableDigitalDash   OdometerValueState = 7
	OdometerValueStateMissing                 OdometerValueState = 8
	OdometerValueStateNotActual               OdometerValueState = 9
	OdometerValueStateNotRequiredOrExempt     OdometerValueState = 10
)

var OdometerValueStateNames = map[int32]string{
	1:  "OTHER",
	2:  "ACTUAL",
	3:  "BROKEN",
	4:  "BURNT",
	5:  "EXCEEDS_MECHANICAL_LIMITS",
	6:  "EXEMPT",
	7:  "INOPERABLE_DIGITAL_DASH",
	8:  "MISSING",
	9:  "NOT_ACTUAL",
	10: "NOT_REQUIRED_OR_EXEMPT",
}

var OdometerValueStateValues = map[string]int32{
	"OTHER":                     1,
	"ACTUAL":                    2,
	"BROKEN":                    3,
	"BURNT":                     4,
	"EXCEEDS_MECHANICAL_LIMITS": 5,
	"EXEMPT":                    6,
	"INOPERABLE_DIGITAL_DASH":   7,
	"MISSING":                   8,
	"NOT_ACTUAL":                9,
	"NOT_REQUIRED_OR_EXEMPT":    10,
}

// OdometerValueStateAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var OdometerValueStateAliases = []state.Alias{
	{Alias: "^(other|unknown|null)$", Value: 1},
	{Alias: "^actual$", Value: 2},
	{Alias: "^broken$", Value: 3},
	{Alias: "^burnt$", Value: 4},
	{Alias: "^exceedsmechanicallimits$", Value: 5},
	{Alias: "^exempt$", Value: 6},
	{Alias: "^inoperabledigitalhash$", Value: 7},
	{Alias: "^missing$", Value: 8},
	{Alias: "^notactual$", Value: 9},
	{Alias: "^notrequiredorexempt$", Value: 10},
}

// OdometerValueStateOrderedValues is a list of sorted values
// of OdometerValueState. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var OdometerValueStateOrderedValues = []int32{
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
}
