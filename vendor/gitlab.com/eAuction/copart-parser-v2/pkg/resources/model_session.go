/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Session struct {
	Key
	Attributes    SessionAttributes    `json:"attributes"`
	Relationships SessionRelationships `json:"relationships"`
}
type SessionResponse struct {
	Data     Session  `json:"data"`
	Included Included `json:"included"`
}

type SessionListResponse struct {
	Data     []Session `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustSession - returns Session from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSession(key Key) *Session {
	var session Session
	if c.tryFindEntry(key, &session) {
		return &session
	}
	return nil
}
