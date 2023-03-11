package responses

import (
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/go/amount"
)

//NewDeposit - populates Deposit resource
func NewDeposit(deposit data.Deposit, withdrawals []data.Withdrawal) resources.Deposit {
	result := resources.Deposit{
		Key: resources.Key{
			ID:   deposit.ID,
			Type: resources.DEPOSITS,
		},
		Attributes: resources.DepositAttributes{
			Amount:    amount.Fiat(deposit.Amount),
			Currency:  deposit.Currency,
			CreatedAt: deposit.CreatedAt.UTC(),
			UpdatedAt: deposit.UpdatedAt.UTC(),
			State:     kafka.DepositV2_State(deposit.State).String(),
		},
		Relationships: resources.DepositRelationships{
			Depositor: resources.Relation{
				Data: &resources.Key{
					ID:   deposit.Depositor,
					Type: resources.IDENTITIES,
				},
			},
		},
	}

	if deposit.LotID != nil {
		result.Relationships.Lot = resources.Key{
			ID:   *deposit.LotID,
			Type: resources.LOTS,
		}.AsRelation()
	}

	result.Relationships.Withdrawals = &resources.RelationCollection{}
	for _, w := range withdrawals {
		result.Relationships.Withdrawals.Data = append(result.Relationships.Withdrawals.Data, resources.Key{
			ID:   w.ID,
			Type: resources.WITHDRAWALS,
		})
	}

	return result
}
