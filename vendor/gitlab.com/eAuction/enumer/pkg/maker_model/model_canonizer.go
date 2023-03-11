package maker_model

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var ErrUnknownModel = errors.New("failed to find canonical form")

// ModelCanonizer - provides methods to convert aliases to canonical form
type ModelCanonizer struct {
	lexemizer *ModelLexemizer
	makers    map[string][]modelAlias
}

type MakerStopWordsRecord struct {
	Maker             string
	StopWords         []string
	ExcludeFromCommon []string
}

type MakerModelRecord struct {
	NormMaker      string
	NormModelRegex string
	MakerModel
}

// NewModelCanonizer - creates new instance of ModelCanonizer
func NewModelCanonizer(makerModels []MakerModelRecord, commonStopWords []string, stopWordRecords []MakerStopWordsRecord) (*ModelCanonizer, error) {
	var err error
	var c = ModelCanonizer{
		makers: map[string][]modelAlias{},
	}
	c.lexemizer, err = NewModelLexemizer(commonStopWords, stopWordRecords)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new lexemizer")
	}

	for _, record := range makerModels {
		maker := Maker(record.NormMaker)
		rawRegexp := record.NormModelRegex
		if c.lexemizer.ContainsStopWords(maker, Model(rawRegexp)) {
			return nil, errors.From(errors.New("raw regexp contains stop words, so it would never match"), logan.F{
				"Maker":      maker,
				"raw_regexp": rawRegexp,
			})
		}

		modelAlias, err := NewModelAlias(record.MakerModel, rawRegexp)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create Model alias")
		}
		c.makers[maker.Normalized()] = append(c.makers[maker.Normalized()], modelAlias)
	}

	return &c, nil
}

// Canonize - converts model to canonical form
func (c ModelCanonizer) Canonize(maker Maker, model Model) (makerModel MakerModel, regexp string, err error) {
	modeAliases, ok := c.makers[maker.Normalized()]
	if !ok {
		return MakerModel{}, "", errors.From(ErrUnknownMaker, logan.F{
			"Maker": maker,
		})
	}

	modelLexeme := c.lexemizer.ModelToLexeme(maker, model)
	normalizedModelLexeme := modelLexeme.Normalised()
	for _, modelAlias := range modeAliases {
		if modelAlias.MatchName(normalizedModelLexeme) {
			return modelAlias.MakerModel, modelAlias.regexp.String(), nil
		}
	}

	return MakerModel{}, "", errors.From(ErrUnknownModel, logan.F{
		"Maker":        maker,
		"Model":        model,
		"model_lexeme": modelLexeme,
	})
}
