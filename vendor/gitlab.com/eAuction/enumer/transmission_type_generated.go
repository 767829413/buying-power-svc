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

// TransmissionType is a public version of "TransmissionType" type of enum value.
type TransmissionType int32

// IsKnownTransmissionType returns true if provided type is a valid
// value of TransmissionType enum.
func IsKnownTransmissionType(v int32) bool {
	_, ok := getState().TransmissionTypeNames[v]
	return ok
}

// TransmissionTypeFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - TransmissionTypeFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - TransmissionTypeFromString panics.
func TransmissionTypeFromString(s string) TransmissionType {
	res, err := TransmissionTypeFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var transmissionTypeRegexCache = map[string]*regexp.Regexp{}
var transmissionTypeRegexMutex = &sync.RWMutex{}

// TransmissionTypeFromStringE does the same as TransmissionTypeFromString does, but instead of
// panic, returns an error.
func TransmissionTypeFromStringE(s string) (TransmissionType, error) {
	state := getState()
	if v, ok := state.TransmissionTypeValues[s]; ok {
		return TransmissionType(v), nil
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

	for _, aliasStruct := range state.TransmissionTypeAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		transmissionTypeRegexMutex.RLock()
		re, ok := transmissionTypeRegexCache[alias]
		transmissionTypeRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			transmissionTypeRegexMutex.Lock()
			transmissionTypeRegexCache[alias] = re
			transmissionTypeRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return TransmissionType(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "TransmissionType ", Value: s}
}

// UnmarshalText - populates TransmissionType using provided text
func (v *TransmissionType) UnmarshalText(text []byte) error {
	value, ok := getState().TransmissionTypeValues[string(text)]
	if !ok {
		return errors.New("cannot find value for TransmissionType")
	}

	*v = TransmissionType(value)
	return nil
}

// MarshalText - converts TransmissionType into text
func (v TransmissionType) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of TransmissionType
func (v TransmissionType) String() string {
	str, ok := getState().TransmissionTypeNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for TransmissionType"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v TransmissionType) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["TRANSMISSION_TYPE"][v.String()]
}

// MarshalJSON encodes TransmissionType value into json value.
func (v TransmissionType) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into TransmissionType.
func (v *TransmissionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
