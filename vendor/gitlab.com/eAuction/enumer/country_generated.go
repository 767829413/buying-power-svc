// GENERATED. DO NOT EDIT.

package enumer

import (
	"encoding/json"
	"regexp"
	"sync"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// Country is a public version of "Country" type of enum value.
type Country int32

// IsKnownCountry returns true if provided type is a valid
// value of Country enum.
func IsKnownCountry(v int32) bool {
	_, ok := getState().CountryNames[v]
	return ok
}

// CountryFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - CountryFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - CountryFromString panics.
func CountryFromString(s string) Country {
	res, err := CountryFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var countryRegexCache = map[string]*regexp.Regexp{}
var countryRegexMutex = &sync.RWMutex{}

// CountryFromStringE does the same as CountryFromString does, but instead of
// panic, returns an error.
func CountryFromStringE(s string) (Country, error) {
	state := getState()
	if v, ok := state.CountryValues[s]; ok {
		return Country(v), nil
	}

	s, err := normalizeCountryE(s)

	if err != nil {
		return -1, errors.Wrap(err, "failed to normalize string", logan.F{
			"string": s,
		})
	}

	if s == "" {
		s = "other"
	}

	for _, aliasStruct := range state.CountryAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		countryRegexMutex.RLock()
		re, ok := countryRegexCache[alias]
		countryRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			countryRegexMutex.Lock()
			countryRegexCache[alias] = re
			countryRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return Country(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "Country ", Value: s}
}

// UnmarshalText - populates Country using provided text
func (v *Country) UnmarshalText(text []byte) error {
	value, ok := getState().CountryValues[string(text)]
	if !ok {
		return errors.New("cannot find value for Country")
	}

	*v = Country(value)
	return nil
}

// MarshalText - converts Country into text
func (v Country) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of Country
func (v Country) String() string {
	str, ok := getState().CountryNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for Country"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v Country) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["COUNTRY"][v.String()]
}

// MarshalJSON encodes Country value into json value.
func (v Country) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into Country.
func (v *Country) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
