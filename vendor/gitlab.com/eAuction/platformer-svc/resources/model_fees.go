/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Fees struct {
	Key
	Attributes FeesAttributes `json:"attributes"`
}
type FeesResponse struct {
	Data     Fees     `json:"data"`
	Included Included `json:"included"`
}

type FeesListResponse struct {
	Data     []Fees   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustFees - returns Fees from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustFees(key Key) *Fees {
	var fees Fees
	if c.tryFindEntry(key, &fees) {
		return &fees
	}
	return nil
}
