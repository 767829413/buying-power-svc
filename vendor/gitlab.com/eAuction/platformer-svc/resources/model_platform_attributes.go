/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PlatformAttributes struct {
	// URL to the web client of the platform
	ClientUrl string `json:"client_url"`
	// defines unique humanreadable identifier of the platform that is the same for all envs
	Code string `json:"code"`
	// List of lot organizers.
	Corporates *[]string `json:"corporates,omitempty"`
	// Same as `country_code` but in ISO 3166-1 alpha-3 format
	CountryAlpha3 string `json:"country_alpha3"`
	// Country code where platform users live.
	CountryCode string `json:"country_code"`
	// List of platforms whose direct sales can be shown on ours one.
	DirectSalesWhitelist *[]string `json:"direct_sales_whitelist,omitempty"`
	// Defines feature flags of the platform to be used by front end. (Do not add backend related feature flags!)
	FrontFeatureFlags map[string]interface{} `json:"front_feature_flags"`
	Logos             PlatformLogos          `json:"logos"`
	// Name of platform just to make it easy to identify.
	Name string `json:"name"`
	// Type of platform.
	PlatformType string `json:"platform_type"`
}
