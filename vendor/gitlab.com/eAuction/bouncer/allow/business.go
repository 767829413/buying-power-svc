package allow

import (
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/go/xdr"
)

func IsBusiness(claims jwt.Claims) bool {
	return claims.AccountType == int(xdr.AccountTypeAccountTypeBusiness)
}

// Business account rule.
// ID is optional and won't be used if empty.
type Business struct {
	ID string
}

func (b Business) Check(claims jwt.Claims) error {
	if claims.IsUnverified || !IsBusiness(claims) || (b.ID != "" && claims.AccountID != b.ID) {
		return bouncer.ErrNotAllowed
	}

	return nil
}
