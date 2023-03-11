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

// Region is a public version of "Region" type of enum value.
type Region int32

// IsKnownRegion returns true if provided type is a valid
// value of Region enum.
func IsKnownRegion(v int32) bool {
	_, ok := getState().RegionNames[v]
	return ok
}

// RegionFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - RegionFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - RegionFromString panics.
func RegionFromString(s string) Region {
	res, err := RegionFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var regionRegexCache = map[string]*regexp.Regexp{}
var regionRegexMutex = &sync.RWMutex{}

// RegionFromStringE does the same as RegionFromString does, but instead of
// panic, returns an error.
func RegionFromStringE(s string) (Region, error) {
	state := getState()
	if v, ok := state.RegionValues[s]; ok {
		return Region(v), nil
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

	for _, aliasStruct := range state.RegionAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		regionRegexMutex.RLock()
		re, ok := regionRegexCache[alias]
		regionRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			regionRegexMutex.Lock()
			regionRegexCache[alias] = re
			regionRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return Region(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "Region ", Value: s}
}

// UnmarshalText - populates Region using provided text
func (v *Region) UnmarshalText(text []byte) error {
	value, ok := getState().RegionValues[string(text)]
	if !ok {
		return errors.New("cannot find value for Region")
	}

	*v = Region(value)
	return nil
}

// MarshalText - converts Region into text
func (v Region) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of Region
func (v Region) String() string {
	str, ok := getState().RegionNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for Region"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v Region) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["REGION"][v.String()]
}

// MarshalJSON encodes Region value into json value.
func (v Region) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into Region.
func (v *Region) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
