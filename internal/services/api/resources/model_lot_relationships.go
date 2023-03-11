/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type LotRelationships struct {
	Organizer Relation `json:"organizer"`
	// Active participations of the lot
	Participations *map[string]interface{} `json:"participations,omitempty"`
	Platform       Relation                `json:"platform"`
	Sale           *Relation               `json:"sale,omitempty"`
	Vehicle        Relation                `json:"vehicle"`
}
