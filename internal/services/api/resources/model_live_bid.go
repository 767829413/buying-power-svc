/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type LiveBid struct {
	Key
	Attributes    LiveBidAttributes    `json:"attributes"`
	Relationships LiveBidRelationships `json:"relationships"`
}
type LiveBidResponse struct {
	Data     LiveBid  `json:"data"`
	Included Included `json:"included"`
}

type LiveBidListResponse struct {
	Data     []LiveBid `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustLiveBid - returns LiveBid from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLiveBid(key Key) *LiveBid {
	var liveBid LiveBid
	if c.tryFindEntry(key, &liveBid) {
		return &liveBid
	}
	return nil
}
