package static

import "gitlab.com/eAuction/enumer/pkg/state"

type Maker int32

const (
	MakerAcura            Maker = 1
	MakerAgusta           Maker = 2
	MakerAlfaRomeo        Maker = 3
	MakerAstonMartin      Maker = 4
	MakerAudi             Maker = 5
	MakerBentley          Maker = 6
	MakerBmw              Maker = 7
	MakerBuick            Maker = 8
	MakerByd              Maker = 9
	MakerCadillac         Maker = 10
	MakerChangan          Maker = 11
	MakerChangfeng        Maker = 12
	MakerChery            Maker = 13
	MakerChevrolet        Maker = 14
	MakerChrysler         Maker = 15
	MakerCitroen          Maker = 16
	MakerCorvette         Maker = 17
	MakerDacia            Maker = 18
	MakerDaewoo           Maker = 19
	MakerDaf              Maker = 20
	MakerDaihatsu         Maker = 21
	MakerDatsun           Maker = 22
	MakerDodge            Maker = 23
	MakerFerrari          Maker = 24
	MakerFiat             Maker = 25
	MakerFord             Maker = 26
	MakerForestRiver      Maker = 27
	MakerGaz              Maker = 28
	MakerGeeli            Maker = 29
	MakerGenesis          Maker = 30
	MakerGenuineScooterCo Maker = 31
	MakerGeo              Maker = 32
	MakerGlobalElectric   Maker = 33
	MakerGmc              Maker = 34
	MakerGreatwall        Maker = 35
	MakerHafei            Maker = 36
	MakerHarleyDavidson   Maker = 37
	MakerHaval            Maker = 38
	MakerHino             Maker = 39
	MakerHonda            Maker = 40
	MakerHummer           Maker = 41
	MakerHyundai          Maker = 42
	MakerInfiniti         Maker = 43
	MakerIsuzu            Maker = 44
	MakerIveco            Maker = 45
	MakerJac              Maker = 46
	MakerJaguar           Maker = 47
	MakerJeep             Maker = 48
	MakerKawasaki         Maker = 49
	MakerKenworth         Maker = 50
	MakerKia              Maker = 51
	MakerLamborghini      Maker = 52
	MakerLancia           Maker = 53
	MakerLandRover        Maker = 54
	MakerLexus            Maker = 55
	MakerLincoln          Maker = 56
	MakerLotus            Maker = 57
	MakerMack             Maker = 58
	MakerMan              Maker = 59
	MakerMaserati         Maker = 60
	MakerMaybach          Maker = 61
	MakerMazda            Maker = 62
	MakerMclaren          Maker = 63
	MakerMercedesBenz     Maker = 64
	MakerMercury          Maker = 65
	MakerMini             Maker = 66
	MakerMitsubishi       Maker = 67
	MakerMoskvich         Maker = 68
	MakerNaz              Maker = 69
	MakerNeoplan          Maker = 70
	MakerNissan           Maker = 71
	MakerOldsmobile       Maker = 72
	MakerOpel             Maker = 73
	MakerPeterbilt        Maker = 74
	MakerPeugeot          Maker = 75
	MakerPlymouth         Maker = 76
	MakerPolaris          Maker = 77
	MakerPontiac          Maker = 78
	MakerPorsche          Maker = 79
	MakerRenault          Maker = 80
	MakerRollsRoyce       Maker = 81
	MakerRover            Maker = 82
	MakerSaab             Maker = 83
	MakerSaturn           Maker = 84
	MakerScion            Maker = 85
	MakerSeat             Maker = 86
	MakerShacman          Maker = 87
	MakerSkoda            Maker = 88
	MakerSmart            Maker = 89
	MakerSpartanMotors    Maker = 90
	MakerSsangyong        Maker = 91
	MakerSubaru           Maker = 92
	MakerSuzuki           Maker = 93
	MakerTata             Maker = 94
	MakerTesla            Maker = 95
	MakerToyota           Maker = 96
	MakerUral             Maker = 97
	MakerVaz              Maker = 98
	MakerVolkswagen       Maker = 99
	MakerVolvo            Maker = 100
	MakerRam              Maker = 101
)

var MakerNames = map[int32]string{
	1:   "ACURA",
	2:   "AGUSTA",
	3:   "ALFA_ROMEO",
	4:   "ASTON_MARTIN",
	5:   "AUDI",
	6:   "BENTLEY",
	7:   "BMW",
	8:   "BUICK",
	9:   "BYD",
	10:  "CADILLAC",
	11:  "CHANGAN",
	12:  "CHANGFENG",
	13:  "CHERY",
	14:  "CHEVROLET",
	15:  "CHRYSLER",
	16:  "CITROEN",
	17:  "CORVETTE",
	18:  "DACIA",
	19:  "DAEWOO",
	20:  "DAF",
	21:  "DAIHATSU",
	22:  "DATSUN",
	23:  "DODGE",
	24:  "FERRARI",
	25:  "FIAT",
	26:  "FORD",
	27:  "FOREST_RIVER",
	28:  "GAZ",
	29:  "GEELI",
	30:  "GENESIS",
	31:  "GENUINE_SCOOTER_CO",
	32:  "GEO",
	33:  "GLOBAL_ELECTRIC",
	34:  "GMC",
	35:  "GREATWALL",
	36:  "HAFEI",
	37:  "HARLEY_DAVIDSON",
	38:  "HAVAL",
	39:  "HINO",
	40:  "HONDA",
	41:  "HUMMER",
	42:  "HYUNDAI",
	43:  "INFINITI",
	44:  "ISUZU",
	45:  "IVECO",
	46:  "JAC",
	47:  "JAGUAR",
	48:  "JEEP",
	49:  "KAWASAKI",
	50:  "KENWORTH",
	51:  "KIA",
	52:  "LAMBORGHINI",
	53:  "LANCIA",
	54:  "LAND_ROVER",
	55:  "LEXUS",
	56:  "LINCOLN",
	57:  "LOTUS",
	58:  "MACK",
	59:  "MAN",
	60:  "MASERATI",
	61:  "MAYBACH",
	62:  "MAZDA",
	63:  "MCLAREN",
	64:  "MERCEDES_BENZ",
	65:  "MERCURY",
	66:  "MINI",
	67:  "MITSUBISHI",
	68:  "MOSKVICH",
	69:  "NAZ",
	70:  "NEOPLAN",
	71:  "NISSAN",
	72:  "OLDSMOBILE",
	73:  "OPEL",
	74:  "PETERBILT",
	75:  "PEUGEOT",
	76:  "PLYMOUTH",
	77:  "POLARIS",
	78:  "PONTIAC",
	79:  "PORSCHE",
	80:  "RENAULT",
	81:  "ROLLS_ROYCE",
	82:  "ROVER",
	83:  "SAAB",
	84:  "SATURN",
	85:  "SCION",
	86:  "SEAT",
	87:  "SHACMAN",
	88:  "SKODA",
	89:  "SMART",
	90:  "SPARTAN_MOTORS",
	91:  "SSANGYONG",
	92:  "SUBARU",
	93:  "SUZUKI",
	94:  "TATA",
	95:  "TESLA",
	96:  "TOYOTA",
	97:  "URAL",
	98:  "VAZ",
	99:  "VOLKSWAGEN",
	100: "VOLVO",
	101: "RAM",
}

var MakerValues = map[string]int32{
	"ACURA":              1,
	"AGUSTA":             2,
	"ALFA_ROMEO":         3,
	"ASTON_MARTIN":       4,
	"AUDI":               5,
	"BENTLEY":            6,
	"BMW":                7,
	"BUICK":              8,
	"BYD":                9,
	"CADILLAC":           10,
	"CHANGAN":            11,
	"CHANGFENG":          12,
	"CHERY":              13,
	"CHEVROLET":          14,
	"CHRYSLER":           15,
	"CITROEN":            16,
	"CORVETTE":           17,
	"DACIA":              18,
	"DAEWOO":             19,
	"DAF":                20,
	"DAIHATSU":           21,
	"DATSUN":             22,
	"DODGE":              23,
	"FERRARI":            24,
	"FIAT":               25,
	"FORD":               26,
	"FOREST_RIVER":       27,
	"GAZ":                28,
	"GEELI":              29,
	"GENESIS":            30,
	"GENUINE_SCOOTER_CO": 31,
	"GEO":                32,
	"GLOBAL_ELECTRIC":    33,
	"GMC":                34,
	"GREATWALL":          35,
	"HAFEI":              36,
	"HARLEY_DAVIDSON":    37,
	"HAVAL":              38,
	"HINO":               39,
	"HONDA":              40,
	"HUMMER":             41,
	"HYUNDAI":            42,
	"INFINITI":           43,
	"ISUZU":              44,
	"IVECO":              45,
	"JAC":                46,
	"JAGUAR":             47,
	"JEEP":               48,
	"KAWASAKI":           49,
	"KENWORTH":           50,
	"KIA":                51,
	"LAMBORGHINI":        52,
	"LANCIA":             53,
	"LAND_ROVER":         54,
	"LEXUS":              55,
	"LINCOLN":            56,
	"LOTUS":              57,
	"MACK":               58,
	"MAN":                59,
	"MASERATI":           60,
	"MAYBACH":            61,
	"MAZDA":              62,
	"MCLAREN":            63,
	"MERCEDES_BENZ":      64,
	"MERCURY":            65,
	"MINI":               66,
	"MITSUBISHI":         67,
	"MOSKVICH":           68,
	"NAZ":                69,
	"NEOPLAN":            70,
	"NISSAN":             71,
	"OLDSMOBILE":         72,
	"OPEL":               73,
	"PETERBILT":          74,
	"PEUGEOT":            75,
	"PLYMOUTH":           76,
	"POLARIS":            77,
	"PONTIAC":            78,
	"PORSCHE":            79,
	"RENAULT":            80,
	"ROLLS_ROYCE":        81,
	"ROVER":              82,
	"SAAB":               83,
	"SATURN":             84,
	"SCION":              85,
	"SEAT":               86,
	"SHACMAN":            87,
	"SKODA":              88,
	"SMART":              89,
	"SPARTAN_MOTORS":     90,
	"SSANGYONG":          91,
	"SUBARU":             92,
	"SUZUKI":             93,
	"TATA":               94,
	"TESLA":              95,
	"TOYOTA":             96,
	"URAL":               97,
	"VAZ":                98,
	"VOLKSWAGEN":         99,
	"VOLVO":              100,
	"RAM":                101,
}

// MakerAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var MakerAliases = []state.Alias{
	{Alias: "^acura$", Value: 1},
	{Alias: "^agusta$", Value: 2},
	{Alias: "^alfaromeo$", Value: 3},
	{Alias: "^astonmartin$", Value: 4},
	{Alias: "^audi$", Value: 5},
	{Alias: "^bentley$", Value: 6},
	{Alias: "^bmw$", Value: 7},
	{Alias: "^buick$", Value: 8},
	{Alias: "^byd$", Value: 9},
	{Alias: "^cadillac$", Value: 10},
	{Alias: "^changan$", Value: 11},
	{Alias: "^changfeng$", Value: 12},
	{Alias: "^chery$", Value: 13},
	{Alias: "^chevrolet$", Value: 14},
	{Alias: "^chrysler$", Value: 15},
	{Alias: "^citroen$", Value: 16},
	{Alias: "^corvette$", Value: 17},
	{Alias: "^dacia$", Value: 18},
	{Alias: "^daewoo$", Value: 19},
	{Alias: "^daf$", Value: 20},
	{Alias: "^daihatsu$", Value: 21},
	{Alias: "^datsun$", Value: 22},
	{Alias: "^dodge$", Value: 23},
	{Alias: "^ferrari$", Value: 24},
	{Alias: "^fiat$", Value: 25},
	{Alias: "^ford$", Value: 26},
	{Alias: "^forestriver$", Value: 27},
	{Alias: "^gaz$", Value: 28},
	{Alias: "^geeli$", Value: 29},
	{Alias: "^genesis$", Value: 30},
	{Alias: "^genuinescooterco$", Value: 31},
	{Alias: "^geo$", Value: 32},
	{Alias: "^globalelectricmotors|globalelectric$", Value: 33},
	{Alias: "^generalmotors|gmc$", Value: 34},
	{Alias: "^greatwall$", Value: 35},
	{Alias: "^hafei$", Value: 36},
	{Alias: "^harleydavidson$", Value: 37},
	{Alias: "^haval$", Value: 38},
	{Alias: "^hino$", Value: 39},
	{Alias: "^honda$", Value: 40},
	{Alias: "^hummer$", Value: 41},
	{Alias: "^hyundai$", Value: 42},
	{Alias: "^infiniti$", Value: 43},
	{Alias: "^isuzu$", Value: 44},
	{Alias: "^iveco$", Value: 45},
	{Alias: "^jac$", Value: 46},
	{Alias: "^jaguar$", Value: 47},
	{Alias: "^jeep$", Value: 48},
	{Alias: "^kawasaki$", Value: 49},
	{Alias: "^kenworth$", Value: 50},
	{Alias: "^kia$", Value: 51},
	{Alias: "^lamborghini$", Value: 52},
	{Alias: "^lancia$", Value: 53},
	{Alias: "^landrover$", Value: 54},
	{Alias: "^lexus$", Value: 55},
	{Alias: "^lincoln$", Value: 56},
	{Alias: "^lotus$", Value: 57},
	{Alias: "^mack$", Value: 58},
	{Alias: "^man$", Value: 59},
	{Alias: "^maserati$", Value: 60},
	{Alias: "^maybach$", Value: 61},
	{Alias: "^mazda$", Value: 62},
	{Alias: "^mclaren|mclarenautomotive$", Value: 63},
	{Alias: "^mercedes|mercedesbenz$", Value: 64},
	{Alias: "^mercury$", Value: 65},
	{Alias: "^mini$", Value: 66},
	{Alias: "^mitsubishi$", Value: 67},
	{Alias: "^moskvich$", Value: 68},
	{Alias: "^naz$", Value: 69},
	{Alias: "^neoplan$", Value: 70},
	{Alias: "^nissan$", Value: 71},
	{Alias: "^oldsmobile$", Value: 72},
	{Alias: "^opel$", Value: 73},
	{Alias: "^peterbilt$", Value: 74},
	{Alias: "^peugeot$", Value: 75},
	{Alias: "^plymouth$", Value: 76},
	{Alias: "^polaris$", Value: 77},
	{Alias: "^pontiac$", Value: 78},
	{Alias: "^porsche$", Value: 79},
	{Alias: "^renault$", Value: 80},
	{Alias: "^rollsroyce$", Value: 81},
	{Alias: "^rover$", Value: 82},
	{Alias: "^saab$", Value: 83},
	{Alias: "^saturn$", Value: 84},
	{Alias: "^scion$", Value: 85},
	{Alias: "^seat$", Value: 86},
	{Alias: "^shacman$", Value: 87},
	{Alias: "^skoda$", Value: 88},
	{Alias: "^smart$", Value: 89},
	{Alias: "^spartanmotors$", Value: 90},
	{Alias: "^ssangyong$", Value: 91},
	{Alias: "^subaru$", Value: 92},
	{Alias: "^suzuki$", Value: 93},
	{Alias: "^tata$", Value: 94},
	{Alias: "^tesla$", Value: 95},
	{Alias: "^toyota$", Value: 96},
	{Alias: "^ural$", Value: 97},
	{Alias: "^vaz$", Value: 98},
	{Alias: "^volkswagen$", Value: 99},
	{Alias: "^volvo$", Value: 100},
	{Alias: "^ram$", Value: 101},
}

// MakerOrderedValues is a list of sorted values
// of Maker. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var MakerOrderedValues = []int32{
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
	48,
	49,
	50,
	51,
	52,
	53,
	54,
	55,
	56,
	57,
	58,
	59,
	60,
	61,
	62,
	63,
	64,
	65,
	66,
	67,
	68,
	69,
	70,
	71,
	72,
	73,
	74,
	75,
	76,
	77,
	78,
	79,
	80,
	81,
	82,
	83,
	84,
	85,
	86,
	87,
	88,
	89,
	90,
	91,
	92,
	93,
	94,
	95,
	96,
	97,
	98,
	99,
	100,
	101,
}
