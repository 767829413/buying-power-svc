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

// Color is a public version of "Color" type of enum value.
type Color int32

// IsKnownColor returns true if provided type is a valid
// value of Color enum.
func IsKnownColor(v int32) bool {
	_, ok := getState().ColorNames[v]
	return ok
}

// ColorFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - ColorFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - ColorFromString panics.
func ColorFromString(s string) Color {
	res, err := ColorFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var colorRegexCache = map[string]*regexp.Regexp{}
var colorRegexMutex = &sync.RWMutex{}

// ColorFromStringE does the same as ColorFromString does, but instead of
// panic, returns an error.
func ColorFromStringE(s string) (Color, error) {
	state := getState()
	if v, ok := state.ColorValues[s]; ok {
		return Color(v), nil
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

	for _, aliasStruct := range state.ColorAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		colorRegexMutex.RLock()
		re, ok := colorRegexCache[alias]
		colorRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			colorRegexMutex.Lock()
			colorRegexCache[alias] = re
			colorRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return Color(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "Color ", Value: s}
}

// UnmarshalText - populates Color using provided text
func (v *Color) UnmarshalText(text []byte) error {
	value, ok := getState().ColorValues[string(text)]
	if !ok {
		return errors.New("cannot find value for Color")
	}

	*v = Color(value)
	return nil
}

// MarshalText - converts Color into text
func (v Color) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of Color
func (v Color) String() string {
	str, ok := getState().ColorNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for Color"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v Color) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["COLOR"][v.String()]
}

// MarshalJSON encodes Color value into json value.
func (v Color) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into Color.
func (v *Color) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
