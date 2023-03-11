package alias

import (
	"sort"
	"strings"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//ModelCanonizer - provides methods to convert aliases to canonical form
type ModelCanonizer struct {
	lexemizer *ModelLexemizer
	makers    map[string][]modelAlias
}

//NewCanonizer - creates new instance of ModelCanonizer
func NewCanonizer() (*ModelCanonizer, error) {
	return newCanonizer("maker_models.csv")
}

//NewBlackListCanonizer - create new instance of Canonizer for black list
func NewBlackListCanonizer() (*ModelCanonizer, error) {
	return newCanonizer("maker_models_black_list.csv")
}

//NewCanonizer - creates new instance of ModelCanonizer
func newCanonizer(boxName string) (*ModelCanonizer, error) {
	var err error
	var c = ModelCanonizer{
		makers: map[string][]modelAlias{},
	}
	c.lexemizer, err = NewModelLexemizer()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new lexemizer")
	}
	records, err := readBoxData(boxName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read Maker models")
	}

	for _, record := range records {
		maker := Maker(record[0])
		rawRegexp := record[1]
		if c.lexemizer.ContainsStopWords(maker, Model(rawRegexp)) {
			return nil, errors.From(errors.New("raw regexp contains stop words, so it would never match"), logan.F{
				"Maker":      maker,
				"raw_regexp": rawRegexp,
			})
		}

		var normMaker, normModel string
		if len(record) >= 3 {
			normMaker = trimSpaceUpper(record[2])
		}

		if len(record) >= 5 {
			normModel = trimSpaceUpper(record[4])
		}

		models := c.makers[maker.Normalized()]
		modelAlias, err := NewModelAlias(normMaker, normModel, rawRegexp)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create Model alias")
		}
		c.makers[maker.Normalized()] = append(models, modelAlias)
	}

	for _, models := range c.makers {
		sort.Slice(models, func(i, j int) bool {
			return models[i].Less(models[j])
		})
	}

	return &c, nil
}

//Canonize - converts model to canonical form
func (c ModelCanonizer) Canonize(maker Maker, model Model) (makerName string, modelName string, regexp string, err error) {
	modeAliases, ok := c.makers[maker.Normalized()]
	if !ok {
		return "", "", "", errors.From(errors.New("unknown maker (one of the reasons for this "+
			"error is that maker was not canonized)"), logan.F{
			"Maker": maker,
		})
	}

	modelLexeme := c.lexemizer.ModelToLexeme(maker, model)
	normalizedModelLexeme := modelLexeme.Normalised()
	for _, modelAlias := range modeAliases {
		if modelAlias.MatchName(normalizedModelLexeme) {
			return modelAlias.Maker, modelAlias.Model, modelAlias.regexp.String(), nil
		}
	}

	return "", "", "", errors.From(errors.New("failed to find canonical form"), logan.F{
		"Maker":        maker,
		"Model":        model,
		"model_lexeme": modelLexeme,
	})
}

func trimSpaceUpper(s string) string {
	return strings.TrimSpace(strings.ToUpper(s))
}
