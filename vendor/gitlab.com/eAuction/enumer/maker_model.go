package enumer

import (
	"errors"
	"fmt"
)

type MakerModel struct {
	Maker Maker
	Model Model
	Class Class
}

func (m MakerModel) String() string {
	classStr := ""
	if m.Class != 0 {
		classStr = fmt.Sprintf(" %s", m.Class)
	}
	return fmt.Sprintf("%s %s%s", m.Maker, m.Model, classStr)
}

// FindMakerModel tries to find specified maker and model using
func FindMakerModel(maker, model string) (MakerModel, error) {
	state, err := getStateE()
	if err != nil {
		return MakerModel{}, err
	}
	result, err := state.ModelFinder.Find(maker, model)
	if err != nil {
		return MakerModel{}, err
	}
	return MakerModel{
		Maker: Maker(result.Maker),
		Model: Model(result.Model),
		Class: Class(result.Class),
	}, nil
}

// GetMakerByModelE returns Maker of the model or panics if cannot find one
func GetMakerByModelE(model Model) (Maker, error) {
	maker, ok := getState().ModelsToMakers[int32(model)]
	if !ok {
		return 0, errors.New(fmt.Sprintf("cannot find maker for given model %d", model))
	}
	return Maker(maker), nil
}
