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

// DamageType is a public version of "DamageType" type of enum value.
type DamageType int32

// IsKnownDamageType returns true if provided type is a valid
// value of DamageType enum.
func IsKnownDamageType(v int32) bool {
	_, ok := getState().DamageTypeNames[v]
	return ok
}

// DamageTypeFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - DamageTypeFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - DamageTypeFromString panics.
func DamageTypeFromString(s string) DamageType {
	res, err := DamageTypeFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var damageTypeRegexCache = map[string]*regexp.Regexp{}
var damageTypeRegexMutex = &sync.RWMutex{}

// DamageTypeFromStringE does the same as DamageTypeFromString does, but instead of
// panic, returns an error.
func DamageTypeFromStringE(s string) (DamageType, error) {
	state := getState()
	if v, ok := state.DamageTypeValues[s]; ok {
		return DamageType(v), nil
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

	for _, aliasStruct := range state.DamageTypeAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		damageTypeRegexMutex.RLock()
		re, ok := damageTypeRegexCache[alias]
		damageTypeRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			damageTypeRegexMutex.Lock()
			damageTypeRegexCache[alias] = re
			damageTypeRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return DamageType(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "DamageType ", Value: s}
}

// UnmarshalText - populates DamageType using provided text
func (v *DamageType) UnmarshalText(text []byte) error {
	value, ok := getState().DamageTypeValues[string(text)]
	if !ok {
		return errors.New("cannot find value for DamageType")
	}

	*v = DamageType(value)
	return nil
}

// MarshalText - converts DamageType into text
func (v DamageType) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of DamageType
func (v DamageType) String() string {
	str, ok := getState().DamageTypeNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for DamageType"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v DamageType) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["DAMAGE_TYPE"][v.String()]
}

// MarshalJSON encodes DamageType value into json value.
func (v DamageType) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into DamageType.
func (v *DamageType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
