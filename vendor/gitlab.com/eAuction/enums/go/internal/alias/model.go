package alias

import (
	"gitlab.com/eAuction/enums/go/pkg/alphanum"
)

//Model - provides methods to get normalized representation of Model without changing it's original value
type Model string

//Normalised - returns normalized representation of Model
func (m Model) Normalised() string {
	return alphanum.Lower(string(m))
}
