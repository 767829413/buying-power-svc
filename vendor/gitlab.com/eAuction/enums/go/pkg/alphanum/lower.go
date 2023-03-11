package alphanum

import (
	"regexp"
	"strings"
	"sync"
	"unicode"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//Transformer - allows to transform string
type Transformer struct {
	t      transform.Transformer
	regexp *regexp.Regexp
	lock   *sync.Mutex
}

//Lower - converts s to lower case, converts all umlauts to their equivalent (Â->A), removes non aplha numeric chars
func (t Transformer) Lower(s string) string {
	t.lock.Lock()
	defer t.lock.Unlock()
	transformedS, _, err := transform.String(t.t, s)
	if err != nil {
		panic(errors.Wrap(err, "failed to transform string", logan.F{
			"raw": s,
		}))
	}
	alphaNumS := t.regexp.ReplaceAllString(transformedS, "")
	return strings.ToLower(alphaNumS)
}

//NewTransformer - creates new transformer
func NewTransformer() *Transformer {
	return &Transformer{
		t:      transform.Chain(norm.NFD, runes.Remove(mnSet{}), norm.NFC),
		regexp: regexp.MustCompile("[^a-zA-Z0-9]+"),
		lock:   &sync.Mutex{},
	}
}

var defaultTransformer = NewTransformer()

//Lower - converts s to lower case, converts all umlauts to their equivalent (Â->A), removes non aplha numeric chars
func Lower(rawS string) string {
	return defaultTransformer.Lower(rawS)
}

type mnSet struct {
}

//Contains - returns true if r is non spacing mark
func (s mnSet) Contains(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}
