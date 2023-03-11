package static

import "gitlab.com/eAuction/enumer/pkg/state"

type Class int32

const (
	ClassBmw1er          Class = 1
	ClassBmw2er          Class = 2
	ClassBmw3er          Class = 3
	ClassBmw4er          Class = 4
	ClassBmw5er          Class = 5
	ClassBmw6er          Class = 6
	ClassBmw7er          Class = 7
	ClassBmw8er          Class = 8
	ClassBmwISeries      Class = 9
	ClassBmwMSeries      Class = 10
	ClassBmwXSeries      Class = 11
	ClassBmwZSeries      Class = 12
	ClassFordF           Class = 13
	ClassGmcCSeries      Class = 14
	ClassLexusEs         Class = 15
	ClassLexusGs         Class = 16
	ClassLexusGx         Class = 17
	ClassLexusIs         Class = 18
	ClassLexusLs         Class = 19
	ClassLexusLx         Class = 20
	ClassLexusNx         Class = 21
	ClassLexusRc         Class = 22
	ClassLexusRx         Class = 23
	ClassLexusSc         Class = 24
	ClassMercedesBenzA   Class = 25
	ClassMercedesBenzB   Class = 26
	ClassMercedesBenzC   Class = 27
	ClassMercedesBenzCl  Class = 28
	ClassMercedesBenzCla Class = 29
	ClassMercedesBenzClk Class = 30
	ClassMercedesBenzCls Class = 31
	ClassMercedesBenzE   Class = 32
	ClassMercedesBenzG   Class = 33
	ClassMercedesBenzGl  Class = 34
	ClassMercedesBenzGla Class = 35
	ClassMercedesBenzGlc Class = 36
	ClassMercedesBenzGle Class = 37
	ClassMercedesBenzGlk Class = 38
	ClassMercedesBenzR   Class = 39
	ClassMercedesBenzGls Class = 40
	ClassMercedesBenzMl  Class = 41
	ClassMercedesBenzS   Class = 42
	ClassMercedesBenzSlr Class = 43
	ClassMercedesBenzSl  Class = 44
	ClassMercedesBenzSlc Class = 45
	ClassMercedesBenzSlk Class = 46
	ClassMercedesBenzV   Class = 47
	ClassMercedesBenzX   Class = 48
	ClassLexusLc         Class = 49
	ClassPorsche718      Class = 50
	ClassPorsche911      Class = 51
	ClassPorscheTaycan   Class = 52
	ClassPorschePanamera Class = 53
	ClassPorscheMacan    Class = 54
	ClassPorscheCayenne  Class = 55
)

var ClassNames = map[int32]string{
	1:  "BMW_1ER",
	2:  "BMW_2ER",
	3:  "BMW_3ER",
	4:  "BMW_4ER",
	5:  "BMW_5ER",
	6:  "BMW_6ER",
	7:  "BMW_7ER",
	8:  "BMW_8ER",
	9:  "BMW_I_SERIES",
	10: "BMW_M_SERIES",
	11: "BMW_X_SERIES",
	12: "BMW_Z_SERIES",
	13: "FORD_F",
	14: "GMC_C_SERIES",
	15: "LEXUS_ES",
	16: "LEXUS_GS",
	17: "LEXUS_GX",
	18: "LEXUS_IS",
	19: "LEXUS_LS",
	20: "LEXUS_LX",
	21: "LEXUS_NX",
	22: "LEXUS_RC",
	23: "LEXUS_RX",
	24: "LEXUS_SC",
	25: "MERCEDES_BENZ_A",
	26: "MERCEDES_BENZ_B",
	27: "MERCEDES_BENZ_C",
	28: "MERCEDES_BENZ_CL",
	29: "MERCEDES_BENZ_CLA",
	30: "MERCEDES_BENZ_CLK",
	31: "MERCEDES_BENZ_CLS",
	32: "MERCEDES_BENZ_E",
	33: "MERCEDES_BENZ_G",
	34: "MERCEDES_BENZ_GL",
	35: "MERCEDES_BENZ_GLA",
	36: "MERCEDES_BENZ_GLC",
	37: "MERCEDES_BENZ_GLE",
	38: "MERCEDES_BENZ_GLK",
	39: "MERCEDES_BENZ_R",
	40: "MERCEDES_BENZ_GLS",
	41: "MERCEDES_BENZ_ML",
	42: "MERCEDES_BENZ_S",
	43: "MERCEDES_BENZ_SLR",
	44: "MERCEDES_BENZ_SL",
	45: "MERCEDES_BENZ_SLC",
	46: "MERCEDES_BENZ_SLK",
	47: "MERCEDES_BENZ_V",
	48: "MERCEDES_BENZ_X",
	49: "LEXUS_LC",
	50: "PORSCHE_718",
	51: "PORSCHE_911",
	52: "PORSCHE_TAYCAN",
	53: "PORSCHE_PANAMERA",
	54: "PORSCHE_MACAN",
	55: "PORSCHE_CAYENNE",
}

var ClassValues = map[string]int32{
	"BMW_1ER":           1,
	"BMW_2ER":           2,
	"BMW_3ER":           3,
	"BMW_4ER":           4,
	"BMW_5ER":           5,
	"BMW_6ER":           6,
	"BMW_7ER":           7,
	"BMW_8ER":           8,
	"BMW_I_SERIES":      9,
	"BMW_M_SERIES":      10,
	"BMW_X_SERIES":      11,
	"BMW_Z_SERIES":      12,
	"FORD_F":            13,
	"GMC_C_SERIES":      14,
	"LEXUS_ES":          15,
	"LEXUS_GS":          16,
	"LEXUS_GX":          17,
	"LEXUS_IS":          18,
	"LEXUS_LS":          19,
	"LEXUS_LX":          20,
	"LEXUS_NX":          21,
	"LEXUS_RC":          22,
	"LEXUS_RX":          23,
	"LEXUS_SC":          24,
	"MERCEDES_BENZ_A":   25,
	"MERCEDES_BENZ_B":   26,
	"MERCEDES_BENZ_C":   27,
	"MERCEDES_BENZ_CL":  28,
	"MERCEDES_BENZ_CLA": 29,
	"MERCEDES_BENZ_CLK": 30,
	"MERCEDES_BENZ_CLS": 31,
	"MERCEDES_BENZ_E":   32,
	"MERCEDES_BENZ_G":   33,
	"MERCEDES_BENZ_GL":  34,
	"MERCEDES_BENZ_GLA": 35,
	"MERCEDES_BENZ_GLC": 36,
	"MERCEDES_BENZ_GLE": 37,
	"MERCEDES_BENZ_GLK": 38,
	"MERCEDES_BENZ_R":   39,
	"MERCEDES_BENZ_GLS": 40,
	"MERCEDES_BENZ_ML":  41,
	"MERCEDES_BENZ_S":   42,
	"MERCEDES_BENZ_SLR": 43,
	"MERCEDES_BENZ_SL":  44,
	"MERCEDES_BENZ_SLC": 45,
	"MERCEDES_BENZ_SLK": 46,
	"MERCEDES_BENZ_V":   47,
	"MERCEDES_BENZ_X":   48,
	"LEXUS_LC":          49,
	"PORSCHE_718":       50,
	"PORSCHE_911":       51,
	"PORSCHE_TAYCAN":    52,
	"PORSCHE_PANAMERA":  53,
	"PORSCHE_MACAN":     54,
	"PORSCHE_CAYENNE":   55,
}

// ClassAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var ClassAliases = []state.Alias{
	{Alias: "^1er$", Value: 1},
	{Alias: "^2er$", Value: 2},
	{Alias: "^3er$", Value: 3},
	{Alias: "^4er$", Value: 4},
	{Alias: "^5er$", Value: 5},
	{Alias: "^6er$", Value: 6},
	{Alias: "^7er$", Value: 7},
	{Alias: "^8er$", Value: 8},
	{Alias: "^iseries$", Value: 9},
	{Alias: "^mseries$", Value: 10},
	{Alias: "^xseries$", Value: 11},
	{Alias: "^zseries$", Value: 12},
	{Alias: "^f$", Value: 13},
	{Alias: "^cseries$", Value: 14},
	{Alias: "^es$", Value: 15},
	{Alias: "^gs$", Value: 16},
	{Alias: "^gx$", Value: 17},
	{Alias: "^is$", Value: 18},
	{Alias: "^ls$", Value: 19},
	{Alias: "^lx$", Value: 20},
	{Alias: "^nx$", Value: 21},
	{Alias: "^rc$", Value: 22},
	{Alias: "^rx$", Value: 23},
	{Alias: "^sc$", Value: 24},
	{Alias: "^a$", Value: 25},
	{Alias: "^b$", Value: 26},
	{Alias: "^c$", Value: 27},
	{Alias: "^cl$", Value: 28},
	{Alias: "^cla$", Value: 29},
	{Alias: "^clk$", Value: 30},
	{Alias: "^cls$", Value: 31},
	{Alias: "^e$", Value: 32},
	{Alias: "^g$", Value: 33},
	{Alias: "^gl$", Value: 34},
	{Alias: "^gla$", Value: 35},
	{Alias: "^glc$", Value: 36},
	{Alias: "^gle$", Value: 37},
	{Alias: "^glk$", Value: 38},
	{Alias: "^r$", Value: 39},
	{Alias: "^gls$", Value: 40},
	{Alias: "^ml$", Value: 41},
	{Alias: "^s$", Value: 42},
	{Alias: "^slr$", Value: 43},
	{Alias: "^sl$", Value: 44},
	{Alias: "^slc$", Value: 45},
	{Alias: "^slk$", Value: 46},
	{Alias: "^v$", Value: 47},
	{Alias: "^x$", Value: 48},
	{Alias: "^lc$", Value: 49},
	{Alias: "^718$", Value: 50},
	{Alias: "^911$", Value: 51},
	{Alias: "^taycan$", Value: 52},
	{Alias: "^panamera$", Value: 53},
	{Alias: "^macan$", Value: 54},
	{Alias: "^cayenne$", Value: 55},
}

// ClassOrderedValues is a list of sorted values
// of Class. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var ClassOrderedValues = []int32{
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
}
