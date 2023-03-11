package static

import "gitlab.com/eAuction/enumer/pkg/state"

type TransmissionType int32

const (
	TransmissionTypeOther     TransmissionType = 1
	TransmissionTypeAutomatic TransmissionType = 2
	TransmissionTypeManual    TransmissionType = 3
)

var TransmissionTypeNames = map[int32]string{
	1: "OTHER",
	2: "AUTOMATIC",
	3: "MANUAL",
}

var TransmissionTypeValues = map[string]int32{
	"OTHER":     1,
	"AUTOMATIC": 2,
	"MANUAL":    3,
}

// TransmissionTypeAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var TransmissionTypeAliases = []state.Alias{
	{Alias: "^other|unknown|null|missing$", Value: 1},
	{Alias: "^automatic$", Value: 2},
	{Alias: "^(manual|sequential|manualgearbox)$", Value: 3},
}

// TransmissionTypeOrderedValues is a list of sorted values
// of TransmissionType. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var TransmissionTypeOrderedValues = []int32{
	1,
	2,
	3,
}
