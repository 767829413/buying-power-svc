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

// Model is a public version of "Model" type of enum value.
type Model int32

// IsKnownModel returns true if provided type is a valid
// value of Model enum.
func IsKnownModel(v int32) bool {
	_, ok := getState().ModelNames[v]
	return ok
}

// ModelFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - ModelFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - ModelFromString panics.
func ModelFromString(s string) Model {
	res, err := ModelFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var modelRegexCache = map[string]*regexp.Regexp{}
var modelRegexMutex = &sync.RWMutex{}

// ModelFromStringE does the same as ModelFromString does, but instead of
// panic, returns an error.
func ModelFromStringE(s string) (Model, error) {
	state := getState()
	if v, ok := state.ModelValues[s]; ok {
		return Model(v), nil
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

	for _, aliasStruct := range state.ModelAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		modelRegexMutex.RLock()
		re, ok := modelRegexCache[alias]
		modelRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			modelRegexMutex.Lock()
			modelRegexCache[alias] = re
			modelRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return Model(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "Model ", Value: s}
}

// UnmarshalText - populates Model using provided text
func (v *Model) UnmarshalText(text []byte) error {
	value, ok := getState().ModelValues[string(text)]
	if !ok {
		return errors.New("cannot find value for Model")
	}

	*v = Model(value)
	return nil
}

// MarshalText - converts Model into text
func (v Model) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of Model
func (v Model) String() string {
	str, ok := getState().ModelNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for Model"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v Model) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["MODEL"][v.String()]
}

// MarshalJSON encodes Model value into json value.
func (v Model) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into Model.
func (v *Model) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
