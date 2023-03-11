package normalization

import (
	"regexp"
	"strings"
	"sync"
	"unicode"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var lowerRegExp = regexp.MustCompile("[^a-zA-Z0-9]+")

// Lower removes all non-alphanumeric symbols and casts it to lower case.
func Lower(s string) string {
	return strings.ToLower(lowerRegExp.ReplaceAllString(s, ""))
}

type normalizer struct {
	lock           sync.Mutex
	transformChain transform.Transformer
	alphanumRe     *regexp.Regexp
}

type mnSet struct{}

// Contains - returns true if r is non spacing mark.
func (s mnSet) Contains(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func (n *normalizer) normalize(s string) (string, error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	// transform string to unicode NFC normalized form (so é becomes e, Ç -> C etc)
	s, _, err := transform.String(n.transformChain, s)
	if err != nil {
		return "", errors.Wrap(err, "failed to transform string")
	}

	s = n.alphanumRe.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	return s, nil
}

var defaultNormalizer = &normalizer{
	transformChain: transform.Chain(norm.NFD, runes.Remove(mnSet{}), norm.NFC),
	alphanumRe:     lowerRegExp,
}

// Normalize takes the value from outside and converts it to normalized
// format (alphanumeric in lower case). Panics if failed to normalize.
func Normalize(s string) string {
	res, err := NormalizeE(s)
	if err != nil {
		panic(err)
	}

	return res
}

// NormalizeE does the same as Normalize does, but
// instead of panic, returns error.
func NormalizeE(s string) (string, error) {
	return defaultNormalizer.normalize(s)
}
