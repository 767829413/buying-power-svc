package enumer

// Maker returns Maker enum of the model.
func (v Model) Maker() Maker {
	return Maker(getState().ModelsToMakers[int32(v)])
}

// Class returns Class (if exists) of the model.
func (v Model) Class() *Class {
	classValue, ok := getState().ModelsToClasses[int32(v)]
	if !ok {
		return nil
	}
	class := Class(classValue)
	return &class
}
