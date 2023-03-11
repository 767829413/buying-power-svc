package allow

import (
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/go/xdr"
)

func IsAdmin(claims *jwt.Claims) bool {
	return claims.AccountType == int(xdr.AccountTypeAccountTypeAdmin)
}

// Admin account rule.
type Admin struct {
	// DEPRECATED: use Platform for consistency
	PlatformID string
	// Platform is optional and if is set will check
	// whether admin is trusted by the platform.
	Platform string
	// Address is optional and if set will check if address in claims matches specified here
	Address  string
}

func (r Admin) Check(claims jwt.Claims) error {
	if !IsAdmin(&claims) {
		return bouncer.ErrNotAllowed
	}

	if !IsPlatformTrusted(&claims, r.PlatformID) {
		return bouncer.ErrNotAllowed
	}

	if !IsPlatformTrusted(&claims, r.Platform) {
		return bouncer.ErrNotAllowed
	}

	if !IsAddress(&claims, r.Address) {
		return bouncer.ErrNotAllowed
	}

	return nil
}
