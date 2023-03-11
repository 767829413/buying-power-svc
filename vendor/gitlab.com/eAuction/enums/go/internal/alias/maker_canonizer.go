package alias

//MakerCanonizer - converts maker to it's canonical form
type MakerCanonizer struct {
}

//NewMakerCanonizer - creates new instance of maker canonizer
func NewMakerCanonizer() (*MakerCanonizer, error) {
	return &MakerCanonizer{}, nil
}

//Canonize - converts maker to it's canonical form. If fails to find transformation, returns maker as is
func (c MakerCanonizer) Canonize(maker Maker) (Maker, error) {
	switch maker.Normalized() {
	case "generalmotors":
		return "gmc", nil
	case "mercedes":
		return "mercedesbenz", nil
	default:
		return maker, nil
	}
}
