package allow

import (
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/go/xdr"
)

func IsPlatform(claims *jwt.Claims) bool {
	return claims.AccountType == int(xdr.AccountTypeAccountTypePlatform)
}

// Platform account rule.
type Platform struct{}

func (r Platform) Check(claims jwt.Claims) error {
	if IsPlatform(&claims) {
		return nil
	}
	return bouncer.ErrNotAllowed
}
