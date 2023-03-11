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

// WheelPosition is a public version of "WheelPosition" type of enum value.
type WheelPosition int32

// IsKnownWheelPosition returns true if provided type is a valid
// value of WheelPosition enum.
func IsKnownWheelPosition(v int32) bool {
	_, ok := getState().WheelPositionNames[v]
	return ok
}

// WheelPositionFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - WheelPositionFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - WheelPositionFromString panics.
func WheelPositionFromString(s string) WheelPosition {
	res, err := WheelPositionFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var wheelPositionRegexCache = map[string]*regexp.Regexp{}
var wheelPositionRegexMutex = &sync.RWMutex{}

// WheelPositionFromStringE does the same as WheelPositionFromString does, but instead of
// panic, returns an error.
func WheelPositionFromStringE(s string) (WheelPosition, error) {
	state := getState()
	if v, ok := state.WheelPositionValues[s]; ok {
		return WheelPosition(v), nil
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

	for _, aliasStruct := range state.WheelPositionAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		wheelPositionRegexMutex.RLock()
		re, ok := wheelPositionRegexCache[alias]
		wheelPositionRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			wheelPositionRegexMutex.Lock()
			wheelPositionRegexCache[alias] = re
			wheelPositionRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return WheelPosition(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "WheelPosition ", Value: s}
}

// UnmarshalText - populates WheelPosition using provided text
func (v *WheelPosition) UnmarshalText(text []byte) error {
	value, ok := getState().WheelPositionValues[string(text)]
	if !ok {
		return errors.New("cannot find value for WheelPosition")
	}

	*v = WheelPosition(value)
	return nil
}

// MarshalText - converts WheelPosition into text
func (v WheelPosition) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of WheelPosition
func (v WheelPosition) String() string {
	str, ok := getState().WheelPositionNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for WheelPosition"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v WheelPosition) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["WHEEL_POSITION"][v.String()]
}

// MarshalJSON encodes WheelPosition value into json value.
func (v WheelPosition) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into WheelPosition.
func (v *WheelPosition) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
