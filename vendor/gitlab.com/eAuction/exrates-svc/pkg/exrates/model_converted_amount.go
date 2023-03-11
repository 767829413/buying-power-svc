/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package exrates

type ConvertedAmount struct {
	Key
	Attributes ConvertedAmountAttributes `json:"attributes"`
}
type ConvertedAmountResponse struct {
	Data     ConvertedAmount `json:"data"`
	Included Included        `json:"included"`
}

type ConvertedAmountListResponse struct {
	Data     []ConvertedAmount `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustConvertedAmount - returns ConvertedAmount from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustConvertedAmount(key Key) *ConvertedAmount {
	var convertedAmount ConvertedAmount
	if c.tryFindEntry(key, &convertedAmount) {
		return &convertedAmount
	}
	return nil
}
