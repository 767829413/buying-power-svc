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

// VehicleState is a public version of "VehicleState" type of enum value.
type VehicleState int32

// IsKnownVehicleState returns true if provided type is a valid
// value of VehicleState enum.
func IsKnownVehicleState(v int32) bool {
	_, ok := getState().VehicleStateNames[v]
	return ok
}

// VehicleStateFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - VehicleStateFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - VehicleStateFromString panics.
func VehicleStateFromString(s string) VehicleState {
	res, err := VehicleStateFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var vehicleStateRegexCache = map[string]*regexp.Regexp{}
var vehicleStateRegexMutex = &sync.RWMutex{}

// VehicleStateFromStringE does the same as VehicleStateFromString does, but instead of
// panic, returns an error.
func VehicleStateFromStringE(s string) (VehicleState, error) {
	state := getState()
	if v, ok := state.VehicleStateValues[s]; ok {
		return VehicleState(v), nil
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

	for _, aliasStruct := range state.VehicleStateAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		vehicleStateRegexMutex.RLock()
		re, ok := vehicleStateRegexCache[alias]
		vehicleStateRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			vehicleStateRegexMutex.Lock()
			vehicleStateRegexCache[alias] = re
			vehicleStateRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return VehicleState(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "VehicleState ", Value: s}
}

// UnmarshalText - populates VehicleState using provided text
func (v *VehicleState) UnmarshalText(text []byte) error {
	value, ok := getState().VehicleStateValues[string(text)]
	if !ok {
		return errors.New("cannot find value for VehicleState")
	}

	*v = VehicleState(value)
	return nil
}

// MarshalText - converts VehicleState into text
func (v VehicleState) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of VehicleState
func (v VehicleState) String() string {
	str, ok := getState().VehicleStateNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for VehicleState"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v VehicleState) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["VEHICLE_STATE"][v.String()]
}

// MarshalJSON encodes VehicleState value into json value.
func (v VehicleState) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into VehicleState.
func (v *VehicleState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
