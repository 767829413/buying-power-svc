package maker_model

import (
	"regexp"
	"strings"
	"sync"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// ModelLexemizer allows to convert Model into lexeme
type ModelLexemizer struct {
	lock            *sync.RWMutex
	commonStopWords []string
	makerStopWords  map[string]*regexp.Regexp
}

// NewModelLexemizer returns new instance of ModelLexemizer
func NewModelLexemizer(commonStopWords []string, stopWords []MakerStopWordsRecord) (*ModelLexemizer, error) {
	var err error
	result := map[string]*regexp.Regexp{}
	for _, makerStopWords := range stopWords {
		makerName := Maker(makerStopWords.Maker).Normalized()
		if makerName == "" {
			return nil, ErrUnknownMaker
		}
		if _, exists := result[makerStopWords.Maker]; exists {
			return nil, errors.From(errors.New("conflict"), logan.F{
				"Maker": makerName,
			})
		}

		result[makerName], err = newMakersStopWords(makerName, commonStopWords, makerStopWords.StopWords, makerStopWords.ExcludeFromCommon)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create Maker stop words", logan.F{
				"Maker": makerName,
			})
		}
	}

	return &ModelLexemizer{
		makerStopWords:  result,
		commonStopWords: commonStopWords,
		lock:            &sync.RWMutex{},
	}, nil
}

// ModelToLexeme converts Model to lexeme
func (l ModelLexemizer) ModelToLexeme(maker Maker, model Model) Model {
	exp := l.getRegexp(maker)
	return Model(exp.ReplaceAllString(model.Normalised(), ""))
}

// ContainsStopWords returns true if normalized Model name contains one of the stop words
func (l ModelLexemizer) ContainsStopWords(maker Maker, model Model) bool {
	exp := l.getRegexp(maker)
	return exp.MatchString(model.Normalised())
}

func (l ModelLexemizer) getRegexp(maker Maker) *regexp.Regexp {
	normalized := maker.Normalized()
	l.lock.Lock()
	defer l.lock.Unlock()
	makerSpecific, ok := l.makerStopWords[normalized]
	if ok {
		return makerSpecific
	}

	result := append(l.commonStopWords, normalized)

	re, err := stopWordsToRegexp(result)
	if err != nil {
		panic(errors.Wrap(err, "failed to create stop words regexp"))
	}

	l.makerStopWords[normalized] = re
	return re
}

func stopWordsToRegexp(stopWords []string) (*regexp.Regexp, error) {
	var nonEmpty []string
	for _, stopWord := range stopWords {
		if stopWord == "" {
			continue
		}
		nonEmpty = append(nonEmpty, stopWord)
	}
	rawExp := strings.Join(nonEmpty, "|")
	exp, err := regexp.Compile(rawExp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile regexp for stop words")
	}

	return exp, nil
}

func newMakersStopWords(name string, common, makerSpecific, exclude []string) (*regexp.Regexp, error) {
	common = append(common, Maker(name).Normalized())
	common = sub(common, exclude)
	makerSpecific = append(makerSpecific, common...)
	exp, err := stopWordsToRegexp(makerSpecific)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert stop words to regexp")
	}

	return exp, nil
}

func sub(as, bs []string) []string {
	var result []string
	for _, a := range as {
		shouldInclude := true
		for _, b := range bs {
			if a == b {
				shouldInclude = false
				break
			}
		}

		if shouldInclude {
			result = append(result, a)
		}
	}

	return result
}
