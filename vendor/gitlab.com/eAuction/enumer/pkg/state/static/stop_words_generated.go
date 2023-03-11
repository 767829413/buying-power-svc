package static

import "gitlab.com/eAuction/enumer/pkg/state"

var CommonStopWords = []string{
	"(series*)",
	"coupe",
	"granturismo",
	"touring",
	"grantourer",
	"new",
}

var MakerStopWords = []state.MakerStopWordsRecord{
	state.MakerStopWordsRecord{
		Maker: 7,
		StopWords: []string{
			"xi",
			"(4drs).*",
		},
		ExcludeFromCommon: []string{
			"",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     3,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"coupe",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     15,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"new",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     42,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"coupe",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     60,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"granturismo",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     64,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"coupe",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     99,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"new",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     25,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"coupe",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     41,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"hummer",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     61,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"maybach",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     101,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"ram",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     54,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"landrover",
		},
	},
	state.MakerStopWordsRecord{
		Maker:     88,
		StopWords: []string(nil),
		ExcludeFromCommon: []string{
			"new",
		}}}
