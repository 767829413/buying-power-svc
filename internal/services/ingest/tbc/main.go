package tbc

import (
	"context"
	"strings"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/deposits"
	"gitlab.com/eAuction/events/go/kafka"
)

//Run - runs ingest for tbc transactions
func Run(ctx context.Context, cfg config.Config) {
	ingest := kafka.NewIngest(&handler{
		log:     cfg.Log(),
		storage: cfg.Storage().Clone(),
	}, cfg.TBCTransactions())

	ingest.Run(ctx)
}

type handler struct {
	log     *logan.Entry
	storage data.Storage
}

// HandleRequest handles kafka messages by passing them to correct event handlers.
func (h *handler) Handle(ctx context.Context, messages []kafka.Message) (err error) {
	err = h.storage.Transaction(func() error {
		for _, raw := range messages {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			msg := raw.MustTbcTransactionV2Message()
			err = h.handleMessage(msg)
			if err != nil {
				if errors.Cause(err) == errUnexpectedState {
					h.log.WithError(err).WithField("offset", msg.Offset()).Error("skipping event due to unexpected state")
					continue
				}

				return errors.Wrap(err, "failed to handle tbc tx")
			}
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "batch handle failed")
	}

	return nil
}

var errUnexpectedState = errors.New("state is unexpected in some way")

func (h *handler) handleMessage(msg kafka.TbcTransactionV2Message) error {
	tx := msg.Value
	if tx == nil {
		return errors.Wrap(errUnexpectedState, "expected TbcTransactionV2Message to have non nil value")
	}

	if tx.Purpose != kafka.TbcTransactionV2Purpose_TBC_TRANSACTION_PURPOSE_DEPOSIT ||
		!strings.HasPrefix(tx.InvoiceId, data.DepositIDPrefix) {
		return nil
	}

	fields := logan.F{
		"tx_id":      tx.Id,
		"invoice_id": tx.InvoiceId,
	}

	deposit, err := h.storage.Deposits().GetForUpdate(tx.InvoiceId)
	if err != nil {
		return errors.Wrap(err, "failed to get deposit by id", fields)
	}

	if deposit == nil {
		return errors.Wrap(errUnexpectedState, "failed to find deposit")
	}

	if deposit.Currency != tx.Currency || deposit.Amount != tx.Amount || deposit.Depositor != tx.AccountId {
		return errors.Wrap(errUnexpectedState, "expected deposit to fully match tx", fields.Merge(logan.F{
			"deposit.currency":  deposit.Currency,
			"tx.currency":       tx.Currency,
			"deposit.amount":    deposit.Amount,
			"tx.amount":         tx.Amount,
			"deposit.depositor": deposit.Depositor,
			"tx.accountID":      tx.AccountId,
		}))
	}

	switch tx.Status {
	// initial status
	case kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_CREATED,
		// transaction is beeing processed
		kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_PENDING:
		return nil
	// money were charged from preauth transactions
	case kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_CONFIRMED:
		return errors.Wrap(h.storage.Deposits().SetState(tx.InvoiceId, int32(kafka.DepositV2_STATE_FEE_CHARGED)),
			"failed to set fee charged for manual charged deposit", fields)
	case // the transaction was refunded
		kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_REFUNDED:
		return errors.From(h.handleRefunded(*deposit), fields)
	case // perauth was reversed
		kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_REVERSED:
		return errors.From(h.handleReversed(*deposit, *tx), fields)
		// successful preauth for DMS, or completed transaction for other types
	case kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_OK:
		return errors.From(h.handleConfirmed(*deposit), fields)
		// transaction has failed
	case kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_FAILED,
		// transaction was declined by the bank
		kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_DECLINED,
		// user hasn't completed the payment process
		kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_TIMEOUT:
		return errors.From(h.handleFailed(*deposit, *tx), fields)
		// expired preauth was autoreversed
	case kafka.TbcTransactionV2Status_TBC_TRANSACTION_STATUS_AUTOREVERSED:
		return errors.Wrap(errUnexpectedState, "autoreverse should never happen", fields)
	default:
		return errors.Wrap(errUnexpectedState, "received tx with unexpected status", fields.Merge(logan.F{
			"status": tx.Status.String(),
		}))
	}
}

func (h *handler) handleReversed(deposit data.Deposit, tx kafka.TbcTransactionV2) error {
	allowedStates := map[kafka.DepositV2_State]struct{}{
		kafka.DepositV2_STATE_REVERSED: {},
		kafka.DepositV2_STATE_PAID:     {},
		kafka.DepositV2_STATE_RETURNED: {},
	}
	if _, ok := allowedStates[kafka.DepositV2_State(deposit.State)]; !ok {
		return errors.Wrap(errUnexpectedState, "expected reverse to only occur for paid deposits", logan.F{
			"deposit.state": deposit.State,
		})
	}

	// TODO: do something if lot is attached to deposit
	return deposits.Return(h.storage, deposit)
}

func (h *handler) handleFailed(deposit data.Deposit, tx kafka.TbcTransactionV2) error {
	if deposit.IsArtificial {
		return nil
	}

	if deposit.State != int32(kafka.DepositV2_STATE_PENDING) && deposit.State != int32(kafka.DepositV2_STATE_FAILED) {
		return errors.Wrap(errUnexpectedState, "expected to have deposit in state pending for tx with deposit purpose", logan.F{
			"state": deposit.State,
		})
	}

	err := h.storage.Deposits().SetState(tx.InvoiceId, int32(kafka.DepositV2_STATE_FAILED))
	if err != nil {
		return errors.Wrap(err, "failed to set failed state for deposit")
	}

	return nil
}

var allowedConfirmedTransition = map[kafka.DepositV2_State]struct{}{
	kafka.DepositV2_STATE_PENDING: {},
	kafka.DepositV2_STATE_PAID:    {},
	kafka.DepositV2_STATE_FAILED:  {},
}

func (h *handler) handleConfirmed(deposit data.Deposit) error {
	_, allowed := allowedConfirmedTransition[kafka.DepositV2_State(deposit.State)]
	if !allowed {
		return errors.Wrap(errUnexpectedState, "expected to have deposit in state pending for confirmed tx", logan.F{
			"state": deposit.State,
		})
	}

	return deposits.Approve(h.storage, deposit)
}

func (h *handler) handleRefunded(deposit data.Deposit) error {
	if deposit.State != int32(kafka.DepositV2_STATE_PAID) && deposit.State != int32(kafka.DepositV2_STATE_RETURNED) {
		return errors.Wrap(errUnexpectedState, "expected to have deposit in state paid or returned for refunded tx", logan.F{
			"state": deposit.State,
		})
	}

	return deposits.Return(h.storage, deposit)
}
