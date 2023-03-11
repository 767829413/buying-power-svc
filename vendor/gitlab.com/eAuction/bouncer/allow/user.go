package allow

import (
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/go/xdr"
)

func IsUser(claims *jwt.Claims) bool {
	return claims.AccountType == int(xdr.AccountTypeAccountTypeUser) ||
		claims.AccountType == int(xdr.AccountTypeAccountTypeBusiness)
}

type User struct {
	Address string
}

func (r User) Check(claims jwt.Claims) error {
	if !IsUser(&claims) {
		return bouncer.ErrNotAllowed
	}

	if !IsAddress(&claims, r.Address) {
		return bouncer.ErrNotAllowed
	}

	return nil
}
