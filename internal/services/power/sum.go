package power

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

//GetTotalDepositAmount - returns sum of deposit amounts in the deposit currency
func GetTotalDepositAmount(deposits []data.Deposit, currency string) (int64, error) {
	return getTotalDepositAmount(deposits ,currency, false)
}

func getTotalDepositAmount(deposits []data.Deposit, currency string, skipArtificial bool) (int64, error) {
	var totalDeposit int64
	for _, deposit := range deposits {
		if skipArtificial && deposit.IsArtificial {
			continue
		}
		if deposit.Currency != currency {
			return 0, errors.From(errors.New("expected all deposits to be in config currency"), logan.F{
				"deposit.id":       deposit.ID,
				"deposit.currency": deposit.Currency,
				"cfg.currency":     currency,
			})
		}

		if deposit.Amount+totalDeposit < totalDeposit {
			return 0, errors.From(errors.New("totalDeposit sum overflow"), logan.F{
				"depositor": deposit.Depositor,
				"lot_id":    deposit.LotID,
			})
		}

		totalDeposit += deposit.Amount
	}

	return totalDeposit, nil
}

//MustTotalDepositAmount - GetTotalDepositAmount, but panics on error
func MustTotalDepositAmount(deposits []data.Deposit, currency string) int64 {
	result, err := GetTotalDepositAmount(deposits, currency)
	if err != nil {
		panic(err)
	}

	return result
}

