/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DepositRelationships struct {
	Depositor   Relation            `json:"depositor"`
	Lot         *Relation           `json:"lot,omitempty"`
	Withdrawals *RelationCollection `json:"withdrawals,omitempty"`
}
