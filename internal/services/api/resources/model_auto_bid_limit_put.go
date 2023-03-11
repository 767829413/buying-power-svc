/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AutoBidLimitPut struct {
	Key
	Attributes    AutoBidLimitPutAttributes    `json:"attributes"`
	Relationships AutoBidLimitPutRelationships `json:"relationships"`
}
type AutoBidLimitPutResponse struct {
	Data     AutoBidLimitPut `json:"data"`
	Included Included        `json:"included"`
}

type AutoBidLimitPutListResponse struct {
	Data     []AutoBidLimitPut `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustAutoBidLimitPut - returns AutoBidLimitPut from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAutoBidLimitPut(key Key) *AutoBidLimitPut {
	var autoBidLimitPut AutoBidLimitPut
	if c.tryFindEntry(key, &autoBidLimitPut) {
		return &autoBidLimitPut
	}
	return nil
}
