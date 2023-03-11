package services

import (
	"context"
	"strconv"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

type participantsSyncer struct {
	cfg    config.Config
	log    *logan.Entry
	writer kafka.AuctionEventWriter
}

//RunParticipantsSyncer - sends participants into kafka
func RunParticipantsSyncer(ctx context.Context, cfg config.Config) {
	s := participantsSyncer{
		cfg:    cfg,
		writer: cfg.AuctionsWriter(),
	}
	running.WithBackOff(ctx, cfg.Log(), "participants_syncer", s.runOnce, time.Second, time.Second, time.Minute)
}

func (s *participantsSyncer) runOnce(ctx context.Context) error {
	listener := s.cfg.NewListener()
	err := listener.Listen("participant_updated")
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to participant_updated")
	}

	s.cfg.Log().Info("listening")
	defer listener.Close()
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	storage := s.cfg.Storage()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case _, ok := <-listener.NotificationChannel():
			if !ok {
				return errors.New("notification channel has been closed")
			}
		case _, ok := <-ticker.C:
			if !ok {
				return errors.New("ticker stopped")
			}
		}

		err = s.handleNotSynced(ctx, storage)
		if err != nil {
			return errors.Wrap(err, "failed to handle not sycned")
		}
	}
}

func (s *participantsSyncer) handleNotSynced(ctx context.Context, storage data.Storage) error {
	participants, err := storage.Participants().SelectNotSynced()
	if err != nil {
		return errors.Wrap(err, "failed to select not synced participants")
	}

	for _, p := range participants {
		var event kafka.AuctionEvent
		event, err = s.participantToEvent(p)
		if err != nil {
			return errors.Wrap(err, "failed to convert participant to event")
		}

		err = s.writer.WriteMessages(ctx, event.Message())
		if err != nil {
			return errors.Wrap(err, "failed to write message to kafka")
		}

		err = storage.Participants().MarkSynced(p)
		if err != nil {
			return errors.Wrap(err, "failed to mark participant as synced")
		}
	}

	return nil
}

func (s *participantsSyncer) participantToEvent(p data.Participant) (kafka.AuctionEvent, error) {
	eventParticipant := kafka.Participant{
		AccountId:       p.AccountAddress,
		RequestedBuyNow: p.RequestedBuyNow,
		CurrentBid:      p.CurrentBid,
		State:           kafka.Participant_State(p.State),
		BidLimit:        p.BidLimit,
		AutoBidLimit:    p.AutoBidLimit,
		OutbidedWith:    p.OutbidedWith.Int64,
	}

	lotID, err := strconv.ParseInt(p.LotID, 10, 64)
	if err != nil {
		return kafka.AuctionEvent{}, errors.Wrap(err, "failed to parse lotID to int")
	}

	if p.PublishedCreation.Valid && p.PublishedCreation.Bool {
		return kafka.AuctionEvent{
			Type: kafka.AuctionEvent_TYPE_PARTICIPANT_UPDATED,
			ParticipantUpdated: &kafka.ParticipantUpdated{
				LotId:       lotID,
				Participant: &eventParticipant,
			},
		}, nil
	}

	return kafka.AuctionEvent{
		Type: kafka.AuctionEvent_TYPE_PARTICIPANT_CREATED,
		ParticipantCreated: &kafka.ParticipantCreated{
			LotId:       lotID,
			Participant: &eventParticipant,
		},
	}, nil
}
