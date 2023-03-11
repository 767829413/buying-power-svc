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

// BodyStyle is a public version of "BodyStyle" type of enum value.
type BodyStyle int32

// IsKnownBodyStyle returns true if provided type is a valid
// value of BodyStyle enum.
func IsKnownBodyStyle(v int32) bool {
	_, ok := getState().BodyStyleNames[v]
	return ok
}

// BodyStyleFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - BodyStyleFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - BodyStyleFromString panics.
func BodyStyleFromString(s string) BodyStyle {
	res, err := BodyStyleFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var bodyStyleRegexCache = map[string]*regexp.Regexp{}
var bodyStyleRegexMutex = &sync.RWMutex{}

// BodyStyleFromStringE does the same as BodyStyleFromString does, but instead of
// panic, returns an error.
func BodyStyleFromStringE(s string) (BodyStyle, error) {
	state := getState()
	if v, ok := state.BodyStyleValues[s]; ok {
		return BodyStyle(v), nil
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

	for _, aliasStruct := range state.BodyStyleAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		bodyStyleRegexMutex.RLock()
		re, ok := bodyStyleRegexCache[alias]
		bodyStyleRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			bodyStyleRegexMutex.Lock()
			bodyStyleRegexCache[alias] = re
			bodyStyleRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return BodyStyle(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "BodyStyle ", Value: s}
}

// UnmarshalText - populates BodyStyle using provided text
func (v *BodyStyle) UnmarshalText(text []byte) error {
	value, ok := getState().BodyStyleValues[string(text)]
	if !ok {
		return errors.New("cannot find value for BodyStyle")
	}

	*v = BodyStyle(value)
	return nil
}

// MarshalText - converts BodyStyle into text
func (v BodyStyle) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of BodyStyle
func (v BodyStyle) String() string {
	str, ok := getState().BodyStyleNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for BodyStyle"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v BodyStyle) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["BODY_STYLE"][v.String()]
}

// MarshalJSON encodes BodyStyle value into json value.
func (v BodyStyle) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into BodyStyle.
func (v *BodyStyle) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
