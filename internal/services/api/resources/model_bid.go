/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Bid struct {
	Key
	Attributes    BidAttributes    `json:"attributes"`
	Relationships BidRelationships `json:"relationships"`
}
type BidResponse struct {
	Data     Bid      `json:"data"`
	Included Included `json:"included"`
}

type BidListResponse struct {
	Data     []Bid    `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustBid - returns Bid from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBid(key Key) *Bid {
	var bid Bid
	if c.tryFindEntry(key, &bid) {
		return &bid
	}
	return nil
}
