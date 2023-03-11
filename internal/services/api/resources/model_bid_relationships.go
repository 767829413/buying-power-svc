/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BidRelationships struct {
	Bidder Relation  `json:"bidder"`
	Lot    Relation  `json:"lot"`
	Sale   *Relation `json:"sale,omitempty"`
}
