package handlers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/eAuction/buying-power-svc/internal/services/power"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/go/amount"

	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/buying-power-svc/internal/types"
)

type depositOptionsHandler interface {
	GetAttachedDepositsWithPending(depositor string) ([]data.Deposit, error)
	GetAvailableDeposits(depositor string) ([]data.Deposit, error)
	IsLimitSufficient(depositor string, limit config.Limit) (bool, *jsonapi.ErrorObject)
	BuyingPowerCurrency() string
	GetDepositLimits(depositor string) ([]config.Limit, error)
}

//DepositOptionsRequest - base request. Must be public to allow to embed into other struct
type DepositOptionsRequest struct {
	Depositor *string `filter:"depositor"`

	IncludeDepositsToAttach bool `include:"deposits_to_attach"`
	IncludeDepositToCreate  bool `include:"deposit_to_create"`
}

//Validate - returns error if request is not valid
func (request DepositOptionsRequest) Validate() error {
	return validation.Errors{
		"filter[depositor]": validation.Validate(request.Depositor, validation.Required),
	}.Filter()
}

func listDepositOptions(w http.ResponseWriter, r *http.Request, request DepositOptionsRequest, handler depositOptionsHandler) {
	identity, _ := GetAuthorizedAccount(r, w, *request.Depositor, false)
	if identity == nil {
		return
	}

	attachedDeposits, err := handler.GetAttachedDepositsWithPending(*request.Depositor)
	if err != nil {
		panic(errors.Wrap(err, "failed to select deposit attached to lot"))
	}

	availableDeposits, err := handler.GetAvailableDeposits(*request.Depositor)
	if err != nil {
		panic(errors.Wrap(err, "failed to get available deposits"))
	}

	var response resources.DepositOptionListResponse

	balance, err := getPotentialBalance(r, *request.Depositor)
	if err != nil {
		panic(errors.Wrap(err, "failed to get potential balance"))
	}

	response.Data = []resources.DepositOption{}

	deposits := Config(r).Deposits()
	depositLimits, err := handler.GetDepositLimits(*request.Depositor)
	if err != nil {
		if errors.Cause(err) == power.ErrNotAvailable {
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"filter[depositor]": err,
			})...)
			return
		}
		panic(errors.Wrap(err, "failed to generate limit for direct sale lot"))
	}
	for _, limit := range depositLimits {
		state, depositsToAttach, depositToCreate := BuyingPowerCalculator(r).GetDepositsOptions(*request.Depositor, balance, attachedDeposits, availableDeposits, limit)
		if state == types.DepositOptionStateAvailable {
			isSufficient, respErr := handler.IsLimitSufficient(*request.Depositor, limit)
			if respErr != nil {
				ape.RenderErr(w, respErr)
			}

			if !isSufficient {
				continue
			}
		}

		var buyingPower int64
		buyingPower, err = BuyingPowerCalculator(r).BuyingPower(*request.Depositor, limit.Deposit, deposits.Currency, handler.BuyingPowerCurrency())
		if err != nil {
			panic(errors.Wrap(err, "failed to calculate buying power"))
		}

		option := resources.DepositOption{
			Key: resources.Key{
				ID:   lorem.ULID(),
				Type: resources.DEPOSIT_OPTIONS,
			},
			Attributes: resources.DepositOptionAttributes{
				Amount:              amount.Fiat(limit.Deposit),
				AmountCurrency:      deposits.Currency,
				BuyingPower:         amount.Fiat(buyingPower),
				BuyingPowerCurrency: handler.BuyingPowerCurrency(),
				State:               resources.DepositOptionState(state.String()),
			},
			Relationships: resources.DepositOptionRelationships{
				DepositsToAttach: &resources.RelationCollection{},
			},
		}

		for _, deposit := range depositsToAttach {
			depositResource := responses.NewDeposit(deposit, nil)
			option.Relationships.DepositsToAttach.Data = append(option.Relationships.DepositsToAttach.Data, depositResource.Key)
			if request.IncludeDepositsToAttach {
				response.Included.Add(&depositResource)
			}
		}

		if depositToCreate != nil {
			depositResource := responses.NewDeposit(*depositToCreate, nil)
			option.Relationships.DepositToCreate = depositResource.AsRelation()
			if request.IncludeDepositToCreate {
				response.Included.Add(&depositResource)
			}
		}

		response.Data = append(response.Data, option)
	}

	json.NewEncoder(w).Encode(&response)
}
