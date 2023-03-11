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

// FuelType is a public version of "FuelType" type of enum value.
type FuelType int32

// IsKnownFuelType returns true if provided type is a valid
// value of FuelType enum.
func IsKnownFuelType(v int32) bool {
	_, ok := getState().FuelTypeNames[v]
	return ok
}

// FuelTypeFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - FuelTypeFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - FuelTypeFromString panics.
func FuelTypeFromString(s string) FuelType {
	res, err := FuelTypeFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var fuelTypeRegexCache = map[string]*regexp.Regexp{}
var fuelTypeRegexMutex = &sync.RWMutex{}

// FuelTypeFromStringE does the same as FuelTypeFromString does, but instead of
// panic, returns an error.
func FuelTypeFromStringE(s string) (FuelType, error) {
	state := getState()
	if v, ok := state.FuelTypeValues[s]; ok {
		return FuelType(v), nil
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

	for _, aliasStruct := range state.FuelTypeAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		fuelTypeRegexMutex.RLock()
		re, ok := fuelTypeRegexCache[alias]
		fuelTypeRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			fuelTypeRegexMutex.Lock()
			fuelTypeRegexCache[alias] = re
			fuelTypeRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return FuelType(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "FuelType ", Value: s}
}

// UnmarshalText - populates FuelType using provided text
func (v *FuelType) UnmarshalText(text []byte) error {
	value, ok := getState().FuelTypeValues[string(text)]
	if !ok {
		return errors.New("cannot find value for FuelType")
	}

	*v = FuelType(value)
	return nil
}

// MarshalText - converts FuelType into text
func (v FuelType) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of FuelType
func (v FuelType) String() string {
	str, ok := getState().FuelTypeNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for FuelType"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v FuelType) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["FUEL_TYPE"][v.String()]
}

// MarshalJSON encodes FuelType value into json value.
func (v FuelType) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into FuelType.
func (v *FuelType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
