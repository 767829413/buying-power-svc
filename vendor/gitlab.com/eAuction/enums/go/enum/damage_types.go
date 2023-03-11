package enum

import "gitlab.com/eAuction/enums/go/pkg/alphanum"

func damageType(s string) bool {
	s = alphanum.Lower(s)
	for _, t := range damageTypes {
		if alphanum.Lower(t.Name) == s {
			return true
		}
	}

	return false
}
