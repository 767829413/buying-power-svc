/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type WithdrawalPatch struct {
	Key
	Attributes *WithdrawalPatchAttributes `json:"attributes,omitempty"`
}
type WithdrawalPatchResponse struct {
	Data     WithdrawalPatch `json:"data"`
	Included Included        `json:"included"`
}

type WithdrawalPatchListResponse struct {
	Data     []WithdrawalPatch `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustWithdrawalPatch - returns WithdrawalPatch from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustWithdrawalPatch(key Key) *WithdrawalPatch {
	var withdrawalPatch WithdrawalPatch
	if c.tryFindEntry(key, &withdrawalPatch) {
		return &withdrawalPatch
	}
	return nil
}
