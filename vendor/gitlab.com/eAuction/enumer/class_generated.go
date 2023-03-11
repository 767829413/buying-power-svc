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

// Class is a public version of "Class" type of enum value.
type Class int32

// IsKnownClass returns true if provided type is a valid
// value of Class enum.
func IsKnownClass(v int32) bool {
	_, ok := getState().ClassNames[v]
	return ok
}

// ClassFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - ClassFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - ClassFromString panics.
func ClassFromString(s string) Class {
	res, err := ClassFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var classRegexCache = map[string]*regexp.Regexp{}
var classRegexMutex = &sync.RWMutex{}

// ClassFromStringE does the same as ClassFromString does, but instead of
// panic, returns an error.
func ClassFromStringE(s string) (Class, error) {
	state := getState()
	if v, ok := state.ClassValues[s]; ok {
		return Class(v), nil
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

	for _, aliasStruct := range state.ClassAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		classRegexMutex.RLock()
		re, ok := classRegexCache[alias]
		classRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			classRegexMutex.Lock()
			classRegexCache[alias] = re
			classRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return Class(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "Class ", Value: s}
}

// UnmarshalText - populates Class using provided text
func (v *Class) UnmarshalText(text []byte) error {
	value, ok := getState().ClassValues[string(text)]
	if !ok {
		return errors.New("cannot find value for Class")
	}

	*v = Class(value)
	return nil
}

// MarshalText - converts Class into text
func (v Class) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of Class
func (v Class) String() string {
	str, ok := getState().ClassNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for Class"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v Class) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["CLASS"][v.String()]
}

// MarshalJSON encodes Class value into json value.
func (v Class) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into Class.
func (v *Class) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
