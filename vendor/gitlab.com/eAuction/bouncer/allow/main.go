package allow

import (
	"gitlab.com/eAuction/bouncer/jwt"
)

// IsAddress checks if claims are for provided address.
// Returns true for empty address
func IsAddress(claims *jwt.Claims, address string) bool {
	return address == "" || address == claims.AccountID
}

// IsPlatformTrusted checks if claims are trusted by given platform.
// Returns true for empty platform
func IsPlatformTrusted(claims *jwt.Claims, platform string) bool {
	if platform == "" || platform == claims.PlatformID {
		return true
	}

	for _, trustingPlatform := range claims.TrustingPlatforms {
		if trustingPlatform == platform {
			return true
		}
	}

	return false
}
