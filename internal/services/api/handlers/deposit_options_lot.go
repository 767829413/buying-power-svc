package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/power"
	"gitlab.com/eAuction/events/go/kafka"
	"gitlab.com/eAuction/go/amount"
)

type depositOptionsLotRequest struct {
	DepositOptionsRequest
	LotID                 string
	RawIsBuyNow           *bool        `filter:"is_buy_now"`
	RawCustomerPrice      *amount.Fiat `filter:"customer_price"`
	CustomerPriceCurrency *string      `filter:"customer_price_currency"`

	Lot *data.Lot
}

func (r depositOptionsLotRequest) IsBuyNow() bool {
	return r.RawIsBuyNow != nil && *r.RawIsBuyNow
}

func (r depositOptionsLotRequest) CustomerPrice() int64 {
	if r.RawCustomerPrice == nil {
		return 0
	}

	return int64(*r.RawCustomerPrice)
}

func newDepositOptionsLotRequest(r *http.Request) (*depositOptionsLotRequest, error) {
	request := depositOptionsLotRequest{
		LotID: chi.URLParam(r, "lot"),
	}
	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to decode url")
	}

	err := request.DepositOptionsRequest.Validate()
	if err != nil {
		return nil, err
	}

	err = validation.Errors{
		"filter[lot]":                     validation.Validate(request.LotID, validation.Required),
		"filter[customer_price_currency]": validation.Validate(request.CustomerPriceCurrency, validation.Required),
	}.Filter()
	if err != nil {
		return nil, err
	}

	request.Lot, err = Storage(r).Lots().ByID(request.LotID)
	if err != nil {
		panic(errors.Wrap(err, "failed to load lot by id"))
	}

	if request.Lot == nil {
		return nil, validation.Errors{
			"lot": errors.New("not found"),
		}
	}

	if request.IsBuyNow() && (request.Lot.Details.AuctionPrices == nil || request.Lot.Details.AuctionPrices.BuyNow <= 0) {
		return nil, validation.Errors{
			"lot": errors.New("lot must be auction supporting buy now to request options for buy now"),
		}
	}

	return &request, nil
}

type depositOptionsLot struct {
	r       *http.Request
	request depositOptionsLotRequest
}

//GetAttachedDepositsWithPending - only applicable for lots
func (d depositOptionsLot) GetAttachedDepositsWithPending(depositor string) ([]data.Deposit, error) {
	return getAttachedDepositsWithPending(d.r, d.request.LotID, depositor, false)
}

//GetAvailableDeposits - only applicable for lots
func (d depositOptionsLot) GetAvailableDeposits(depositor string) ([]data.Deposit, error) {
	no := false
	deposits := Config(d.r).Deposits()
	depositMinCreatedAt := d.request.Lot.EndsAt.Add(-deposits.ExpiresIn).Add(deposits.LotProcessingDuration).Add(time.Minute).UTC()
	return Storage(d.r).Deposits().Select(data.DepositsSelector{
		States:       []int32{int32(kafka.DepositV2_STATE_PAID)},
		HasLot:       &no,
		Depositors:   []string{depositor},
		MinCreatedAt: &depositMinCreatedAt,
		Currency:     &deposits.Currency,
	})
}

//IsLimitSufficient - only applicable for lots
func (d depositOptionsLot) IsLimitSufficient(depositor string, limit config.Limit) (bool, *jsonapi.ErrorObject) {
	if d.request.IsBuyNow() && !limit.IsBuyNowAllowed {
		return false, nil
	}

	isSufficient, err := BuyingPowerCalculator(d.r).IsDepositSufficientToPayFee(depositor, *d.request.Lot, limit.Deposit,
		Config(d.r).Deposits().Currency)

	if err != nil {
		if errors.Cause(err) == power.ErrNotAvailable {
			return false, problems.BadRequest(validation.Errors{
				"lot": errors.New("fees are not available for specified lot"),
			})[0]
		}

		panic(errors.Wrap(err, "failed to check if fee is deposit sufficient to pay the fee"))
	}

	if !isSufficient {
		return false, nil
	}

	customerPrice, _, err := Config(d.r).ExRatesConverter().ConvertCents(*d.request.CustomerPriceCurrency, d.request.Lot.Currency, d.request.CustomerPrice())
	if err != nil {
		panic(errors.Wrap(err, "failed to convert customer price"))
	}

	isSufficient, err = BuyingPowerCalculator(d.r).IsLimitSufficientForLot(*d.request.Lot, limit.Limit,
		Config(d.r).Deposits().Currency, d.request.IsBuyNow(), customerPrice)
	if err != nil {
		panic(errors.Wrap(err, "failed to check if limit is sufficient for lot"))
	}

	if !isSufficient {
		return false, nil
	}

	return true, nil
}

func (d depositOptionsLot) GetDepositLimits(depositor string) ([]config.Limit, error) {
	return BuyingPowerCalculator(d.r).GetDepositLimits(depositor, d.Lot(), d.request.CustomerPrice(), *d.request.CustomerPriceCurrency)
}

//BuyingPowerCurrency - returns currency in which we should return buying power
func (d depositOptionsLot) BuyingPowerCurrency() string {
	return d.request.Lot.Currency
}

func (d depositOptionsLot) Lot() *data.Lot {
	return d.request.Lot
}

//ListDepositOptionsLot - renders deposit options with lot context
func ListDepositOptionsLot(w http.ResponseWriter, r *http.Request) {
	request, err := newDepositOptionsLotRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	listDepositOptions(w, r, request.DepositOptionsRequest, depositOptionsLot{
		r:       r,
		request: *request,
	})
}

func getAttachedDepositsWithPending(r *http.Request, lotID string, depositor string, forUpdate bool) ([]data.Deposit, error) {
	deposits := Config(r).Deposits()
	return Storage(r).Deposits().Select(data.DepositsSelector{
		States:     []int32{int32(kafka.DepositV2_STATE_PAID), int32(kafka.DepositV2_STATE_LOCKED), int32(kafka.DepositV2_STATE_PENDING)},
		Lots:       []string{lotID},
		Depositors: []string{depositor},
		Currency:   &deposits.Currency,
		ForUpdate:  forUpdate,
	})
}
