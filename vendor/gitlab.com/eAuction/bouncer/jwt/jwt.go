package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	StandardClaims = jwt.StandardClaims
)

var (
	NewWithClaims      = jwt.NewWithClaims
	SigningMethodES256 = jwt.SigningMethodES256
)

const (
	AudienceExternalService = "external_service"
	AudienceShelfUser       = "shelf_user"
)

type Claims struct {
	AccountID                  string   `json:"account_id"`
	AccountType                int      `json:"account_type"`
	PlatformID                 string   `json:"platform_id"`
	TrustingPlatforms          []string `json:"trusting_platforms"`
	EmailVerified              bool     `json:"email_verified"`
	PhoneVerified              bool     `json:"phone_verified"`
	ImportSuggestionsAllowed   bool     `json:"import_suggestions_allowed,omitempty"`
	IsImporterBlacklistManager bool     `json:"is_importer_blacklist_manager,omitempty"`

	IsUnverified bool `json:"is_unverified"`

	AccountPaused bool `json:"-"`

	jwt.StandardClaims `log:"-"`
}

func (c *Claims) SetClaims(claims *Claims) {
	*c = *claims
}

func (a *Claims) GetClaims() *Claims {
	return a
}
