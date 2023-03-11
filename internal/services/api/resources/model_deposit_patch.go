/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DepositPatch struct {
	Key
	Attributes DepositPatchAttributes `json:"attributes"`
}
type DepositPatchResponse struct {
	Data     DepositPatch `json:"data"`
	Included Included     `json:"included"`
}

type DepositPatchListResponse struct {
	Data     []DepositPatch `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustDepositPatch - returns DepositPatch from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDepositPatch(key Key) *DepositPatch {
	var depositPatch DepositPatch
	if c.tryFindEntry(key, &depositPatch) {
		return &depositPatch
	}
	return nil
}
