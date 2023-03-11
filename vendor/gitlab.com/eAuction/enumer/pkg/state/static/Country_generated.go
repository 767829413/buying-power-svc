package static

import "gitlab.com/eAuction/enumer/pkg/state"

type Country int32

const (
	CountryOther Country = 1
	CountryAfg   Country = 2
	CountryAlb   Country = 3
	CountryDza   Country = 4
	CountryAsm   Country = 5
	CountryAnd   Country = 6
	CountryAgo   Country = 7
	CountryAia   Country = 8
	CountryAta   Country = 9
	CountryAtg   Country = 10
	CountryArg   Country = 11
	CountryArm   Country = 12
	CountryAbw   Country = 13
	CountryAus   Country = 14
	CountryAut   Country = 15
	CountryAze   Country = 16
	CountryBhs   Country = 17
	CountryBhr   Country = 18
	CountryBgd   Country = 19
	CountryBrb   Country = 20
	CountryBlr   Country = 21
	CountryBel   Country = 22
	CountryBlz   Country = 23
	CountryBen   Country = 24
	CountryBmu   Country = 25
	CountryBtn   Country = 26
	CountryBol   Country = 27
	CountryBes   Country = 28
	CountryBih   Country = 29
	CountryBwa   Country = 30
	CountryBvt   Country = 31
	CountryBra   Country = 32
	CountryIot   Country = 33
	CountryBrn   Country = 34
	CountryBgr   Country = 35
	CountryBfa   Country = 36
	CountryBdi   Country = 37
	CountryKhm   Country = 38
	CountryCmr   Country = 39
	CountryCan   Country = 40
	CountryCpv   Country = 41
	CountryCym   Country = 42
	CountryCaf   Country = 43
	CountryTcd   Country = 44
	CountryChl   Country = 45
	CountryChn   Country = 46
	CountryCxr   Country = 47
	CountryCck   Country = 48
	CountryCol   Country = 49
	CountryCom   Country = 50
	CountryCog   Country = 51
	CountryCod   Country = 52
	CountryCok   Country = 53
	CountryCri   Country = 54
	CountryHrv   Country = 55
	CountryCub   Country = 56
	CountryCuw   Country = 57
	CountryCyp   Country = 58
	CountryCze   Country = 59
	CountryCiv   Country = 60
	CountryDnk   Country = 61
	CountryDji   Country = 62
	CountryDma   Country = 63
	CountryDom   Country = 64
	CountryEcu   Country = 65
	CountryEgy   Country = 66
	CountrySlv   Country = 67
	CountryGnq   Country = 68
	CountryEri   Country = 69
	CountryEst   Country = 70
	CountryEth   Country = 71
	CountryFlk   Country = 72
	CountryFro   Country = 73
	CountryFji   Country = 74
	CountryFin   Country = 75
	CountryFra   Country = 76
	CountryGuf   Country = 77
	CountryPyf   Country = 78
	CountryAtf   Country = 79
	CountryGab   Country = 80
	CountryGmb   Country = 81
	CountryGeo   Country = 82
	CountryDeu   Country = 83
	CountryGha   Country = 84
	CountryGib   Country = 85
	CountryGrc   Country = 86
	CountryGrl   Country = 87
	CountryGrd   Country = 88
	CountryGlp   Country = 89
	CountryGum   Country = 90
	CountryGtm   Country = 91
	CountryGgy   Country = 92
	CountryGin   Country = 93
	CountryGnb   Country = 94
	CountryGuy   Country = 95
	CountryHti   Country = 96
	CountryHmd   Country = 97
	CountryVat   Country = 98
	CountryHnd   Country = 99
	CountryHkg   Country = 100
	CountryHun   Country = 101
	CountryIsl   Country = 102
	CountryInd   Country = 103
	CountryIdn   Country = 104
	CountryIrn   Country = 105
	CountryIrq   Country = 106
	CountryIrl   Country = 107
	CountryImn   Country = 108
	CountryIsr   Country = 109
	CountryIta   Country = 110
	CountryJam   Country = 111
	CountryJpn   Country = 112
	CountryJey   Country = 113
	CountryJor   Country = 114
	CountryKaz   Country = 115
	CountryKen   Country = 116
	CountryKir   Country = 117
	CountryPrk   Country = 118
	CountryKor   Country = 119
	CountryKwt   Country = 120
	CountryKgz   Country = 121
	CountryLao   Country = 122
	CountryLva   Country = 123
	CountryLbn   Country = 124
	CountryLso   Country = 125
	CountryLbr   Country = 126
	CountryLby   Country = 127
	CountryLie   Country = 128
	CountryLtu   Country = 129
	CountryLux   Country = 130
	CountryMac   Country = 131
	CountryMkd   Country = 132
	CountryMdg   Country = 133
	CountryMwi   Country = 134
	CountryMys   Country = 135
	CountryMdv   Country = 136
	CountryMli   Country = 137
	CountryMlt   Country = 138
	CountryMhl   Country = 139
	CountryMtq   Country = 140
	CountryMrt   Country = 141
	CountryMus   Country = 142
	CountryMyt   Country = 143
	CountryMex   Country = 144
	CountryFsm   Country = 145
	CountryMda   Country = 146
	CountryMco   Country = 147
	CountryMng   Country = 148
	CountryMne   Country = 149
	CountryMsr   Country = 150
	CountryMar   Country = 151
	CountryMoz   Country = 152
	CountryMmr   Country = 153
	CountryNam   Country = 154
	CountryNru   Country = 155
	CountryNpl   Country = 156
	CountryNld   Country = 157
	CountryNcl   Country = 158
	CountryNzl   Country = 159
	CountryNic   Country = 160
	CountryNer   Country = 161
	CountryNga   Country = 162
	CountryNiu   Country = 163
	CountryNfk   Country = 164
	CountryMnp   Country = 165
	CountryNor   Country = 166
	CountryOmn   Country = 167
	CountryPak   Country = 168
	CountryPlw   Country = 169
	CountryPse   Country = 170
	CountryPan   Country = 171
	CountryPng   Country = 172
	CountryPry   Country = 173
	CountryPer   Country = 174
	CountryPhl   Country = 175
	CountryPcn   Country = 176
	CountryPol   Country = 177
	CountryPrt   Country = 178
	CountryPri   Country = 179
	CountryQat   Country = 180
	CountryRou   Country = 181
	CountryRus   Country = 182
	CountryRwa   Country = 183
	CountryReu   Country = 184
	CountryBlm   Country = 185
	CountryShn   Country = 186
	CountryKna   Country = 187
	CountryLca   Country = 188
	CountryMaf   Country = 189
	CountrySpm   Country = 190
	CountryVct   Country = 191
	CountryWsm   Country = 192
	CountrySmr   Country = 193
	CountryStp   Country = 194
	CountrySau   Country = 195
	CountrySen   Country = 196
	CountrySrb   Country = 197
	CountrySyc   Country = 198
	CountrySle   Country = 199
	CountrySgp   Country = 200
	CountrySxm   Country = 201
	CountrySvk   Country = 202
	CountrySvn   Country = 203
	CountrySlb   Country = 204
	CountrySom   Country = 205
	CountryZaf   Country = 206
	CountrySgs   Country = 207
	CountrySsd   Country = 208
	CountryEsp   Country = 209
	CountryLka   Country = 210
	CountrySdn   Country = 211
	CountrySur   Country = 212
	CountrySjm   Country = 213
	CountrySwz   Country = 214
	CountrySwe   Country = 215
	CountryChe   Country = 216
	CountrySyr   Country = 217
	CountryTwn   Country = 218
	CountryTjk   Country = 219
	CountryTza   Country = 220
	CountryTha   Country = 221
	CountryTls   Country = 222
	CountryTgo   Country = 223
	CountryTkl   Country = 224
	CountryTon   Country = 225
	CountryTto   Country = 226
	CountryTun   Country = 227
	CountryTur   Country = 228
	CountryTkm   Country = 229
	CountryTca   Country = 230
	CountryTuv   Country = 231
	CountryUga   Country = 232
	CountryUkr   Country = 233
	CountryAre   Country = 234
	CountryGbr   Country = 235
	CountryUsa   Country = 236
	CountryUmi   Country = 237
	CountryUry   Country = 238
	CountryUzb   Country = 239
	CountryVut   Country = 240
	CountryVen   Country = 241
	CountryVnm   Country = 242
	CountryVgb   Country = 243
	CountryVir   Country = 244
	CountryWlf   Country = 245
	CountryEsh   Country = 246
	CountryYem   Country = 247
	CountryZmb   Country = 248
	CountryZwe   Country = 249
)

var CountryNames = map[int32]string{
	1:   "OTHER",
	2:   "AFG",
	3:   "ALB",
	4:   "DZA",
	5:   "ASM",
	6:   "AND",
	7:   "AGO",
	8:   "AIA",
	9:   "ATA",
	10:  "ATG",
	11:  "ARG",
	12:  "ARM",
	13:  "ABW",
	14:  "AUS",
	15:  "AUT",
	16:  "AZE",
	17:  "BHS",
	18:  "BHR",
	19:  "BGD",
	20:  "BRB",
	21:  "BLR",
	22:  "BEL",
	23:  "BLZ",
	24:  "BEN",
	25:  "BMU",
	26:  "BTN",
	27:  "BOL",
	28:  "BES",
	29:  "BIH",
	30:  "BWA",
	31:  "BVT",
	32:  "BRA",
	33:  "IOT",
	34:  "BRN",
	35:  "BGR",
	36:  "BFA",
	37:  "BDI",
	38:  "KHM",
	39:  "CMR",
	40:  "CAN",
	41:  "CPV",
	42:  "CYM",
	43:  "CAF",
	44:  "TCD",
	45:  "CHL",
	46:  "CHN",
	47:  "CXR",
	48:  "CCK",
	49:  "COL",
	50:  "COM",
	51:  "COG",
	52:  "COD",
	53:  "COK",
	54:  "CRI",
	55:  "HRV",
	56:  "CUB",
	57:  "CUW",
	58:  "CYP",
	59:  "CZE",
	60:  "CIV",
	61:  "DNK",
	62:  "DJI",
	63:  "DMA",
	64:  "DOM",
	65:  "ECU",
	66:  "EGY",
	67:  "SLV",
	68:  "GNQ",
	69:  "ERI",
	70:  "EST",
	71:  "ETH",
	72:  "FLK",
	73:  "FRO",
	74:  "FJI",
	75:  "FIN",
	76:  "FRA",
	77:  "GUF",
	78:  "PYF",
	79:  "ATF",
	80:  "GAB",
	81:  "GMB",
	82:  "GEO",
	83:  "DEU",
	84:  "GHA",
	85:  "GIB",
	86:  "GRC",
	87:  "GRL",
	88:  "GRD",
	89:  "GLP",
	90:  "GUM",
	91:  "GTM",
	92:  "GGY",
	93:  "GIN",
	94:  "GNB",
	95:  "GUY",
	96:  "HTI",
	97:  "HMD",
	98:  "VAT",
	99:  "HND",
	100: "HKG",
	101: "HUN",
	102: "ISL",
	103: "IND",
	104: "IDN",
	105: "IRN",
	106: "IRQ",
	107: "IRL",
	108: "IMN",
	109: "ISR",
	110: "ITA",
	111: "JAM",
	112: "JPN",
	113: "JEY",
	114: "JOR",
	115: "KAZ",
	116: "KEN",
	117: "KIR",
	118: "PRK",
	119: "KOR",
	120: "KWT",
	121: "KGZ",
	122: "LAO",
	123: "LVA",
	124: "LBN",
	125: "LSO",
	126: "LBR",
	127: "LBY",
	128: "LIE",
	129: "LTU",
	130: "LUX",
	131: "MAC",
	132: "MKD",
	133: "MDG",
	134: "MWI",
	135: "MYS",
	136: "MDV",
	137: "MLI",
	138: "MLT",
	139: "MHL",
	140: "MTQ",
	141: "MRT",
	142: "MUS",
	143: "MYT",
	144: "MEX",
	145: "FSM",
	146: "MDA",
	147: "MCO",
	148: "MNG",
	149: "MNE",
	150: "MSR",
	151: "MAR",
	152: "MOZ",
	153: "MMR",
	154: "NAM",
	155: "NRU",
	156: "NPL",
	157: "NLD",
	158: "NCL",
	159: "NZL",
	160: "NIC",
	161: "NER",
	162: "NGA",
	163: "NIU",
	164: "NFK",
	165: "MNP",
	166: "NOR",
	167: "OMN",
	168: "PAK",
	169: "PLW",
	170: "PSE",
	171: "PAN",
	172: "PNG",
	173: "PRY",
	174: "PER",
	175: "PHL",
	176: "PCN",
	177: "POL",
	178: "PRT",
	179: "PRI",
	180: "QAT",
	181: "ROU",
	182: "RUS",
	183: "RWA",
	184: "REU",
	185: "BLM",
	186: "SHN",
	187: "KNA",
	188: "LCA",
	189: "MAF",
	190: "SPM",
	191: "VCT",
	192: "WSM",
	193: "SMR",
	194: "STP",
	195: "SAU",
	196: "SEN",
	197: "SRB",
	198: "SYC",
	199: "SLE",
	200: "SGP",
	201: "SXM",
	202: "SVK",
	203: "SVN",
	204: "SLB",
	205: "SOM",
	206: "ZAF",
	207: "SGS",
	208: "SSD",
	209: "ESP",
	210: "LKA",
	211: "SDN",
	212: "SUR",
	213: "SJM",
	214: "SWZ",
	215: "SWE",
	216: "CHE",
	217: "SYR",
	218: "TWN",
	219: "TJK",
	220: "TZA",
	221: "THA",
	222: "TLS",
	223: "TGO",
	224: "TKL",
	225: "TON",
	226: "TTO",
	227: "TUN",
	228: "TUR",
	229: "TKM",
	230: "TCA",
	231: "TUV",
	232: "UGA",
	233: "UKR",
	234: "ARE",
	235: "GBR",
	236: "USA",
	237: "UMI",
	238: "URY",
	239: "UZB",
	240: "VUT",
	241: "VEN",
	242: "VNM",
	243: "VGB",
	244: "VIR",
	245: "WLF",
	246: "ESH",
	247: "YEM",
	248: "ZMB",
	249: "ZWE",
}

var CountryValues = map[string]int32{
	"OTHER": 1,
	"AFG":   2,
	"ALB":   3,
	"DZA":   4,
	"ASM":   5,
	"AND":   6,
	"AGO":   7,
	"AIA":   8,
	"ATA":   9,
	"ATG":   10,
	"ARG":   11,
	"ARM":   12,
	"ABW":   13,
	"AUS":   14,
	"AUT":   15,
	"AZE":   16,
	"BHS":   17,
	"BHR":   18,
	"BGD":   19,
	"BRB":   20,
	"BLR":   21,
	"BEL":   22,
	"BLZ":   23,
	"BEN":   24,
	"BMU":   25,
	"BTN":   26,
	"BOL":   27,
	"BES":   28,
	"BIH":   29,
	"BWA":   30,
	"BVT":   31,
	"BRA":   32,
	"IOT":   33,
	"BRN":   34,
	"BGR":   35,
	"BFA":   36,
	"BDI":   37,
	"KHM":   38,
	"CMR":   39,
	"CAN":   40,
	"CPV":   41,
	"CYM":   42,
	"CAF":   43,
	"TCD":   44,
	"CHL":   45,
	"CHN":   46,
	"CXR":   47,
	"CCK":   48,
	"COL":   49,
	"COM":   50,
	"COG":   51,
	"COD":   52,
	"COK":   53,
	"CRI":   54,
	"HRV":   55,
	"CUB":   56,
	"CUW":   57,
	"CYP":   58,
	"CZE":   59,
	"CIV":   60,
	"DNK":   61,
	"DJI":   62,
	"DMA":   63,
	"DOM":   64,
	"ECU":   65,
	"EGY":   66,
	"SLV":   67,
	"GNQ":   68,
	"ERI":   69,
	"EST":   70,
	"ETH":   71,
	"FLK":   72,
	"FRO":   73,
	"FJI":   74,
	"FIN":   75,
	"FRA":   76,
	"GUF":   77,
	"PYF":   78,
	"ATF":   79,
	"GAB":   80,
	"GMB":   81,
	"GEO":   82,
	"DEU":   83,
	"GHA":   84,
	"GIB":   85,
	"GRC":   86,
	"GRL":   87,
	"GRD":   88,
	"GLP":   89,
	"GUM":   90,
	"GTM":   91,
	"GGY":   92,
	"GIN":   93,
	"GNB":   94,
	"GUY":   95,
	"HTI":   96,
	"HMD":   97,
	"VAT":   98,
	"HND":   99,
	"HKG":   100,
	"HUN":   101,
	"ISL":   102,
	"IND":   103,
	"IDN":   104,
	"IRN":   105,
	"IRQ":   106,
	"IRL":   107,
	"IMN":   108,
	"ISR":   109,
	"ITA":   110,
	"JAM":   111,
	"JPN":   112,
	"JEY":   113,
	"JOR":   114,
	"KAZ":   115,
	"KEN":   116,
	"KIR":   117,
	"PRK":   118,
	"KOR":   119,
	"KWT":   120,
	"KGZ":   121,
	"LAO":   122,
	"LVA":   123,
	"LBN":   124,
	"LSO":   125,
	"LBR":   126,
	"LBY":   127,
	"LIE":   128,
	"LTU":   129,
	"LUX":   130,
	"MAC":   131,
	"MKD":   132,
	"MDG":   133,
	"MWI":   134,
	"MYS":   135,
	"MDV":   136,
	"MLI":   137,
	"MLT":   138,
	"MHL":   139,
	"MTQ":   140,
	"MRT":   141,
	"MUS":   142,
	"MYT":   143,
	"MEX":   144,
	"FSM":   145,
	"MDA":   146,
	"MCO":   147,
	"MNG":   148,
	"MNE":   149,
	"MSR":   150,
	"MAR":   151,
	"MOZ":   152,
	"MMR":   153,
	"NAM":   154,
	"NRU":   155,
	"NPL":   156,
	"NLD":   157,
	"NCL":   158,
	"NZL":   159,
	"NIC":   160,
	"NER":   161,
	"NGA":   162,
	"NIU":   163,
	"NFK":   164,
	"MNP":   165,
	"NOR":   166,
	"OMN":   167,
	"PAK":   168,
	"PLW":   169,
	"PSE":   170,
	"PAN":   171,
	"PNG":   172,
	"PRY":   173,
	"PER":   174,
	"PHL":   175,
	"PCN":   176,
	"POL":   177,
	"PRT":   178,
	"PRI":   179,
	"QAT":   180,
	"ROU":   181,
	"RUS":   182,
	"RWA":   183,
	"REU":   184,
	"BLM":   185,
	"SHN":   186,
	"KNA":   187,
	"LCA":   188,
	"MAF":   189,
	"SPM":   190,
	"VCT":   191,
	"WSM":   192,
	"SMR":   193,
	"STP":   194,
	"SAU":   195,
	"SEN":   196,
	"SRB":   197,
	"SYC":   198,
	"SLE":   199,
	"SGP":   200,
	"SXM":   201,
	"SVK":   202,
	"SVN":   203,
	"SLB":   204,
	"SOM":   205,
	"ZAF":   206,
	"SGS":   207,
	"SSD":   208,
	"ESP":   209,
	"LKA":   210,
	"SDN":   211,
	"SUR":   212,
	"SJM":   213,
	"SWZ":   214,
	"SWE":   215,
	"CHE":   216,
	"SYR":   217,
	"TWN":   218,
	"TJK":   219,
	"TZA":   220,
	"THA":   221,
	"TLS":   222,
	"TGO":   223,
	"TKL":   224,
	"TON":   225,
	"TTO":   226,
	"TUN":   227,
	"TUR":   228,
	"TKM":   229,
	"TCA":   230,
	"TUV":   231,
	"UGA":   232,
	"UKR":   233,
	"ARE":   234,
	"GBR":   235,
	"USA":   236,
	"UMI":   237,
	"URY":   238,
	"UZB":   239,
	"VUT":   240,
	"VEN":   241,
	"VNM":   242,
	"VGB":   243,
	"VIR":   244,
	"WLF":   245,
	"ESH":   246,
	"YEM":   247,
	"ZMB":   248,
	"ZWE":   249,
}

// CountryAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var CountryAliases = []state.Alias{
	{Alias: "^(other|unknown|undefined)$", Value: 1},
	{Alias: "^(afghanistan|af|afg)$", Value: 2},
	{Alias: "^(albania|al|alb)$", Value: 3},
	{Alias: "^(algeria|dz|dza)$", Value: 4},
	{Alias: "^(americansamoa|as|asm)$", Value: 5},
	{Alias: "^(andorra|ad|and)$", Value: 6},
	{Alias: "^(angola|ao|ago)$", Value: 7},
	{Alias: "^(anguilla|ai|aia)$", Value: 8},
	{Alias: "^(antarctica|aq|ata)$", Value: 9},
	{Alias: "^(antiguaandbarbuda|ag|atg)$", Value: 10},
	{Alias: "^(argentina|ar|arg)$", Value: 11},
	{Alias: "^(armenia|am|arm)$", Value: 12},
	{Alias: "^(aruba|aw|abw)$", Value: 13},
	{Alias: "^(australia|au|aus)$", Value: 14},
	{Alias: "^(austria|at|aut)$", Value: 15},
	{Alias: "^(azerbaijan|az|aze)$", Value: 16},
	{Alias: "^(bahamas|bs|bhs)$", Value: 17},
	{Alias: "^(bahrain|bh|bhr)$", Value: 18},
	{Alias: "^(bangladesh|bd|bgd)$", Value: 19},
	{Alias: "^(barbados|bb|brb)$", Value: 20},
	{Alias: "^(belarus|by|blr)$", Value: 21},
	{Alias: "^(belgium|be|bel)$", Value: 22},
	{Alias: "^(belize|bz|blz)$", Value: 23},
	{Alias: "^(benin|bj|ben)$", Value: 24},
	{Alias: "^(bermuda|bm|bmu)$", Value: 25},
	{Alias: "^(bhutan|bt|btn)$", Value: 26},
	{Alias: "^(bolivia|bo|bol)$", Value: 27},
	{Alias: "^(bonaire|bq|bes)$", Value: 28},
	{Alias: "^(bosniaandherzegovina|ba|bih)$", Value: 29},
	{Alias: "^(botswana|bw|bwa)$", Value: 30},
	{Alias: "^(bouvetisland|bv|bvt)$", Value: 31},
	{Alias: "^(brazil|br|bra)$", Value: 32},
	{Alias: "^(britishindianoceanterritory|io|iot)$", Value: 33},
	{Alias: "^(bruneidarussalam|bn|brn)$", Value: 34},
	{Alias: "^(bulgaria|bg|bgr)$", Value: 35},
	{Alias: "^(burkinafaso|bf|bfa)$", Value: 36},
	{Alias: "^(burundi|bi|bdi)$", Value: 37},
	{Alias: "^(cambodia|kh|khm)$", Value: 38},
	{Alias: "^(cameroon|cm|cmr)$", Value: 39},
	{Alias: "^(canada|ca|can)$", Value: 40},
	{Alias: "^(capeverde|cv|cpv)$", Value: 41},
	{Alias: "^(caymanislands|ky|cym)$", Value: 42},
	{Alias: "^(centralafricanrepublic|cf|caf)$", Value: 43},
	{Alias: "^(chad|td|tcd)$", Value: 44},
	{Alias: "^(chile|cl|chl)$", Value: 45},
	{Alias: "^(china|cn|chn)$", Value: 46},
	{Alias: "^(christmasisland|cx|cxr)$", Value: 47},
	{Alias: "^(cocoskeelingislands|cc|cck)$", Value: 48},
	{Alias: "^(colombia|co|col)$", Value: 49},
	{Alias: "^(comoros|km|com)$", Value: 50},
	{Alias: "^(congo|cg|cog)$", Value: 51},
	{Alias: "^(democraticrepublicofthecongo|cd|cod)$", Value: 52},
	{Alias: "^(cookislands|ck|cok)$", Value: 53},
	{Alias: "^(costarica|cr|cri)$", Value: 54},
	{Alias: "^(croatia|hr|hrv)$", Value: 55},
	{Alias: "^(cuba|cu|cub)$", Value: 56},
	{Alias: "^(curacao|cw|cuw)$", Value: 57},
	{Alias: "^(cyprus|cy|cyp)$", Value: 58},
	{Alias: "^(czechrepublic|cz|cze)$", Value: 59},
	{Alias: "^(cotedivoire|ci|civ)$", Value: 60},
	{Alias: "^(denmark|dk|dnk)$", Value: 61},
	{Alias: "^(djibouti|dj|dji)$", Value: 62},
	{Alias: "^(dominica|dm|dma)$", Value: 63},
	{Alias: "^(dominicanrepublic|do|dom)$", Value: 64},
	{Alias: "^(ecuador|ec|ecu)$", Value: 65},
	{Alias: "^(egypt|eg|egy)$", Value: 66},
	{Alias: "^(elsalvador|sv|slv)$", Value: 67},
	{Alias: "^(equatorialguinea|gq|gnq)$", Value: 68},
	{Alias: "^(eritrea|er|eri)$", Value: 69},
	{Alias: "^(estonia|ee|est)$", Value: 70},
	{Alias: "^(ethiopia|et|eth)$", Value: 71},
	{Alias: "^(falklandislands|fk|flk)$", Value: 72},
	{Alias: "^(faroeislands|fo|fro)$", Value: 73},
	{Alias: "^(fiji|fj|fji)$", Value: 74},
	{Alias: "^(finland|fi|fin)$", Value: 75},
	{Alias: "^(france|fr|fra)$", Value: 76},
	{Alias: "^(frenchguiana|gf|guf)$", Value: 77},
	{Alias: "^(frenchpolynesia|pf|pyf)$", Value: 78},
	{Alias: "^(frenchsouthernterritories|tf|atf)$", Value: 79},
	{Alias: "^(gabon|ga|gab)$", Value: 80},
	{Alias: "^(gambia|gm|gmb)$", Value: 81},
	{Alias: "^(georgia|ge|geo)$", Value: 82},
	{Alias: "^(germany|de|deu)$", Value: 83},
	{Alias: "^(ghana|gh|gha)$", Value: 84},
	{Alias: "^(gibraltar|gi|gib)$", Value: 85},
	{Alias: "^(greece|gr|grc)$", Value: 86},
	{Alias: "^(greenland|gl|grl)$", Value: 87},
	{Alias: "^(grenada|gd|grd)$", Value: 88},
	{Alias: "^(guadeloupe|gp|glp)$", Value: 89},
	{Alias: "^(guam|gu|gum)$", Value: 90},
	{Alias: "^(guatemala|gt|gtm)$", Value: 91},
	{Alias: "^(guernsey|gg|ggy)$", Value: 92},
	{Alias: "^(guinea|gn|gin)$", Value: 93},
	{Alias: "^(guineabissau|gw|gnb)$", Value: 94},
	{Alias: "^(guyana|gy|guy)$", Value: 95},
	{Alias: "^(haiti|ht|hti)$", Value: 96},
	{Alias: "^(heardislandandmcdonaldislands|hm|hmd)$", Value: 97},
	{Alias: "^(holysee|va|vat)$", Value: 98},
	{Alias: "^(honduras|hn|hnd)$", Value: 99},
	{Alias: "^(hongkong|hk|hkg)$", Value: 100},
	{Alias: "^(hungary|hu|hun)$", Value: 101},
	{Alias: "^(iceland|is|isl)$", Value: 102},
	{Alias: "^(india|in|ind)$", Value: 103},
	{Alias: "^(indonesia|id|idn)$", Value: 104},
	{Alias: "^(iran|ir|irn)$", Value: 105},
	{Alias: "^(iraq|iq|irq)$", Value: 106},
	{Alias: "^(ireland|ie|irl)$", Value: 107},
	{Alias: "^(isleofman|im|imn)$", Value: 108},
	{Alias: "^(israel|il|isr)$", Value: 109},
	{Alias: "^(italy|it|ita)$", Value: 110},
	{Alias: "^(jamaica|jm|jam)$", Value: 111},
	{Alias: "^(japan|jp|jpn)$", Value: 112},
	{Alias: "^(jersey|je|jey)$", Value: 113},
	{Alias: "^(jordan|jo|jor)$", Value: 114},
	{Alias: "^(kazakhstan|kz|kaz)$", Value: 115},
	{Alias: "^(kenya|ke|ken)$", Value: 116},
	{Alias: "^(kiribati|ki|kir)$", Value: 117},
	{Alias: "^(northkorea|kp|prk)$", Value: 118},
	{Alias: "^(southkorea|kr|kor)$", Value: 119},
	{Alias: "^(kuwait|kw|kwt)$", Value: 120},
	{Alias: "^(kyrgyzstan|kg|kgz)$", Value: 121},
	{Alias: "^(laopeoplesdemocraticrepublic|la|lao)$", Value: 122},
	{Alias: "^(latvia|lv|lva)$", Value: 123},
	{Alias: "^(lebanon|lb|lbn)$", Value: 124},
	{Alias: "^(lesotho|ls|lso)$", Value: 125},
	{Alias: "^(liberia|lr|lbr)$", Value: 126},
	{Alias: "^(libya|ly|lby)$", Value: 127},
	{Alias: "^(liechtenstein|li|lie)$", Value: 128},
	{Alias: "^(lithuania|lt|ltu)$", Value: 129},
	{Alias: "^(luxembourg|lu|lux)$", Value: 130},
	{Alias: "^(macao|mo|mac)$", Value: 131},
	{Alias: "^(northmacedonia|mk|mkd)$", Value: 132},
	{Alias: "^(madagascar|mg|mdg)$", Value: 133},
	{Alias: "^(malawi|mw|mwi)$", Value: 134},
	{Alias: "^(malaysia|my|mys)$", Value: 135},
	{Alias: "^(maldives|mv|mdv)$", Value: 136},
	{Alias: "^(mali|ml|mli)$", Value: 137},
	{Alias: "^(malta|mt|mlt)$", Value: 138},
	{Alias: "^(marshallislands|mh|mhl)$", Value: 139},
	{Alias: "^(martinique|mq|mtq)$", Value: 140},
	{Alias: "^(mauritania|mr|mrt)$", Value: 141},
	{Alias: "^(mauritius|mu|mus)$", Value: 142},
	{Alias: "^(mayotte|yt|myt)$", Value: 143},
	{Alias: "^(mexico|mx|mex)$", Value: 144},
	{Alias: "^(micronesia|fm|fsm)$", Value: 145},
	{Alias: "^(moldova|md|mda)$", Value: 146},
	{Alias: "^(monaco|mc|mco)$", Value: 147},
	{Alias: "^(mongolia|mn|mng)$", Value: 148},
	{Alias: "^(montenegro|me|mne)$", Value: 149},
	{Alias: "^(montserrat|ms|msr)$", Value: 150},
	{Alias: "^(morocco|ma|mar)$", Value: 151},
	{Alias: "^(mozambique|mz|moz)$", Value: 152},
	{Alias: "^(myanmar|mm|mmr)$", Value: 153},
	{Alias: "^(namibia|na|nam)$", Value: 154},
	{Alias: "^(nauru|nr|nru)$", Value: 155},
	{Alias: "^(nepal|np|npl)$", Value: 156},
	{Alias: "^(netherlands|nl|nld)$", Value: 157},
	{Alias: "^(newcaledonia|nc|ncl)$", Value: 158},
	{Alias: "^(newzealand|nz|nzl)$", Value: 159},
	{Alias: "^(nicaragua|ni|nic)$", Value: 160},
	{Alias: "^(niger|ne|ner)$", Value: 161},
	{Alias: "^(nigeria|ng|nga)$", Value: 162},
	{Alias: "^(niue|nu|niu)$", Value: 163},
	{Alias: "^(norfolkisland|nf|nfk)$", Value: 164},
	{Alias: "^(northernmarianaislands|mp|mnp)$", Value: 165},
	{Alias: "^(norway|no|nor)$", Value: 166},
	{Alias: "^(oman|om|omn)$", Value: 167},
	{Alias: "^(pakistan|pk|pak)$", Value: 168},
	{Alias: "^(palau|pw|plw)$", Value: 169},
	{Alias: "^(palestine|ps|pse)$", Value: 170},
	{Alias: "^(panama|pa|pan)$", Value: 171},
	{Alias: "^(papuanewguinea|pg|png)$", Value: 172},
	{Alias: "^(paraguay|py|pry)$", Value: 173},
	{Alias: "^(peru|pe|per)$", Value: 174},
	{Alias: "^(philippines|ph|phl)$", Value: 175},
	{Alias: "^(pitcairn|pn|pcn)$", Value: 176},
	{Alias: "^(poland|pl|pol)$", Value: 177},
	{Alias: "^(portugal|pt|prt)$", Value: 178},
	{Alias: "^(puertorico|pr|pri)$", Value: 179},
	{Alias: "^(qatar|qa|qat)$", Value: 180},
	{Alias: "^(romania|ro|rou)$", Value: 181},
	{Alias: "^(russianfederation|ru|rus)$", Value: 182},
	{Alias: "^(rwanda|rw|rwa)$", Value: 183},
	{Alias: "^(reunion|re|reu)$", Value: 184},
	{Alias: "^(saintbarthelemy|bl|blm)$", Value: 185},
	{Alias: "^(sainthelena|sh|shn)$", Value: 186},
	{Alias: "^(saintkittsandnevis|kn|kna)$", Value: 187},
	{Alias: "^(saintlucia|lc|lca)$", Value: 188},
	{Alias: "^(saintmartinfrenchpart|mf|maf)$", Value: 189},
	{Alias: "^(saintpierreandmiquelon|pm|spm)$", Value: 190},
	{Alias: "^(saintvincentandthegrenadines|vc|vct)$", Value: 191},
	{Alias: "^(samoa|ws|wsm)$", Value: 192},
	{Alias: "^(sanmarino|sm|smr)$", Value: 193},
	{Alias: "^(saotomeandprincipe|st|stp)$", Value: 194},
	{Alias: "^(saudiarabia|sa|sau)$", Value: 195},
	{Alias: "^(senegal|sn|sen)$", Value: 196},
	{Alias: "^(serbia|rs|srb)$", Value: 197},
	{Alias: "^(seychelles|sc|syc)$", Value: 198},
	{Alias: "^(sierraleone|sl|sle)$", Value: 199},
	{Alias: "^(singapore|sg|sgp)$", Value: 200},
	{Alias: "^(sintmaartendutchpart|sx|sxm)$", Value: 201},
	{Alias: "^(slovakia|sk|svk)$", Value: 202},
	{Alias: "^(slovenia|si|svn)$", Value: 203},
	{Alias: "^(solomonislands|sb|slb)$", Value: 204},
	{Alias: "^(somalia|so|som)$", Value: 205},
	{Alias: "^(southafrica|za|zaf)$", Value: 206},
	{Alias: "^(southgeorgiaandthesouthsandwichislands|gs|sgs)$", Value: 207},
	{Alias: "^(southsudan|ss|ssd)$", Value: 208},
	{Alias: "^(spain|es|esp)$", Value: 209},
	{Alias: "^(srilanka|lk|lka)$", Value: 210},
	{Alias: "^(sudan|sd|sdn)$", Value: 211},
	{Alias: "^(suriname|sr|sur)$", Value: 212},
	{Alias: "^(svalbardandjanmayen|sj|sjm)$", Value: 213},
	{Alias: "^(swaziland|sz|swz)$", Value: 214},
	{Alias: "^(sweden|se|swe)$", Value: 215},
	{Alias: "^(switzerland|ch|che)$", Value: 216},
	{Alias: "^(syrianarabrepublic|sy|syr)$", Value: 217},
	{Alias: "^(taiwan|tw|twn)$", Value: 218},
	{Alias: "^(tajikistan|tj|tjk)$", Value: 219},
	{Alias: "^(unitedrepublicoftanzania|tz|tza)$", Value: 220},
	{Alias: "^(thailand|th|tha)$", Value: 221},
	{Alias: "^(timorleste|tl|tls)$", Value: 222},
	{Alias: "^(togo|tg|tgo)$", Value: 223},
	{Alias: "^(tokelau|tk|tkl)$", Value: 224},
	{Alias: "^(tonga|to|ton)$", Value: 225},
	{Alias: "^(trinidadandtobago|tt|tto)$", Value: 226},
	{Alias: "^(tunisia|tn|tun)$", Value: 227},
	{Alias: "^(turkey|tr|tur)$", Value: 228},
	{Alias: "^(turkmenistan|tm|tkm)$", Value: 229},
	{Alias: "^(turksandcaicosislands|tc|tca)$", Value: 230},
	{Alias: "^(tuvalu|tv|tuv)$", Value: 231},
	{Alias: "^(uganda|ug|uga)$", Value: 232},
	{Alias: "^(ukraine|ua|ukr|україна)$", Value: 233},
	{Alias: "^(unitedarabemirates|ae|are|dubai)$", Value: 234},
	{Alias: "^(unitedkingdom|gb|gbr)$", Value: 235},
	{Alias: "^(unitedstates|us|usa)$", Value: 236},
	{Alias: "^(unitedstatesminoroutlyingislands|um|umi)$", Value: 237},
	{Alias: "^(uruguay|uy|ury)$", Value: 238},
	{Alias: "^(uzbekistan|uz|uzb)$", Value: 239},
	{Alias: "^(vanuatu|vu|vut)$", Value: 240},
	{Alias: "^(venezuela|ve|ven)$", Value: 241},
	{Alias: "^(vietnam|vn|vnm)$", Value: 242},
	{Alias: "^(britishvirginislands|vg|vgb)$", Value: 243},
	{Alias: "^(usvirginislands|vi|vir)$", Value: 244},
	{Alias: "^(wallisandfutuna|wf|wlf)$", Value: 245},
	{Alias: "^(westernsahara|eh|esh)$", Value: 246},
	{Alias: "^(yemen|ye|yem)$", Value: 247},
	{Alias: "^(zambia|zm|zmb)$", Value: 248},
	{Alias: "^(zimbabwe|zw|zwe)$", Value: 249},
}

// CountryOrderedValues is a list of sorted values
// of Country. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var CountryOrderedValues = []int32{
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
	102,
	103,
	104,
	105,
	106,
	107,
	108,
	109,
	110,
	111,
	112,
	113,
	114,
	115,
	116,
	117,
	118,
	119,
	120,
	121,
	122,
	123,
	124,
	125,
	126,
	127,
	128,
	129,
	130,
	131,
	132,
	133,
	134,
	135,
	136,
	137,
	138,
	139,
	140,
	141,
	142,
	143,
	144,
	145,
	146,
	147,
	148,
	149,
	150,
	151,
	152,
	153,
	154,
	155,
	156,
	157,
	158,
	159,
	160,
	161,
	162,
	163,
	164,
	165,
	166,
	167,
	168,
	169,
	170,
	171,
	172,
	173,
	174,
	175,
	176,
	177,
	178,
	179,
	180,
	181,
	182,
	183,
	184,
	185,
	186,
	187,
	188,
	189,
	190,
	191,
	192,
	193,
	194,
	195,
	196,
	197,
	198,
	199,
	200,
	201,
	202,
	203,
	204,
	205,
	206,
	207,
	208,
	209,
	210,
	211,
	212,
	213,
	214,
	215,
	216,
	217,
	218,
	219,
	220,
	221,
	222,
	223,
	224,
	225,
	226,
	227,
	228,
	229,
	230,
	231,
	232,
	233,
	234,
	235,
	236,
	237,
	238,
	239,
	240,
	241,
	242,
	243,
	244,
	245,
	246,
	247,
	248,
	249,
}
