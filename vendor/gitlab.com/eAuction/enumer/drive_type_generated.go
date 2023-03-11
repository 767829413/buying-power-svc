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

// DriveType is a public version of "DriveType" type of enum value.
type DriveType int32

// IsKnownDriveType returns true if provided type is a valid
// value of DriveType enum.
func IsKnownDriveType(v int32) bool {
	_, ok := getState().DriveTypeNames[v]
	return ok
}

// DriveTypeFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - DriveTypeFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - DriveTypeFromString panics.
func DriveTypeFromString(s string) DriveType {
	res, err := DriveTypeFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var driveTypeRegexCache = map[string]*regexp.Regexp{}
var driveTypeRegexMutex = &sync.RWMutex{}

// DriveTypeFromStringE does the same as DriveTypeFromString does, but instead of
// panic, returns an error.
func DriveTypeFromStringE(s string) (DriveType, error) {
	state := getState()
	if v, ok := state.DriveTypeValues[s]; ok {
		return DriveType(v), nil
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

	for _, aliasStruct := range state.DriveTypeAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		driveTypeRegexMutex.RLock()
		re, ok := driveTypeRegexCache[alias]
		driveTypeRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			driveTypeRegexMutex.Lock()
			driveTypeRegexCache[alias] = re
			driveTypeRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return DriveType(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "DriveType ", Value: s}
}

// UnmarshalText - populates DriveType using provided text
func (v *DriveType) UnmarshalText(text []byte) error {
	value, ok := getState().DriveTypeValues[string(text)]
	if !ok {
		return errors.New("cannot find value for DriveType")
	}

	*v = DriveType(value)
	return nil
}

// MarshalText - converts DriveType into text
func (v DriveType) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of DriveType
func (v DriveType) String() string {
	str, ok := getState().DriveTypeNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for DriveType"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v DriveType) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["DRIVE_TYPE"][v.String()]
}

// MarshalJSON encodes DriveType value into json value.
func (v DriveType) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into DriveType.
func (v *DriveType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
