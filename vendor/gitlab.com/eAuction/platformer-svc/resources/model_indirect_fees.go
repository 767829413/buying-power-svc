/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type IndirectFees struct {
	Key
	Attributes IndirectFeesAttributes `json:"attributes"`
}
type IndirectFeesResponse struct {
	Data     IndirectFees `json:"data"`
	Included Included     `json:"included"`
}

type IndirectFeesListResponse struct {
	Data     []IndirectFees `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustIndirectFees - returns IndirectFees from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustIndirectFees(key Key) *IndirectFees {
	var indirectFees IndirectFees
	if c.tryFindEntry(key, &indirectFees) {
		return &indirectFees
	}
	return nil
}
