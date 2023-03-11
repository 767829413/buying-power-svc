/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PostTransaction struct {
	Key
	Attributes PostTransactionAttributes `json:"attributes"`
}
type PostTransactionResponse struct {
	Data     PostTransaction `json:"data"`
	Included Included        `json:"included"`
}

type PostTransactionListResponse struct {
	Data     []PostTransaction `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustPostTransaction - returns PostTransaction from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPostTransaction(key Key) *PostTransaction {
	var postTransaction PostTransaction
	if c.tryFindEntry(key, &postTransaction) {
		return &postTransaction
	}
	return nil
}
