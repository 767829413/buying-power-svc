/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EnvelopeRelationships struct {
	Bid             *Relation           `json:"bid,omitempty"`
	BonusTime       *Relation           `json:"bonus_time,omitempty"`
	CurrentLot      *Relation           `json:"current_lot,omitempty"`
	Error           *Relation           `json:"error,omitempty"`
	Ok              *Relation           `json:"ok,omitempty"`
	Participations  *RelationCollection `json:"participations,omitempty"`
	SaleEnd         *Relation           `json:"sale_end,omitempty"`
	Sold            *Relation           `json:"sold,omitempty"`
	UnknownLotEvent *Relation           `json:"unknown_lot_event,omitempty"`
}
