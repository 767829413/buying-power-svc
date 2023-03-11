// GENERATED. DO NOT EDIT.

package enumer

import (
	"encoding/json"
	"regexp"
	"sync"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// OwnershipDocument is a public version of "OwnershipDocument" type of enum value.
type OwnershipDocument int32

// IsKnownOwnershipDocument returns true if provided type is a valid
// value of OwnershipDocument enum.
func IsKnownOwnershipDocument(v int32) bool {
	_, ok := getState().OwnershipDocumentNames[v]
	return ok
}

// OwnershipDocumentFromString converts tries to figure out enum value from it's string
// representation. No need to normalize input strings - OwnershipDocumentFromString will
// normalize it by itself. If a valid name or alias passed (even not
// normalized) - the value will be returned. Otherwise - OwnershipDocumentFromString panics.
func OwnershipDocumentFromString(s string) OwnershipDocument {
	res, err := OwnershipDocumentFromStringE(s)
	if err != nil {
		panic(err)
	}
	return res
}

var ownershipDocumentRegexCache = map[string]*regexp.Regexp{}
var ownershipDocumentRegexMutex = &sync.RWMutex{}

// OwnershipDocumentFromStringE does the same as OwnershipDocumentFromString does, but instead of
// panic, returns an error.
func OwnershipDocumentFromStringE(s string) (OwnershipDocument, error) {
	state := getState()
	if v, ok := state.OwnershipDocumentValues[s]; ok {
		return OwnershipDocument(v), nil
	}

	s, err := normalizeOwnershipDocumentE(s)

	if err != nil {
		return -1, errors.Wrap(err, "failed to normalize string", logan.F{
			"string": s,
		})
	}

	if s == "" {
		s = "other"
	}

	for _, aliasStruct := range state.OwnershipDocumentAliases {
		alias := aliasStruct.Alias
		v := aliasStruct.Value
		ownershipDocumentRegexMutex.RLock()
		re, ok := ownershipDocumentRegexCache[alias]
		ownershipDocumentRegexMutex.RUnlock()
		if !ok {
			re = regexp.MustCompile(alias)
			ownershipDocumentRegexMutex.Lock()
			ownershipDocumentRegexCache[alias] = re
			ownershipDocumentRegexMutex.Unlock()
		}

		if re.MatchString(s) {
			return OwnershipDocument(v), nil
		}
	}

	return -1, UnknownAliasError{Enum: "OwnershipDocument ", Value: s}
}

// UnmarshalText - populates OwnershipDocument using provided text
func (v *OwnershipDocument) UnmarshalText(text []byte) error {
	value, ok := getState().OwnershipDocumentValues[string(text)]
	if !ok {
		return errors.New("cannot find value for OwnershipDocument")
	}

	*v = OwnershipDocument(value)
	return nil
}

// MarshalText - converts OwnershipDocument into text
func (v OwnershipDocument) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String return string representation of OwnershipDocument
func (v OwnershipDocument) String() string {
	str, ok := getState().OwnershipDocumentNames[int32(v)]
	if !ok {
		panic(errors.From(errors.New("cannot find name for OwnershipDocument"), logan.F{
			"value": int32(v),
		}))
	}
	return str
}

func (v OwnershipDocument) Translation(locale string) string {
	translationSet, ok := getState().Translations[locale]
	if !ok {
		return ""
	}
	return translationSet["OWNERSHIP_DOCUMENT"][v.String()]
}

// MarshalJSON encodes OwnershipDocument value into json value.
func (v OwnershipDocument) MarshalJSON() ([]byte, error) {
	str := v.String()
	return json.Marshal(str)
}

// UnmarshalJSON decodes json value into OwnershipDocument.
func (v *OwnershipDocument) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return v.UnmarshalText([]byte(s))
}
