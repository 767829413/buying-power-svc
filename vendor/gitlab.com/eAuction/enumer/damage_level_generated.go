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

// DamageLevel is a public version of "DamageLevel" type of enum value.
type DamageLevel int32

// IsKnownDamageLevel returns true if provided type is a valid
// value of DamageLevel enum.
func IsKnownDamageLevel(v int32) bool {
	_, ok := getState().DamageLevelNames[v]
	return ok
}

// DamageLevelFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - DamageLevelFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - DamageLevelFromString panics.
func DamageLevelFromString(s string) DamageLevel {
	res, err := DamageLevelFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var damageLevelRegexCache = map[string]*regexp.Regexp{}
var damageLevelRegexMutex = &sync.RWMutex{}

// DamageLevelFromStringE does the same as DamageLevelFromString does, but instead of
// panic, returns an error.
func DamageLevelFromStringE(s string) (DamageLevel, error) {
	state := getState()
	if v, ok := state.DamageLevelValues[s]; ok {
		return DamageLevel(v), nil
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

	for _, aliasStruct := range state.DamageLevelAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		damageLevelRegexMutex.RLock()
		re, ok := damageLevelRegexCache[alias]
		damageLevelRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			damageLevelRegexMutex.Lock()
			damageLevelRegexCache[alias] = re
			damageLevelRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return DamageLevel(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "DamageLevel ", Value: s}
}

// UnmarshalText - populates DamageLevel using provided text
func (v *DamageLevel) UnmarshalText(text []byte) error {
	value, ok := getState().DamageLevelValues[string(text)]
	if !ok {
		return errors.New("cannot find value for DamageLevel")
	}

	*v = DamageLevel(value)
	return nil
}

// MarshalText - converts DamageLevel into text
func (v DamageLevel) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of DamageLevel
func (v DamageLevel) String() string {
	str, ok := getState().DamageLevelNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for DamageLevel"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v DamageLevel) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["DAMAGE_LEVEL"][v.String()]
}

// MarshalJSON encodes DamageLevel value into json value.
func (v DamageLevel) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into DamageLevel.
func (v *DamageLevel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
