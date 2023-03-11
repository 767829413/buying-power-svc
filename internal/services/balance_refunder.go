package services

import (
	"context"
	"gitlab.com/eAuction/buying-power-svc/internal/services/deposits"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/events/go/kafka"

	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

//RunBalanceRefund - returns deposits for lost lots to the balance
func RunBalanceRefund(ctx context.Context, cfg config.Config) {
	r := &balanceRefund{
		log: cfg.Log().WithField("service", "balance_refund"),
		cfg: cfg,
		//depositStates: ,
		returnToCard: cfg.Deposits().ReturnDepositsToCard,
	}

	running.WithBackOff(ctx, r.log, "balance_refund", r.runOnce, time.Minute, time.Second*10, time.Minute)
}

type balanceRefund struct {
	log          *logan.Entry
	cfg          config.Config
	returnToCard bool
}

func (c *balanceRefund) runOnce(ctx context.Context) error {
	storage := c.cfg.Storage()
	participants, err := storage.Participants().Select(data.ParticipantsSelector{
		State:         []kafka.Participant_State{kafka.Participant_STATE_REJECTED, kafka.Participant_STATE_LOST},
		DepositStates: []int32{int32(kafka.DepositV2_STATE_LOCKED), int32(kafka.DepositV2_STATE_PAID)},
	})

	for _, participant := range participants {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		txStorage := storage.Clone()
		err = running.RunSafely(ctx, "", func(ctx context.Context) error {
			return txStorage.Transaction(func() error {
				return deposits.ReturnDepositForParticipant(txStorage, participant, c.cfg.Deposits())
			})
		})
		if err != nil {
			c.log.WithError(err).WithFields(logan.F{
				"lot_id":  participant.LotID,
				"address": participant.AccountAddress,
			}).Error("failed to return deposit to balance")
			continue
		}

	}

	return nil
}
