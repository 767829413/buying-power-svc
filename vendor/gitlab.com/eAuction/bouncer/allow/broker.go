package allow

import (
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/go/xdr"
)

func IsBroker(claims *jwt.Claims) bool {
	return claims.AccountType == int(xdr.AccountTypeAccountTypeBroker)
}

type Broker struct {
	Address  string
	Platform string
}

func (r Broker) Check(claims jwt.Claims) error {
	if !IsBroker(&claims) {
		return bouncer.ErrNotAllowed
	}

	if !IsAddress(&claims, r.Address) {
		return bouncer.ErrNotAllowed
	}

	if !IsPlatformTrusted(&claims, r.Platform) {
		return bouncer.ErrNotAllowed
	}

	return nil
}
