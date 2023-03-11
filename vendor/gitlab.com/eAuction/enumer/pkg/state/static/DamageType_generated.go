package static

import "gitlab.com/eAuction/enumer/pkg/state"

type DamageType int32

const (
	DamageTypeOther                DamageType = 1
	DamageTypeRejectedRepair       DamageType = 2
	DamageTypeElectrical           DamageType = 3
	DamageTypeSide                 DamageType = 4
	DamageTypeTheft                DamageType = 5
	DamageTypeStrip                DamageType = 6
	DamageTypeRepossession         DamageType = 7
	DamageTypeFrame                DamageType = 8
	DamageTypeNormalWear           DamageType = 9
	DamageTypePartialRepair        DamageType = 10
	DamageTypeFrontEnd             DamageType = 11
	DamageTypeAllOver              DamageType = 12
	DamageTypeMinorDentOrScratches DamageType = 13
	DamageTypeLeftRear             DamageType = 14
	DamageTypeTransmission         DamageType = 15
	DamageTypeVandalism            DamageType = 16
	DamageTypeRightFront           DamageType = 17
	DamageTypeSuspension           DamageType = 18
	DamageTypeRearEnd              DamageType = 19
	DamageTypeRightSide            DamageType = 20
	DamageTypeFlood                DamageType = 21
	DamageTypeFreshWater           DamageType = 22
	DamageTypeRollover             DamageType = 23
	DamageTypeUndercarriage        DamageType = 24
	DamageTypeRightRear            DamageType = 25
	DamageTypeLeftFront            DamageType = 26
	DamageTypeInteriorBurn         DamageType = 27
	DamageTypeDamageHistory        DamageType = 28
	DamageTypeRoof                 DamageType = 29
	DamageTypeFrontAndRear         DamageType = 30
	DamageTypeHail                 DamageType = 31
	DamageTypeMechanical           DamageType = 32
	DamageTypeSaltWater            DamageType = 33
	DamageTypeTotalBurn            DamageType = 34
	DamageTypeBurn                 DamageType = 35
	DamageTypeReplacedVin          DamageType = 36
	DamageTypeBiohazardOrChemical  DamageType = 37
	DamageTypeLeftAndRight         DamageType = 38
	DamageTypeEngineBurn           DamageType = 39
	DamageTypeLeftSide             DamageType = 40
	DamageTypeNeedsRepair          DamageType = 41
	DamageTypeStormDamage          DamageType = 42
	DamageTypeNone                 DamageType = 43
	DamageTypeExteriorBurn         DamageType = 44
	DamageTypeEngineDamage         DamageType = 45
	DamageTypeMissingOrAlteredVin  DamageType = 46
	DamageTypeCashForClunkers      DamageType = 47
)

var DamageTypeNames = map[int32]string{
	1:  "OTHER",
	2:  "REJECTED_REPAIR",
	3:  "ELECTRICAL",
	4:  "SIDE",
	5:  "THEFT",
	6:  "STRIP",
	7:  "REPOSSESSION",
	8:  "FRAME",
	9:  "NORMAL_WEAR",
	10: "PARTIAL_REPAIR",
	11: "FRONT_END",
	12: "ALL_OVER",
	13: "MINOR_DENT_OR_SCRATCHES",
	14: "LEFT_REAR",
	15: "TRANSMISSION",
	16: "VANDALISM",
	17: "RIGHT_FRONT",
	18: "SUSPENSION",
	19: "REAR_END",
	20: "RIGHT_SIDE",
	21: "FLOOD",
	22: "FRESH_WATER",
	23: "ROLLOVER",
	24: "UNDERCARRIAGE",
	25: "RIGHT_REAR",
	26: "LEFT_FRONT",
	27: "INTERIOR_BURN",
	28: "DAMAGE_HISTORY",
	29: "ROOF",
	30: "FRONT_AND_REAR",
	31: "HAIL",
	32: "MECHANICAL",
	33: "SALT_WATER",
	34: "TOTAL_BURN",
	35: "BURN",
	36: "REPLACED_VIN",
	37: "BIOHAZARD_OR_CHEMICAL",
	38: "LEFT_AND_RIGHT",
	39: "ENGINE_BURN",
	40: "LEFT_SIDE",
	41: "NEEDS_REPAIR",
	42: "STORM_DAMAGE",
	43: "NONE",
	44: "EXTERIOR_BURN",
	45: "ENGINE_DAMAGE",
	46: "MISSING_OR_ALTERED_VIN",
	47: "CASH_FOR_CLUNKERS",
}

var DamageTypeValues = map[string]int32{
	"OTHER":                   1,
	"REJECTED_REPAIR":         2,
	"ELECTRICAL":              3,
	"SIDE":                    4,
	"THEFT":                   5,
	"STRIP":                   6,
	"REPOSSESSION":            7,
	"FRAME":                   8,
	"NORMAL_WEAR":             9,
	"PARTIAL_REPAIR":          10,
	"FRONT_END":               11,
	"ALL_OVER":                12,
	"MINOR_DENT_OR_SCRATCHES": 13,
	"LEFT_REAR":               14,
	"TRANSMISSION":            15,
	"VANDALISM":               16,
	"RIGHT_FRONT":             17,
	"SUSPENSION":              18,
	"REAR_END":                19,
	"RIGHT_SIDE":              20,
	"FLOOD":                   21,
	"FRESH_WATER":             22,
	"ROLLOVER":                23,
	"UNDERCARRIAGE":           24,
	"RIGHT_REAR":              25,
	"LEFT_FRONT":              26,
	"INTERIOR_BURN":           27,
	"DAMAGE_HISTORY":          28,
	"ROOF":                    29,
	"FRONT_AND_REAR":          30,
	"HAIL":                    31,
	"MECHANICAL":              32,
	"SALT_WATER":              33,
	"TOTAL_BURN":              34,
	"BURN":                    35,
	"REPLACED_VIN":            36,
	"BIOHAZARD_OR_CHEMICAL":   37,
	"LEFT_AND_RIGHT":          38,
	"ENGINE_BURN":             39,
	"LEFT_SIDE":               40,
	"NEEDS_REPAIR":            41,
	"STORM_DAMAGE":            42,
	"NONE":                    43,
	"EXTERIOR_BURN":           44,
	"ENGINE_DAMAGE":           45,
	"MISSING_OR_ALTERED_VIN":  46,
	"CASH_FOR_CLUNKERS":       47,
}

// DamageTypeAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var DamageTypeAliases = []state.Alias{
	{Alias: "^other|unknown|null$", Value: 1},
	{Alias: "^(rejectedrepair)$", Value: 2},
	{Alias: "^(electrical|electricaldamage)$", Value: 3},
	{Alias: "^(side|sidedamage)$", Value: 4},
	{Alias: "^(theft)$", Value: 5},
	{Alias: "^(strip|stripped)$", Value: 6},
	{Alias: "^(repossession)$", Value: 7},
	{Alias: "^(frame|framedamage)$", Value: 8},
	{Alias: "^(normalwear)$", Value: 9},
	{Alias: "^(partialrepair)$", Value: 10},
	{Alias: "^(frontend|frontenddamage)$", Value: 11},
	{Alias: "^(allover|alloverdamage)$", Value: 12},
	{Alias: "^(minordentorscratches$|minordentscratches)$", Value: 13},
	{Alias: "^(leftrear|leftreardamage)$", Value: 14},
	{Alias: "^(transmission)$", Value: 15},
	{Alias: "^(vandalized|vandalised|vandalism)$", Value: 16},
	{Alias: "^(rightfront|rightfrontdamage)$", Value: 17},
	{Alias: "^(suspension|suspensiondamage)$", Value: 18},
	{Alias: "^(rear|rearend|rearenddamage)$", Value: 19},
	{Alias: "^(rightside|rightsidedamage)$", Value: 20},
	{Alias: "^(flood|waterflood)$", Value: 21},
	{Alias: "^(freshwater)$", Value: 22},
	{Alias: "^(rollover)$", Value: 23},
	{Alias: "^(undercarriage|undercarriagedamage)$", Value: 24},
	{Alias: "^(rightrear|rightreardamage)$", Value: 25},
	{Alias: "^(leftfront|leftfrontdamage)$", Value: 26},
	{Alias: "^(interiorburn|burninterior)$", Value: 27},
	{Alias: "^(damagehistory)$", Value: 28},
	{Alias: "^(toproof|roof|roofdamage)$", Value: 29},
	{Alias: "^(frontrear|frontandrear|frontreardamage)$", Value: 30},
	{Alias: "^(hail|haildamage)$", Value: 31},
	{Alias: "^(mechanical|mechanicaldamage)$", Value: 32},
	{Alias: "^(saltwater)$", Value: 33},
	{Alias: "^(totalburn)$", Value: 34},
	{Alias: "^(burn)$", Value: 35},
	{Alias: "^(replacedvin)$", Value: 36},
	{Alias: "^(biohazardchemical|biohazardorchemical|biohazardorchemicaldamage|biohazard)$", Value: 37},
	{Alias: "^(leftright|leftrightdamage|leftandright)$", Value: 38},
	{Alias: "^(engineburn|burnengine)$", Value: 39},
	{Alias: "^(leftside|leftsidedamage)$", Value: 40},
	{Alias: "^(needsrepair)$", Value: 41},
	{Alias: "^(stormdamage)$", Value: 42},
	{Alias: "^(none|nodamage)$", Value: 43},
	{Alias: "^(burnexterior|exteriorburn)$", Value: 44},
	{Alias: "^(enginedamage)$", Value: 45},
	{Alias: "^(missingalteredvin|missingoralteredvin)$", Value: 46},
	{Alias: "^(cashforclunkers|cash4clunker|cash4clunkers)$", Value: 47},
}

// DamageTypeOrderedValues is a list of sorted values
// of DamageType. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var DamageTypeOrderedValues = []int32{
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
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
	35,
	36,
	37,
	38,
	39,
	40,
	41,
	42,
	43,
	44,
	45,
	46,
	47,
}
