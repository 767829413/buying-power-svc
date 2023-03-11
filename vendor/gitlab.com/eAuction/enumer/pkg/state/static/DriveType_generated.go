package static

import "gitlab.com/eAuction/enumer/pkg/state"

type DriveType int32

const (
	DriveTypeOther          DriveType = 1
	DriveTypeFrontDrive     DriveType = 2
	DriveTypeRearDrive      DriveType = 3
	DriveTypeFourXFourDrive DriveType = 4
	DriveTypeAllDrive       DriveType = 5
)

var DriveTypeNames = map[int32]string{
	1: "OTHER",
	2: "FRONT_DRIVE",
	3: "REAR_DRIVE",
	4: "FOUR_X_FOUR_DRIVE",
	5: "ALL_DRIVE",
}

var DriveTypeValues = map[string]int32{
	"OTHER":             1,
	"FRONT_DRIVE":       2,
	"REAR_DRIVE":        3,
	"FOUR_X_FOUR_DRIVE": 4,
	"ALL_DRIVE":         5,
}

// DriveTypeAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var DriveTypeAliases = []state.Alias{
	{Alias: "^other|unknown|null$", Value: 1},
	{Alias: "^(frontdrive|frontwheeldrive)$", Value: 2},
	{Alias: "^(reardrive|rearwheeldrive)$", Value: 3},
	{Alias: "^(4x4wfrontwhldrv|4x4wrearwheeldrv|fourbyfour|rearwheeldrvw4x4|4x4drive|fourxfourdrive)$", Value: 4},
	{Alias: "^(alldrive|allwheeldrive)$", Value: 5},
}

// DriveTypeOrderedValues is a list of sorted values
// of DriveType. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var DriveTypeOrderedValues = []int32{
	1,
	2,
	3,
	4,
	5,
}
