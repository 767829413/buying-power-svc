/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package exrates

type ExchangeRate struct {
	Key
	Attributes    ExchangeRateAttributes    `json:"attributes"`
	Relationships ExchangeRateRelationships `json:"relationships"`
}
type ExchangeRateResponse struct {
	Data     ExchangeRate `json:"data"`
	Included Included     `json:"included"`
}

type ExchangeRateListResponse struct {
	Data     []ExchangeRate `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustExchangeRate - returns ExchangeRate from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustExchangeRate(key Key) *ExchangeRate {
	var exchangeRate ExchangeRate
	if c.tryFindEntry(key, &exchangeRate) {
		return &exchangeRate
	}
	return nil
}
