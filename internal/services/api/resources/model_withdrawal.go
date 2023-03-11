/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Withdrawal struct {
	Key
	Attributes    WithdrawalAttributes    `json:"attributes"`
	Relationships WithdrawalRelationships `json:"relationships"`
}
type WithdrawalResponse struct {
	Data     Withdrawal `json:"data"`
	Included Included   `json:"included"`
}

type WithdrawalListResponse struct {
	Data     []Withdrawal `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustWithdrawal - returns Withdrawal from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustWithdrawal(key Key) *Withdrawal {
	var withdrawal Withdrawal
	if c.tryFindEntry(key, &withdrawal) {
		return &withdrawal
	}
	return nil
}
