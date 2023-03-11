package handlers

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/power"
	"gitlab.com/eAuction/events/go/kafka"
)

type attachDepositsV2 struct {
	lotID     string
	deposits  map[string]struct{}
	depositor string
	isAdmin bool
}

func AttachDepositsV2(r *http.Request, request attachDepositsV2) error {
	var lot *data.Lot
	lot, err := Storage(r).Lots().ByID(request.lotID)
	if err != nil {
		return errors.Wrap(err, "failed to load lot by id")
	}

	if lot == nil {
		return problems.NotFound()
	}

	if !request.isAdmin && lot.State != int32(kafka.Lot_STATE_PENDING) {
		return problems.BadRequest(validation.Errors{
			"lot": errors.New("lot must be in state pending"),
		}.Filter())[0]
	}

	participant, err := Storage(r).Participants().ByLotIDAccountAddress(lot.ID, request.depositor)
	if err != nil {
		return errors.Wrap(err, "failed to select participant")
	}

	if participant == nil ||
		!(participant.State == int32(kafka.Participant_STATE_PENDING) || participant.State == int32(kafka.Participant_STATE_DEPOSIT_PENDING)) {
		return problems.BadRequest(validation.Errors{
			"lot": errors.New("participation for the lot must exist in state pending or deposit pending"),
		}.Filter())[0]
	}

	err = detachDeposits(r, request)
	if err != nil {
		return errors.Wrap(err, "failed to detach deposits")
	}

	for depositID := range request.deposits {
		err = attachDeposit(r, *lot, depositID, request.depositor)
		if err != nil {
			return errors.Wrap(err, "failed to attach deposit")
		}
	}

	ok, err := isDepositAmountSufficient(r, *lot, request.depositor)
	if err != nil {
		return errors.Wrap(err, "failed to check if deposit amount is sufficient")
	}

	if !ok {
		return problems.BadRequest(validation.Errors{
			"lot": errors.New("deposit amount is not sufficient to pay fees"),
		})[0]
	}

	return nil
}

func detachDeposits(r *http.Request, request attachDepositsV2) error {
	attachedDeposits, err := getAttachedDepositsWithPending(r, request.lotID, request.depositor, true)
	if err != nil {
		return errors.Wrap(err, "failed to get attached deposits")
	}

	lockedStates := map[int32]struct{}{
		int32(kafka.DepositV2_STATE_PAID):   {},
		int32(kafka.DepositV2_STATE_LOCKED): {},
	}
	for _, attachedDeposit := range attachedDeposits {
		if _, ok := request.deposits[attachedDeposit.ID]; ok {
			continue
		}

		if _, ok := lockedStates[attachedDeposit.State]; ok {
			return problems.BadRequest(validation.Errors{
				"data": errors.New("not allowed to detach deposits in locked state"),
			})[0]
		}

		err = Storage(r).Deposits().SetLot(attachedDeposit.ID, nil)
		if err != nil {
			return errors.Wrap(err, "failed to detach lot")
		}
	}

	return nil
}

func attachDeposit(r *http.Request, lot data.Lot, depositID, depositor string) error {
	deposit, err := Storage(r).Deposits().GetForUpdate(depositID)
	if err != nil {
		return errors.Wrap(err, "failed to get deposit for update")
	}

	depositKey := fmt.Sprintf("data[%s]", depositID)
	if deposit == nil || deposit.Depositor != depositor {
		return problems.BadRequest(validation.Errors{
			depositKey: errors.New("not found"),
		})[0]
	}

	if deposit.LotID != nil && *deposit.LotID != lot.ID {
		return problems.BadRequest(validation.Errors{
			depositKey: errors.New("not allowed to reattach lot. Detach it explicitly first"),
		})[0]
	}

	if deposit.LotID != nil && *deposit.LotID == lot.ID {
		return nil
	}

	if deposit.State != int32(kafka.DepositV2_STATE_PENDING) && deposit.State != int32(kafka.DepositV2_STATE_PAID) {
		return problems.BadRequest(validation.Errors{
			depositKey: errors.New("deposit state does not allow to perform reattachment"),
		}.Filter())[0]
	}

	maxLotEndTime := deposit.CreatedAt.Add(Config(r).Deposits().ExpiresIn).Add(-Config(r).Deposits().LotProcessingDuration)
	if lot.EndsAt.After(maxLotEndTime) {
		return problems.BadRequest(validation.Errors{
			depositKey: errors.New("deposit expires before lot can be fully processed"),
		}.Filter())[0]
	}

	Log(r).WithFields(logan.F{
		"depositor":  deposit.Depositor,
		"deposit_id": depositID,
		"lot_id":     lot.ID,
	}).Info("attaching deposit to lot")
	err = Storage(r).Deposits().SetLot(depositID, &lot.ID)
	if err != nil {
		return errors.Wrap(err, "failed to set lot for deposit")
	}

	return nil
}

func isDepositAmountSufficient(r *http.Request, lot data.Lot, depositorID string) (bool, error) {
	attachedDeposits, err := getAttachedDepositsWithPending(r, lot.ID, depositorID, false)
	if err != nil {
		return false, errors.Wrap(err, "failed to get attached deposits")
	}

	totalDepositAmount := power.MustTotalDepositAmount(attachedDeposits, Config(r).Deposits().Currency)
	ok, err := BuyingPowerCalculator(r).IsDepositSufficientToPayFee(depositorID, lot, totalDepositAmount, Config(r).Deposits().Currency)
	if err != nil {
		if errors.Cause(err) == power.ErrNotAvailable {
			return false, nil
		}

		return false, errors.Wrap(err, "failed to check if deposit amount is sufficient")
	}

	return ok, nil
}
