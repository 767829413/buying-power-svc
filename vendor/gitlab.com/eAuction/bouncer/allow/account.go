package allow

import (
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
)

// Account allows to check constraints not bound to particular account type.
type Account struct {
	Platform string
	Address  string
}

func (a Account) Check(claims jwt.Claims) error {
	if !IsAddress(&claims, a.Address) {
		return bouncer.ErrNotAllowed
	}

	if !IsPlatformTrusted(&claims, a.Platform) {
		return bouncer.ErrNotAllowed
	}

	return nil
}
