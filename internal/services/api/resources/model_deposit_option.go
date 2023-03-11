/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DepositOption struct {
	Key
	Attributes    DepositOptionAttributes    `json:"attributes"`
	Relationships DepositOptionRelationships `json:"relationships"`
}
type DepositOptionResponse struct {
	Data     DepositOption `json:"data"`
	Included Included      `json:"included"`
}

type DepositOptionListResponse struct {
	Data     []DepositOption `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustDepositOption - returns DepositOption from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDepositOption(key Key) *DepositOption {
	var depositOption DepositOption
	if c.tryFindEntry(key, &depositOption) {
		return &depositOption
	}
	return nil
}
