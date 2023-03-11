package maker_model

import (
	"regexp"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/enumer/pkg/state"
)

var ErrUnknownMaker = errors.New("unknown maker")

// MakerMatcher - converts maker to it's canonical form
type canonMakerAlias struct {
	alias *regexp.Regexp
	maker int32
}

type MakerMatcher struct {
	aliases map[string]canonMakerAlias
}

// NewMakerMatcher - creates new instance of maker canonizer.
func NewMakerMatcher(makerNames map[int32]string, makerAliases []state.Alias) (*MakerMatcher, error) {
	aliases := map[string]canonMakerAlias{}
	for _, makerAlias := range makerAliases {
		alias := makerAlias.Alias
		maker := makerAlias.Value
		re, err := regexp.Compile(alias)
		if err != nil {
			return nil, errors.Wrap(err, "failed to compile alias", logan.F{"alias": alias, "maker": maker})
		}
		aliases[Maker(makerNames[maker]).Normalized()] = canonMakerAlias{
			alias: re,
			maker: maker,
		}
	}
	return &MakerMatcher{
		aliases: aliases,
	}, nil
}

// Match - converts maker to it enum value.
func (c MakerMatcher) Match(maker Maker) (int32, error) {
	makerAlias, ok := c.aliases[maker.Normalized()]
	if ok {
		return makerAlias.maker, nil
	}
	for _, makerAlias = range c.aliases {
		if makerAlias.alias.MatchString(maker.Normalized()) {
			return makerAlias.maker, nil
		}
	}
	return 0, errors.Wrap(ErrUnknownMaker, "", logan.F{"maker": maker})
}
