/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ShelfFee struct {
	Key
	Attributes ShelfFeeAttributes `json:"attributes"`
}
type ShelfFeeResponse struct {
	Data     ShelfFee `json:"data"`
	Included Included `json:"included"`
}

type ShelfFeeListResponse struct {
	Data     []ShelfFee `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustShelfFee - returns ShelfFee from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustShelfFee(key Key) *ShelfFee {
	var shelfFee ShelfFee
	if c.tryFindEntry(key, &shelfFee) {
		return &shelfFee
	}
	return nil
}
