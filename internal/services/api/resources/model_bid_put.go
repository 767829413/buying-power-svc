/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BidPut struct {
	Key
	Attributes    BidPutAttributes    `json:"attributes"`
	Relationships BidPutRelationships `json:"relationships"`
}
type BidPutResponse struct {
	Data     BidPut   `json:"data"`
	Included Included `json:"included"`
}

type BidPutListResponse struct {
	Data     []BidPut `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustBidPut - returns BidPut from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBidPut(key Key) *BidPut {
	var bidPut BidPut
	if c.tryFindEntry(key, &bidPut) {
		return &bidPut
	}
	return nil
}
