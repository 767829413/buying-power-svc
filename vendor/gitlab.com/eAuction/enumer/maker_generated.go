// GENERATED. DO NOT EDIT.

package enumer

import (
	"encoding/json"
	"regexp"
	"sync"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/eAuction/enumer/pkg/normalization"
)

// Maker is a public version of "Maker" type of enum value.
type Maker int32

// IsKnownMaker returns true if provided type is a valid
// value of Maker enum.
func IsKnownMaker(v int32) bool {
	_, ok := getState().MakerNames[v]
	return ok
}

// MakerFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - MakerFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - MakerFromString panics.
func MakerFromString(s string) Maker {
	res, err := MakerFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var makerRegexCache = map[string]*regexp.Regexp{}
var makerRegexMutex = &sync.RWMutex{}

// MakerFromStringE does the same as MakerFromString does, but instead of
// panic, returns an error.
func MakerFromStringE(s string) (Maker, error) {
	state := getState()
	if v, ok := state.MakerValues[s]; ok {
		return Maker(v), nil
	}

	s, err := normalization.NormalizeE(s)

	if err != nil {
		return -1, errors.Wrap(err, "failed to normalize string", logan.F{
			"string": s,
		})
	}

	if s == "" {
		s = "other"
	}

	for _, aliasStruct := range state.MakerAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		makerRegexMutex.RLock()
		re, ok := makerRegexCache[alias]
		makerRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			makerRegexMutex.Lock()
			makerRegexCache[alias] = re
			makerRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return Maker(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "Maker ", Value: s}
}

// UnmarshalText - populates Maker using provided text
func (v *Maker) UnmarshalText(text []byte) error {
	value, ok := getState().MakerValues[string(text)]
	if !ok {
		return errors.New("cannot find value for Maker")
	}

	*v = Maker(value)
	return nil
}

// MarshalText - converts Maker into text
func (v Maker) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of Maker
func (v Maker) String() string {
	str, ok := getState().MakerNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for Maker"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v Maker) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["MAKER"][v.String()]
}

// MarshalJSON encodes Maker value into json value.
func (v Maker) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into Maker.
func (v *Maker) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
