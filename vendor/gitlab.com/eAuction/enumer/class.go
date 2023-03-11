package enumer

// Maker returns maker for the class.
func (v Class) Maker() Maker {
	return Maker(getState().ClassesToMakers[int32(v)])

}

// Models returns all known models of the class.
func (v Class) Models() []Model {
	modelValues := getState().ClassesToModels[int32(v)]

	if len(modelValues) == 0 {
		return nil
	}
	models := make([]Model, 0, len(modelValues))
	for _, value := range modelValues {
		models = append(models, Model(value))
	}
	return models
}
