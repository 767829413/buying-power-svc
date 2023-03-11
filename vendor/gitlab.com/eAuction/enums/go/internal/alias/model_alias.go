package alias

import (
	"regexp"

	"github.com/pkg/errors"
)

//modelAlias - represents alias for models
type modelAlias struct {
	regexp *regexp.Regexp
	Model  string
	Maker  string
}

//NewModelAlias -
func NewModelAlias(maker, name, rawRegexp string) (modelAlias, error) {
	exp, err := regexp.Compile(rawRegexp)
	if err != nil {
		return modelAlias{}, errors.Wrap(err, "failed to compile regexp")
	}

	return modelAlias{
		regexp: exp,
		Model:  name,
		Maker:  maker,
	}, nil
}

//MatchName - returns true if name matches regexp
func (m modelAlias) MatchName(name string) bool {
	return m.regexp.MatchString(name)
}

//Less - returns true if m < 0
func (m modelAlias) Less(o modelAlias) bool {
	cmp := len(m.regexp.String()) - len(o.regexp.String())
	switch {
	case cmp < 0:
		return false
	case cmp > 0:
		return true
	default:
		return m.regexp.String() > o.regexp.String()
	}
}
