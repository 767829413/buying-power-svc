/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SessionRequest struct {
	Key
	Attributes SessionRequestAttributes `json:"attributes"`
}
type SessionRequestResponse struct {
	Data     SessionRequest `json:"data"`
	Included Included       `json:"included"`
}

type SessionRequestListResponse struct {
	Data     []SessionRequest `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustSessionRequest - returns SessionRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSessionRequest(key Key) *SessionRequest {
	var sessionRequest SessionRequest
	if c.tryFindEntry(key, &sessionRequest) {
		return &sessionRequest
	}
	return nil
}
