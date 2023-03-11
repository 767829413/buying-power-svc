package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/eAuction/bouncer/allow"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/events/go/kafka"
)

type depositListRequest struct {
	Page pgdb.OffsetPageParams

	FilterDepositor *string  `filter:"depositor"`
	RawState        []string `filter:"state"`
	Lot             []string `filter:"lot"`

	IncludeWithdrawals bool `include:"withdrawals"`

	FilterState []int32
}

func newDepositListRequest(r *http.Request) (*depositListRequest, error) {
	var request depositListRequest
	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to decode url")
	}

	for _, state := range request.RawState {
		v, ok := kafka.DepositV2_State_value[state]
		if !ok {
			return nil, validation.Errors{
				"filter[state]": errors.New("invalid value"),
			}
		}

		request.FilterState = append(request.FilterState, v)
	}

	return &request, validation.Errors{
		"filter[depositor]": validation.Validate(request.FilterDepositor, validation.Required),
	}.Filter()
}

func ListDeposits(w http.ResponseWriter, r *http.Request) {
	request, err := newDepositListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, err = Check(r, allow.Account{Address: *request.FilterDepositor})
	if err != nil {
		ape.RenderErr(w, problems.NotAllowed(err))
		return
	}

	deposits, err := Storage(r).Deposits().Select(data.DepositsSelector{
		States:     request.FilterState,
		Depositors: []string{*request.FilterDepositor},
		Lots:       request.Lot,
		Page:       &request.Page})
	if err != nil {
		panic(errors.Wrap(err, "failed to select deposits"))
	}

	response := resources.DepositListResponse{
		Data: make([]resources.Deposit, 0, len(deposits)),
	}

	withdrawals := map[string][]data.Withdrawal{}
	if request.IncludeWithdrawals {
		withdrawals, err = includeWithdrawals(r, deposits, &response.Included)
		if err != nil {
			panic(errors.Wrap(err, "failed to include withdrawals"))
		}
	}

	for _, deposit := range deposits {
		response.Data = append(response.Data, responses.NewDeposit(deposit, withdrawals[deposit.ID]))
	}

	populateDepositsLinks(r, request, &response)

	json.NewEncoder(w).Encode(&response)
}

func includeWithdrawals(r *http.Request, deposits []data.Deposit, includes *resources.Included) (map[string][]data.Withdrawal, error) {
	if len(deposits) == 0 {
		return map[string][]data.Withdrawal{}, nil
	}

	var depositsIDs []string
	for _, d := range deposits {
		depositsIDs = append(depositsIDs, d.ID)
	}

	ws, err := Storage(r).Withdrawals().Select(data.WithdrawalsSelector{
		Deposits: depositsIDs,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to select withdrawals for deposits")
	}

	result := map[string][]data.Withdrawal{}
	for _, w := range ws {
		wR := responses.NewWithdrawal(w)
		includes.Add(&wR)
		result[w.DepositID] = append(result[w.DepositID], w)
	}

	return result, nil
}

func populateDepositsLinks(r *http.Request, request *depositListRequest, response *resources.DepositListResponse) {
	response.Links = &resources.Links{}
	response.Links.Self = fmt.Sprintf("%s?%s", r.URL.Path, mustUrlValEncode(request))
	request.Page.PageNumber++
	response.Links.Next = fmt.Sprintf("%s?%s", r.URL.Path, mustUrlValEncode(request))
}

func mustUrlValEncode(src interface{}) string {
	res, err := urlval.Encode(src)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode request"))
	}

	return res
}
