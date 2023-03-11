/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EndSale struct {
	Key
	Attributes    EndSaleAttributes    `json:"attributes"`
	Relationships EndSaleRelationships `json:"relationships"`
}
type EndSaleResponse struct {
	Data     EndSale  `json:"data"`
	Included Included `json:"included"`
}

type EndSaleListResponse struct {
	Data     []EndSale `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustEndSale - returns EndSale from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEndSale(key Key) *EndSale {
	var endSale EndSale
	if c.tryFindEntry(key, &endSale) {
		return &endSale
	}
	return nil
}
