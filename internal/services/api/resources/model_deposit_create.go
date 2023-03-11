/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DepositCreate struct {
	Key
	Attributes    DepositCreateAttributes    `json:"attributes"`
	Relationships DepositCreateRelationships `json:"relationships"`
}
type DepositCreateResponse struct {
	Data     DepositCreate `json:"data"`
	Included Included      `json:"included"`
}

type DepositCreateListResponse struct {
	Data     []DepositCreate `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustDepositCreate - returns DepositCreate from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDepositCreate(key Key) *DepositCreate {
	var depositCreate DepositCreate
	if c.tryFindEntry(key, &depositCreate) {
		return &depositCreate
	}
	return nil
}
