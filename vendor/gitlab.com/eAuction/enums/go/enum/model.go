package enum

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/enums/go/internal/alias"
)

//ErrModelBlackListed - specifies that model has been blacklisted
var ErrModelBlackListed = errors.New("specified model is in black list")

//ModelFinder - finds Model by maker, model names
type ModelFinder struct {
	makerCanonizer     *alias.MakerCanonizer
	canonizer          *alias.ModelCanonizer
	blackListCanonizer *alias.ModelCanonizer
}

//NewModelFinder - creates new instance of models finder
func NewModelFinder() (*ModelFinder, error) {
	canonizer, err := alias.NewCanonizer()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create canonizer")
	}

	blackListCanonizer, err := alias.NewBlackListCanonizer()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create black list canonizer")
	}

	makerCanonizer, err := alias.NewMakerCanonizer()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create maker canonizer")
	}

	return &ModelFinder{
		canonizer:          canonizer,
		blackListCanonizer: blackListCanonizer,
		makerCanonizer:     makerCanonizer,
	}, nil
}

//Find - returns Model for specific maker, model names or error if fails to find
func (m ModelFinder) Find(makerName, modelName string) (Model, error) {
	maker, err := m.makerCanonizer.Canonize(alias.Maker(makerName))
	if err != nil {
		return Model{}, errors.Wrap(err, "failed to canonize maker")
	}
	model := alias.Model(modelName)
	_, _, regexp, err := m.blackListCanonizer.Canonize(maker, model)
	if err == nil {
		return Model{}, errors.From(ErrModelBlackListed, logan.F{
			"maker_name":     makerName,
			"model_name":     modelName,
			"blacklisted_by": regexp,
		})
	}

	canonMaker, canonModel, _, err := m.canonizer.Canonize(maker, model)
	if err != nil {
		return Model{}, errors.Wrap(err, "failed to canonize maker model", logan.F{
			"maker_name": makerName,
			"model_name": modelName,
		})
	}

	for _, makerCandidate := range Makers() {
		if makerCandidate.Name != canonMaker {
			continue
		}

		for _, classCandidate := range makerCandidate.Classes {
			for _, modelCandidate := range classCandidate.Models {
				if modelCandidate == canonModel {
					return Model{
						Name:  modelCandidate,
						Class: classCandidate.Name,
						Maker: makerCandidate.Name,
					}, nil
				}
			}
		}
	}

	return Model{}, errors.From(errors.New("failed to find model any model for canonical form"), logan.F{
		"canon_model": canonMaker,
		"canon_maker": canonModel,
	})

}
