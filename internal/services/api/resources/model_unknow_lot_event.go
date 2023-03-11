/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UnknowLotEvent struct {
	Key
	Attributes    UnknowLotEventAttributes    `json:"attributes"`
	Relationships UnknowLotEventRelationships `json:"relationships"`
}
type UnknowLotEventResponse struct {
	Data     UnknowLotEvent `json:"data"`
	Included Included       `json:"included"`
}

type UnknowLotEventListResponse struct {
	Data     []UnknowLotEvent `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustUnknowLotEvent - returns UnknowLotEvent from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUnknowLotEvent(key Key) *UnknowLotEvent {
	var unknowLotEvent UnknowLotEvent
	if c.tryFindEntry(key, &unknowLotEvent) {
		return &unknowLotEvent
	}
	return nil
}
