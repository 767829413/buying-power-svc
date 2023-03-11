/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PostConfirmation struct {
	Key
	Attributes *PostConfirmationAttributes `json:"attributes,omitempty"`
}
type PostConfirmationResponse struct {
	Data     PostConfirmation `json:"data"`
	Included Included         `json:"included"`
}

type PostConfirmationListResponse struct {
	Data     []PostConfirmation `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustPostConfirmation - returns PostConfirmation from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPostConfirmation(key Key) *PostConfirmation {
	var postConfirmation PostConfirmation
	if c.tryFindEntry(key, &postConfirmation) {
		return &postConfirmation
	}
	return nil
}
