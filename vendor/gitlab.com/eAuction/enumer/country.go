package enumer

import (
	"regexp"
	"strings"
)

// Alpha2 returns ISO 3166-1 alpha-2 code.
func (v Country) Alpha2() string {
	alpha2 := getState().CountryToAlpha2[int32(v)]
	if alpha2 == "" {
		panic("unknown country")
	}
	return alpha2
}

// Alpha3 returns ISO 3166-1 alpha-3 code.
func (v Country) Alpha3() string {
	return v.String()
}

// PhoneCodes returns known phone codes for the country.
func (v Country) PhoneCodes() []string {
	return getState().CountryToPhoneCodes[int32(v)]
}

var (
	redundantInCountry = regexp.MustCompile(`\s+|'|-`)
)

func normalizeCountryE(s string) (string, error) {
	return strings.ToLower(redundantInCountry.ReplaceAllString(s, "")), nil
}
