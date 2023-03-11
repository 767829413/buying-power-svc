/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Sold struct {
	Key
	Attributes    SoldAttributes    `json:"attributes"`
	Relationships SoldRelationships `json:"relationships"`
}
type SoldResponse struct {
	Data     Sold     `json:"data"`
	Included Included `json:"included"`
}

type SoldListResponse struct {
	Data     []Sold   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustSold - returns Sold from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSold(key Key) *Sold {
	var sold Sold
	if c.tryFindEntry(key, &sold) {
		return &sold
	}
	return nil
}
