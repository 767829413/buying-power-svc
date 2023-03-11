package enum

import (
	"gitlab.com/eAuction/enums/go/internal/alias"
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

// TODO: do not use
func FindDamageType(raw string) (string, bool) {
	result, ok := alias.DamageTypes[alphanum.Lower(raw)]
	return result, ok
}

//EnumValue - represents generic enum value
type EnumValue struct {
	// Name - used for filtering
	Name string
	// ReadableName - used by frontend for localizations
	ReadableName string
}

type (
	MakerSet      map[string]struct{}
	MakerModelSet map[[2]string]struct{}
)

func (s MakerSet) Contains(maker string) bool {
	_, ok := s[alphanum.Lower(maker)]
	return ok
}

func (s MakerModelSet) Contains(maker, model string) bool {
	_, ok := s[[2]string{alphanum.Lower(maker), alphanum.Lower(model)}]
	return ok
}

func Makers() []Maker {
	return makerModels
}

func DamageTypes() []EnumValue {
	return damageTypes
}

func DriveTypes() []EnumValue {
	return driveTypes
}

func TransmissionTypes() []EnumValue {
	return transmissions
}

func FuelTypes() []EnumValue {
	return fuelTypes
}

func BodyStyles() []EnumValue {
	return bodyStyles
}

type Model struct {
	Name  string
	Maker string
	Class string
}

type Maker struct {
	Name    string
	Classes []Class
}

type Class struct {
	Name   string
	Models []string
}
