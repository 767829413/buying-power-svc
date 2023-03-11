package static

import "gitlab.com/eAuction/enumer/pkg/state"

type Region int32

const (
	RegionAl Region = 1
	RegionAk Region = 2
	RegionAz Region = 3
	RegionAr Region = 4
	RegionCa Region = 5
	RegionCo Region = 6
	RegionCt Region = 7
	RegionDe Region = 8
	RegionFl Region = 9
	RegionGa Region = 10
	RegionHi Region = 11
	RegionId Region = 12
	RegionIl Region = 13
	RegionIn Region = 14
	RegionIa Region = 15
	RegionKs Region = 16
	RegionKy Region = 17
	RegionLa Region = 18
	RegionMe Region = 19
	RegionMd Region = 20
	RegionMa Region = 21
	RegionMi Region = 22
	RegionMn Region = 23
	RegionMs Region = 24
	RegionMo Region = 25
	RegionMt Region = 26
	RegionNe Region = 27
	RegionNv Region = 28
	RegionNh Region = 29
	RegionNj Region = 30
	RegionNm Region = 31
	RegionNy Region = 32
	RegionNc Region = 33
	RegionNd Region = 34
	RegionOh Region = 35
	RegionOk Region = 36
	RegionOr Region = 37
	RegionPa Region = 38
	RegionRi Region = 39
	RegionSc Region = 40
	RegionSd Region = 41
	RegionTn Region = 42
	RegionTx Region = 43
	RegionUt Region = 44
	RegionVt Region = 45
	RegionVa Region = 46
	RegionWa Region = 47
	RegionWv Region = 48
	RegionWi Region = 49
	RegionWy Region = 50
	RegionDc Region = 51
	RegionVi Region = 52
	RegionPr Region = 53
	RegionGu Region = 54
)

var RegionNames = map[int32]string{
	1:  "AL",
	2:  "AK",
	3:  "AZ",
	4:  "AR",
	5:  "CA",
	6:  "CO",
	7:  "CT",
	8:  "DE",
	9:  "FL",
	10: "GA",
	11: "HI",
	12: "ID",
	13: "IL",
	14: "IN",
	15: "IA",
	16: "KS",
	17: "KY",
	18: "LA",
	19: "ME",
	20: "MD",
	21: "MA",
	22: "MI",
	23: "MN",
	24: "MS",
	25: "MO",
	26: "MT",
	27: "NE",
	28: "NV",
	29: "NH",
	30: "NJ",
	31: "NM",
	32: "NY",
	33: "NC",
	34: "ND",
	35: "OH",
	36: "OK",
	37: "OR",
	38: "PA",
	39: "RI",
	40: "SC",
	41: "SD",
	42: "TN",
	43: "TX",
	44: "UT",
	45: "VT",
	46: "VA",
	47: "WA",
	48: "WV",
	49: "WI",
	50: "WY",
	51: "DC",
	52: "VI",
	53: "PR",
	54: "GU",
}

var RegionValues = map[string]int32{
	"AL": 1,
	"AK": 2,
	"AZ": 3,
	"AR": 4,
	"CA": 5,
	"CO": 6,
	"CT": 7,
	"DE": 8,
	"FL": 9,
	"GA": 10,
	"HI": 11,
	"ID": 12,
	"IL": 13,
	"IN": 14,
	"IA": 15,
	"KS": 16,
	"KY": 17,
	"LA": 18,
	"ME": 19,
	"MD": 20,
	"MA": 21,
	"MI": 22,
	"MN": 23,
	"MS": 24,
	"MO": 25,
	"MT": 26,
	"NE": 27,
	"NV": 28,
	"NH": 29,
	"NJ": 30,
	"NM": 31,
	"NY": 32,
	"NC": 33,
	"ND": 34,
	"OH": 35,
	"OK": 36,
	"OR": 37,
	"PA": 38,
	"RI": 39,
	"SC": 40,
	"SD": 41,
	"TN": 42,
	"TX": 43,
	"UT": 44,
	"VT": 45,
	"VA": 46,
	"WA": 47,
	"WV": 48,
	"WI": 49,
	"WY": 50,
	"DC": 51,
	"VI": 52,
	"PR": 53,
	"GU": 54,
}

// RegionAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var RegionAliases = []state.Alias{
	{Alias: "^(alabama|al)$", Value: 1},
	{Alias: "^(alaska|ak)$", Value: 2},
	{Alias: "^(arizona|az)$", Value: 3},
	{Alias: "^(arkansas|ar)$", Value: 4},
	{Alias: "^(california|ca)$", Value: 5},
	{Alias: "^(colorado|co)$", Value: 6},
	{Alias: "^(connecticut|ct)$", Value: 7},
	{Alias: "^(delaware|de)$", Value: 8},
	{Alias: "^(florida|fl)$", Value: 9},
	{Alias: "^(georgia|ga)$", Value: 10},
	{Alias: "^(hawaii|hi)$", Value: 11},
	{Alias: "^(idaho|id)$", Value: 12},
	{Alias: "^(illinois|il)$", Value: 13},
	{Alias: "^(indiana|in)$", Value: 14},
	{Alias: "^(iowa|ia)$", Value: 15},
	{Alias: "^(kansas|ks)$", Value: 16},
	{Alias: "^(kentucky|ky)$", Value: 17},
	{Alias: "^(louisiana|la)$", Value: 18},
	{Alias: "^(maine|me)$", Value: 19},
	{Alias: "^(maryland|md)$", Value: 20},
	{Alias: "^(massachusetts|ma)$", Value: 21},
	{Alias: "^(michigan|mi)$", Value: 22},
	{Alias: "^(minnesota|mn)$", Value: 23},
	{Alias: "^(mississippi|ms)$", Value: 24},
	{Alias: "^(missouri|mo)$", Value: 25},
	{Alias: "^(montana|mt)$", Value: 26},
	{Alias: "^(nebraska|ne)$", Value: 27},
	{Alias: "^(nevada|nv)$", Value: 28},
	{Alias: "^(newhampshire|nh)$", Value: 29},
	{Alias: "^(newjersey|nj)$", Value: 30},
	{Alias: "^(newmexico|nm)$", Value: 31},
	{Alias: "^(newyork|ny)$", Value: 32},
	{Alias: "^(northcarolina|nc)$", Value: 33},
	{Alias: "^(northdakota|nd)$", Value: 34},
	{Alias: "^(ohio|oh)$", Value: 35},
	{Alias: "^(oklahoma|ok)$", Value: 36},
	{Alias: "^(oregon|or)$", Value: 37},
	{Alias: "^(pennsylvania|pa)$", Value: 38},
	{Alias: "^(rhodeisland|ri)$", Value: 39},
	{Alias: "^(southcarolina|sc)$", Value: 40},
	{Alias: "^(southdakota|sd)$", Value: 41},
	{Alias: "^(tennessee|tn)$", Value: 42},
	{Alias: "^(texas|tx)$", Value: 43},
	{Alias: "^(utah|ut)$", Value: 44},
	{Alias: "^(vermont|vt)$", Value: 45},
	{Alias: "^(virginia|va)$", Value: 46},
	{Alias: "^(washington|wa)$", Value: 47},
	{Alias: "^(westvirginia|wv)$", Value: 48},
	{Alias: "^(wisconsin|wi)$", Value: 49},
	{Alias: "^(wyoming|wy)$", Value: 50},
	{Alias: "^(districtofcolumbia|dc|washingtondc)$", Value: 51},
	{Alias: "^(virginislands|vi)$", Value: 52},
	{Alias: "^(puertorico|pr)$", Value: 53},
	{Alias: "^(guam|gu)$", Value: 54},
}

// RegionOrderedValues is a list of sorted values
// of Region. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var RegionOrderedValues = []int32{
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
}
