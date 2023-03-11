/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BidLimitPut struct {
	Key
	Attributes    BidLimitPutAttributes    `json:"attributes"`
	Relationships BidLimitPutRelationships `json:"relationships"`
}
type BidLimitPutResponse struct {
	Data     BidLimitPut `json:"data"`
	Included Included    `json:"included"`
}

type BidLimitPutListResponse struct {
	Data     []BidLimitPut `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustBidLimitPut - returns BidLimitPut from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBidLimitPut(key Key) *BidLimitPut {
	var bidLimitPut BidLimitPut
	if c.tryFindEntry(key, &bidLimitPut) {
		return &bidLimitPut
	}
	return nil
}
