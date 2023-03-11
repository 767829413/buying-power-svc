/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PostRefund struct {
	Key
	Attributes *PostRefundAttributes `json:"attributes,omitempty"`
}
type PostRefundResponse struct {
	Data     PostRefund `json:"data"`
	Included Included   `json:"included"`
}

type PostRefundListResponse struct {
	Data     []PostRefund `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustPostRefund - returns PostRefund from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPostRefund(key Key) *PostRefund {
	var postRefund PostRefund
	if c.tryFindEntry(key, &postRefund) {
		return &postRefund
	}
	return nil
}
