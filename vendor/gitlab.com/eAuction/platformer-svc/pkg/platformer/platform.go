package platformer

import (
	"gitlab.com/eAuction/go/amount"
	"net/url"
)

//Logos -
type Logos struct {
	Dark  *url.URL `fig:"dark,required"`
	Light *url.URL `fig:"light,required"`
	Icon  *url.URL `fig:"icon,required"`
}

type TransportationFee struct {
	Min      string `fig:"min,required"`
	Max      string `fig:"max,required"`
	Currency string `fig:"currency,required"`
}

//ShelfFee - defines fee charged by shelf network
type ShelfFee struct {
	Fixed    amount.Fiat `fig:"fixed,required"`
	Currency string      `fig:"currency,required"`
}

// Platform represents a single platform.
type Platform struct {
	ID                   string
	Corporates           []string               `fig:"corporates"`
	Code                 string                 `fig:"code,required"`
	DirectSalesWhitelist []string               `fig:"direct_sales_whitelist"`
	Name                 string                 `fig:"name,required"`
	CountryCode          string                 `fig:"country,required"`
	CountryAlpha3        string                 `fig:"country3,required"`
	PlatformType         PlatformType           `fig:"type,required"`
	Logos                Logos                  `fig:"logos,required"`
	ClientURL            *url.URL               `fig:"client_url,required"`
	FrontFeatureFlags    map[string]interface{} `fig:"front_feature_flags"`
	Fees                 *string                `fig:"fees"`
	TransportationFee    *TransportationFee     `fig:"transportation_fee"`
	ShelfFee             *ShelfFee              `fig:"shelf_fee"`
	IndirectFees         *string                `fig:"indirect_fees"`
	LiontransFees        *string                `fig:"liontrans_fees"`
}
