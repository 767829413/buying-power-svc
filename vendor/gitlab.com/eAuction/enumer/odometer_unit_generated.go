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

// OdometerUnit is a public version of "OdometerUnit" type of enum value.
type OdometerUnit int32

// IsKnownOdometerUnit returns true if provided type is a valid
// value of OdometerUnit enum.
func IsKnownOdometerUnit(v int32) bool {
	_, ok := getState().OdometerUnitNames[v]
	return ok
}

// OdometerUnitFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - OdometerUnitFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - OdometerUnitFromString panics.
func OdometerUnitFromString(s string) OdometerUnit {
	res, err := OdometerUnitFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var odometerUnitRegexCache = map[string]*regexp.Regexp{}
var odometerUnitRegexMutex = &sync.RWMutex{}

// OdometerUnitFromStringE does the same as OdometerUnitFromString does, but instead of
// panic, returns an error.
func OdometerUnitFromStringE(s string) (OdometerUnit, error) {
	state := getState()
	if v, ok := state.OdometerUnitValues[s]; ok {
		return OdometerUnit(v), nil
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

	for _, aliasStruct := range state.OdometerUnitAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		odometerUnitRegexMutex.RLock()
		re, ok := odometerUnitRegexCache[alias]
		odometerUnitRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			odometerUnitRegexMutex.Lock()
			odometerUnitRegexCache[alias] = re
			odometerUnitRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return OdometerUnit(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "OdometerUnit ", Value: s}
}

// UnmarshalText - populates OdometerUnit using provided text
func (v *OdometerUnit) UnmarshalText(text []byte) error {
	value, ok := getState().OdometerUnitValues[string(text)]
	if !ok {
		return errors.New("cannot find value for OdometerUnit")
	}

	*v = OdometerUnit(value)
	return nil
}

// MarshalText - converts OdometerUnit into text
func (v OdometerUnit) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of OdometerUnit
func (v OdometerUnit) String() string {
	str, ok := getState().OdometerUnitNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for OdometerUnit"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v OdometerUnit) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["ODOMETER_UNIT"][v.String()]
}

// MarshalJSON encodes OdometerUnit value into json value.
func (v OdometerUnit) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into OdometerUnit.
func (v *OdometerUnit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
