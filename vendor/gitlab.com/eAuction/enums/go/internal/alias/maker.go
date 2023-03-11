package alias

import "gitlab.com/eAuction/enums/go/pkg/alphanum"

//Maker - provides methods to get normalized representation of Model without changing it's original value
type Maker string

//Normalized - returns normalized representation of Maker
func (m Maker) Normalized() string {
	return alphanum.Lower(string(m))
}
