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

// Condition is a public version of "Condition" type of enum value.
type Condition int32

// IsKnownCondition returns true if provided type is a valid
// value of Condition enum.
func IsKnownCondition(v int32) bool {
	_, ok := getState().ConditionNames[v]
	return ok
}

// ConditionFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - ConditionFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - ConditionFromString panics.
func ConditionFromString(s string) Condition {
	res, err := ConditionFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var conditionRegexCache = map[string]*regexp.Regexp{}
var conditionRegexMutex = &sync.RWMutex{}

// ConditionFromStringE does the same as ConditionFromString does, but instead of
// panic, returns an error.
func ConditionFromStringE(s string) (Condition, error) {
	state := getState()
	if v, ok := state.ConditionValues[s]; ok {
		return Condition(v), nil
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

	for _, aliasStruct := range state.ConditionAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		conditionRegexMutex.RLock()
		re, ok := conditionRegexCache[alias]
		conditionRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			conditionRegexMutex.Lock()
			conditionRegexCache[alias] = re
			conditionRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return Condition(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "Condition ", Value: s}
}

// UnmarshalText - populates Condition using provided text
func (v *Condition) UnmarshalText(text []byte) error {
	value, ok := getState().ConditionValues[string(text)]
	if !ok {
		return errors.New("cannot find value for Condition")
	}

	*v = Condition(value)
	return nil
}

// MarshalText - converts Condition into text
func (v Condition) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of Condition
func (v Condition) String() string {
	str, ok := getState().ConditionNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for Condition"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v Condition) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["CONDITION"][v.String()]
}

// MarshalJSON encodes Condition value into json value.
func (v Condition) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into Condition.
func (v *Condition) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
