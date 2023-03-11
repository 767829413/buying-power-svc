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

// OdometerValueState is a public version of "OdometerValueState" type of enum value.
type OdometerValueState int32

// IsKnownOdometerValueState returns true if provided type is a valid
// value of OdometerValueState enum.
func IsKnownOdometerValueState(v int32) bool {
	_, ok := getState().OdometerValueStateNames[v]
	return ok
}

// OdometerValueStateFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - OdometerValueStateFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - OdometerValueStateFromString panics.
func OdometerValueStateFromString(s string) OdometerValueState {
	res, err := OdometerValueStateFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var odometerValueStateRegexCache = map[string]*regexp.Regexp{}
var odometerValueStateRegexMutex = &sync.RWMutex{}

// OdometerValueStateFromStringE does the same as OdometerValueStateFromString does, but instead of
// panic, returns an error.
func OdometerValueStateFromStringE(s string) (OdometerValueState, error) {
	state := getState()
	if v, ok := state.OdometerValueStateValues[s]; ok {
		return OdometerValueState(v), nil
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

	for _, aliasStruct := range state.OdometerValueStateAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		odometerValueStateRegexMutex.RLock()
		re, ok := odometerValueStateRegexCache[alias]
		odometerValueStateRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			odometerValueStateRegexMutex.Lock()
			odometerValueStateRegexCache[alias] = re
			odometerValueStateRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return OdometerValueState(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "OdometerValueState ", Value: s}
}

// UnmarshalText - populates OdometerValueState using provided text
func (v *OdometerValueState) UnmarshalText(text []byte) error {
	value, ok := getState().OdometerValueStateValues[string(text)]
	if !ok {
		return errors.New("cannot find value for OdometerValueState")
	}

	*v = OdometerValueState(value)
	return nil
}

// MarshalText - converts OdometerValueState into text
func (v OdometerValueState) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of OdometerValueState
func (v OdometerValueState) String() string {
	str, ok := getState().OdometerValueStateNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for OdometerValueState"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v OdometerValueState) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["ODOMETER_VALUE_STATE"][v.String()]
}

// MarshalJSON encodes OdometerValueState value into json value.
func (v OdometerValueState) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into OdometerValueState.
func (v *OdometerValueState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
