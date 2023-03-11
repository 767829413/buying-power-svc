package maker_model

import (
	"gitlab.com/eAuction/enumer/pkg/normalization"
)

//Model - provides methods to get normalized representation of Model without changing it's original value
type Model string

//Normalised - returns normalized representation of Model
func (m Model) Normalised() string {
	return normalization.Lower(string(m))
}

//Maker - provides methods to get normalized representation of Model without changing it's original value
type Maker string

//Normalized - returns normalized representation of Maker
func (m Maker) Normalized() string {
	return normalization.Lower(string(m))
}
