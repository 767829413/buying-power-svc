package maker_model

import (
	"regexp"

	"github.com/pkg/errors"
)

//modelAlias - represents alias for models
type modelAlias struct {
	regexp *regexp.Regexp
	MakerModel
}

//NewModelAlias -
func NewModelAlias(makerModel MakerModel, rawRegexp string) (modelAlias, error) {
	exp, err := regexp.Compile(rawRegexp)
	if err != nil {
		return modelAlias{}, errors.Wrap(err, "failed to compile regexp")
	}

	return modelAlias{
		regexp:     exp,
		MakerModel: makerModel,
	}, nil
}

//MatchName - returns true if name matches regexp
func (m modelAlias) MatchName(name string) bool {
	return m.regexp.MatchString(name)
}
