package static

import "gitlab.com/eAuction/enumer/pkg/state"

type BodyStyle int32

const (
	BodyStyleOther       BodyStyle = 1
	BodyStyleCoupe       BodyStyle = 2
	BodyStyleHatchback   BodyStyle = 3
	BodyStyleTruck       BodyStyle = 4
	BodyStyleConvertible BodyStyle = 5
	BodyStyleLimousine   BodyStyle = 6
	BodyStyleSuv         BodyStyle = 7
	BodyStyleSedan       BodyStyle = 8
	BodyStyleVan         BodyStyle = 9
)

var BodyStyleNames = map[int32]string{
	1: "OTHER",
	2: "COUPE",
	3: "HATCHBACK",
	4: "TRUCK",
	5: "CONVERTIBLE",
	6: "LIMOUSINE",
	7: "SUV",
	8: "SEDAN",
	9: "VAN",
}

var BodyStyleValues = map[string]int32{
	"OTHER":       1,
	"COUPE":       2,
	"HATCHBACK":   3,
	"TRUCK":       4,
	"CONVERTIBLE": 5,
	"LIMOUSINE":   6,
	"SUV":         7,
	"SEDAN":       8,
	"VAN":         9,
}

// BodyStyleAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var BodyStyleAliases = []state.Alias{
	{Alias: "^(other|unknown|null)$", Value: 1},
	{Alias: "^(gliders|sportutilityvehicle2door|coupe3d|2door|hatchback2door|coupe|3door|2dcoupe|coupe2door|enduro|roadster|sportsv|3doorextcabpk|converti|2drspor|coupe3door|sport|incompletestripchassis2door|motorhome2door|competition|sedan2door)$", Value: 2},
	{Alias: "^(hatchback4door|hatchback|wagon|wagon4door|hearse4door|hatchbac|wagon2door|station|4dextwagonawd|hatchback5door)$", Value: 3},
	{Alias: "^(dump|busnonschool|doublecab|grain|garbagerefuse|livestock|straighttruck|chassiscab|logging|platformcab|tractortrkdiesel|busschool|tiltcab|straighttruck4dr|straighttruck2doo|truck|beverage|cementmixer|tractortruck|firetruck|tank)$", Value: 4},
	{Alias: "^(convertible|convertible2door)$", Value: 5},
	{Alias: "^(limousine4door|limousine|extended|saloon)$", Value: 6},
	{Alias: "^(sportutilitytruck4door|chassis|4dsr54x4|cutaway2door|cabchassis|sportuttruck4dr|cabchassis2door|pickup2door|4dutility2wd4cyl|clubcabpickup|utility|incompletepickup2|incompletepickup2door|incpickup2door|incompletepickup4|suv5door|crossover|crewcab2wdswb|flatbed|4dle4x2|cabchassis4door|suv|4dsuv4wd|incpickup4door|incompletepickup4door|sportutilityvehicle|cruiser|pickup4door|crewpic|incstripchassis|straighttruck2dr|clubcab|commercialchassis4door|4dsedanls|suv4door|sportutilityvehicl|4drext|crewpickup|pickup|suv2door|4dsuv|pickup3door|sportutilitytruck)$", Value: 7},
	{Alias: "^(traditional|sidebyside|incompletestripch|touring|sportutilityvehicle4door|incstripchass2dr|4drspor|4dutilityfwd|coupe4door|sedan4d|conventionalcab|scooter|atv|cabandchassis|containchassis|cutaway3door|4dsedan|incompletestripchassis|sedan|4door|sedan4door|cutaway)$", Value: 8},
	{Alias: "^(vancargo5door|vanpassenger3door|motorhome|vanpassenger|stepvan2door|vanpassenger5door|cargovan|stationwagon|vanpassenger4door|wagon3door|stepvan|van|vancamper3door|passengervan|vancargo3door|cargova|panelvan|vancargo4door|mpv|sportvan|vancargo2door|vancargo|motorhome3door)$", Value: 9},
}

// BodyStyleOrderedValues is a list of sorted values
// of BodyStyle. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var BodyStyleOrderedValues = []int32{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}
