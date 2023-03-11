package alias

import (
	"regexp"
	"strings"
	"sync"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//ModelLexemizer - allows to convert Model into lexeme
type ModelLexemizer struct {
	lock            *sync.Mutex
	commonStopWords []string
	makerStopWords  map[string]*regexp.Regexp
}

//NewModelLexemizer - returns new instance of ModelLexemizer
func NewModelLexemizer() (*ModelLexemizer, error) {
	rawStopWords, err := readBoxData("model_stop_words.csv")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read stop words")
	}

	if len(rawStopWords) < 2 {
		return nil, errors.New("expected stop words to contain at least 2 lines: header and common")
	}

	commonStopWords := strings.Split(rawStopWords[1][1], " ")
	rawStopWords = rawStopWords[2:]
	result := map[string]*regexp.Regexp{}
	for _, rawMakerStopWords := range rawStopWords {
		makerName := Maker(rawMakerStopWords[0]).Normalized()
		if _, exists := result[makerName]; exists {
			return nil, errors.From(errors.New("conflict"), logan.F{
				"Maker": makerName,
			})
		}

		stopWords := strings.Split(rawMakerStopWords[1], " ")
		excludeFromCommon := strings.Split(rawMakerStopWords[2], " ")
		result[makerName], err = newMakersStopWords(makerName, commonStopWords, stopWords, excludeFromCommon)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create Maker stop words", logan.F{
				"Maker": makerName,
			})
		}
	}

	return &ModelLexemizer{
		makerStopWords:  result,
		commonStopWords: commonStopWords,
		lock:            &sync.Mutex{},
	}, nil
}

//ModelToLexeme - converts Model to lexeme
func (l ModelLexemizer) ModelToLexeme(maker Maker, model Model) Model {
	exp := l.getRegexp(maker, model)
	return Model(exp.ReplaceAllString(model.Normalised(), ""))
}

//ContainsStopWords - returns true if normalized Model name contains one of the stop words
func (l ModelLexemizer) ContainsStopWords(maker Maker, model Model) bool {
	exp := l.getRegexp(maker, model)
	return exp.MatchString(model.Normalised())
}

func (l ModelLexemizer) getRegexp(maker Maker, model Model) *regexp.Regexp {
	l.lock.Lock()
	defer l.lock.Unlock()
	makerSpecific, ok := l.makerStopWords[maker.Normalized()]
	if ok {
		return makerSpecific
	}

	result := append(l.commonStopWords, maker.Normalized())
	var err error
	l.makerStopWords[maker.Normalized()], err = stopWordsToRegexp(result)
	if err != nil {
		panic(errors.Wrap(err, "failed to create stop words regexp"))
	}

	return l.makerStopWords[maker.Normalized()]
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
