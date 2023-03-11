/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Deposit struct {
	Key
	Attributes    DepositAttributes    `json:"attributes"`
	Relationships DepositRelationships `json:"relationships"`
}
type DepositResponse struct {
	Data     Deposit  `json:"data"`
	Included Included `json:"included"`
}

type DepositListResponse struct {
	Data     []Deposit `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustDeposit - returns Deposit from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDeposit(key Key) *Deposit {
	var deposit Deposit
	if c.tryFindEntry(key, &deposit) {
		return &deposit
	}
	return nil
}
