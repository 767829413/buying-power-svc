package enumer

// Classes return all know classes for the maker.
func (v Maker) Classes() []Class {
	classValues := getState().MakersToClasses[int32(v)]
	if len(classValues) == 0 {
		return nil
	}
	classes := make([]Class, 0, len(classValues))
	for _, value := range classValues {
		classes = append(classes, Class(value))
	}
	return classes
}

// Models returns all known models for the maker.
func (v Maker) Models() []Model {
	modelValues := getState().MakersToModels[int32(v)]

	if len(modelValues) == 0 {
		return nil
	}
	models := make([]Model, 0, len(modelValues))
	for _, value := range modelValues {
		models = append(models, Model(value))
	}
	return models
}
