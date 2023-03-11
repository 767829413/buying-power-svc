/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Lot struct {
	Key
	Attributes    LotAttributes    `json:"attributes"`
	Relationships LotRelationships `json:"relationships"`
}
type LotResponse struct {
	Data     Lot      `json:"data"`
	Included Included `json:"included"`
}

type LotListResponse struct {
	Data     []Lot    `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustLot - returns Lot from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLot(key Key) *Lot {
	var lot Lot
	if c.tryFindEntry(key, &lot) {
		return &lot
	}
	return nil
}
