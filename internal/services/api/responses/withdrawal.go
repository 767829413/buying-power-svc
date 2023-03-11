package responses

import (
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/events/go/kafka"
)

//NewWithdrawal - populates withdrawal resource
func NewWithdrawal(w data.Withdrawal) resources.Withdrawal {
	result := resources.Withdrawal{
		Key: resources.Key{
			ID:   w.ID,
			Type: resources.WITHDRAWALS,
		},
		Attributes: resources.WithdrawalAttributes{
			CreatedAt: w.CreatedAt.UTC(),
			UpdatedAt: w.UpdatedAt.UTC(),
			State:     kafka.Withdrawal_State(w.State).String(),
		},
		Relationships: resources.WithdrawalRelationships{
			Deposit: resources.Relation{
				Data: &resources.Key{
					ID:   w.DepositID,
					Type: resources.DEPOSITS,
				},
			},
		},
	}

	return result
}
