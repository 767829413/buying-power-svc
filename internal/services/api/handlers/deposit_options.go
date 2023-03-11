package handlers

import (
	"net/http"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

func newDepositOptionsRequest(r *http.Request) (*DepositOptionsRequest, error) {
	var request DepositOptionsRequest
	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to decode url")
	}

	err := request.Validate()
	if err != nil {
		return nil, err
	}

	return &request, nil
}

type depositOptions struct {
	r *http.Request
}

//GetAttachedDepositsWithPending - only applicable for lots
func (d depositOptions) GetAttachedDepositsWithPending(_ string) ([]data.Deposit, error) {
	return nil, nil
}

//GetAvailableDeposits - only applicable for lots
func (d depositOptions) GetAvailableDeposits(_ string) ([]data.Deposit, error) {
	return nil, nil
}

//IsLimitSufficient - only applicable for lots
func (d depositOptions) IsLimitSufficient(_ string, _ config.Limit) (bool, *jsonapi.ErrorObject) {
	return true, nil
}

//BuyingPowerCurrency - returns currency in which we should return buying power
func (d depositOptions) BuyingPowerCurrency() string {
	return Config(d.r).Deposits().Currency
}

func (d depositOptions) GetDepositLimits(depositor string) ([]config.Limit, error) {
	return BuyingPowerCalculator(d.r).GetDepositLimits(depositor, nil, 0, "USD")
}

//ListDepositOptions - renders default deposit options without lot context
func ListDepositOptions(w http.ResponseWriter, r *http.Request) {
	request, err := newDepositOptionsRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	listDepositOptions(w, r, *request, depositOptions{r: r})
}
