package maker_model

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/enumer/pkg/state"
)

// ErrModelBlackListed specifies that model has been blacklisted
var ErrModelBlackListed = errors.New("specified model is in black list")

// ModelFinder finds Model by maker, model names
type ModelFinder struct {
	makerCanonizer     *MakerMatcher
	canonizer          *ModelCanonizer
	blackListCanonizer *ModelCanonizer
	makerNames         map[int32]string
}

// NewModelFinder creates new instance of models finder
func NewModelFinder(state *state.State) (*ModelFinder, error) {
	var makerStopWords []MakerStopWordsRecord
	for _, v := range state.MakerStopWords {
		makerStopWords = append(makerStopWords, MakerStopWordsRecord{
			Maker:             state.MakerNames[v.Maker],
			StopWords:         v.StopWords,
			ExcludeFromCommon: v.ExcludeFromCommon,
		})
	}

	var makersModels []MakerModelRecord
	for _, modelAlias := range state.ModelAliases {
		model := modelAlias.Value
		maker := state.ModelsToMakers[model]
		alias := modelAlias.Alias
		makersModels = append(makersModels, MakerModelRecord{
			NormMaker:      Maker(state.MakerNames[maker]).Normalized(),
			NormModelRegex: alias,
			MakerModel: MakerModel{
				Maker: maker,
				Model: model,
				Class: state.ModelsToClasses[model],
			},
		})
	}

	canonizer, err := NewModelCanonizer(makersModels, state.CommonStopWords, makerStopWords)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create canonizer")
	}

	var blacklist []MakerModelRecord
	for _, v := range state.MakerModelBlacklist {
		blacklist = append(blacklist, MakerModelRecord{
			NormMaker:      v.NormMaker,
			NormModelRegex: v.NormModelRegex,
		})
	}
	blackListCanonizer, err := NewModelCanonizer(blacklist, state.CommonStopWords, makerStopWords)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create black list canonizer")
	}

	makerCanonizer, err := NewMakerMatcher(state.MakerNames, state.MakerAliases)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create maker canonizer")
	}

	return &ModelFinder{
		canonizer:          canonizer,
		blackListCanonizer: blackListCanonizer,
		makerCanonizer:     makerCanonizer,
		makerNames:         state.MakerNames,
	}, nil
}

type MakerModel struct {
	Maker int32
	Model int32
	Class int32 // may be 0
}

// Find returns Model for specific maker, model names or error if fails to find
func (m ModelFinder) Find(makerName, modelName string) (MakerModel, error) {
	_, regexp, err := m.blackListCanonizer.Canonize(Maker(makerName), Model(modelName))
	if err == nil {
		return MakerModel{}, errors.From(ErrModelBlackListed, logan.F{
			"maker_name":     makerName,
			"model_name":     modelName,
			"blacklisted_by": regexp,
		})
	}
	canonMaker, err := m.makerCanonizer.Match(Maker(makerName))
	if err != nil {
		return MakerModel{}, errors.Wrap(err, "failed to canonize maker")
	}
	makerModel, _, err := m.canonizer.Canonize(Maker(m.makerNames[canonMaker]), Model(modelName))
	if err != nil {
		return MakerModel{}, errors.Wrap(err, "failed to canonize maker model", logan.F{
			"maker_name": makerName,
			"model_name": modelName,
		})
	}
	return makerModel, nil
}
