package participants

import (
	"context"
	"encoding/json"
	"github.com/lib/pq"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	router "gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/router"
	"gitlab.com/eAuction/events/go/kafka"
	"time"
)

//Router - routes participants' updates to dedicated connections
type Router interface {
	Subscribe(ctx context.Context, connectionID, accountID string) <-chan []byte
}

type notifier struct {
	log              *logan.Entry
	ctx              context.Context
	storage          data.Storage
	getNewPqListener func() *pq.Listener

	router            router.Router
	lastUpdatedAt     time.Time
	participantStates []kafka.Participant_State
}

//NewRouter - returns new notifier
func NewRouter(ctx context.Context, log *logan.Entry, storage data.Storage, getNewPqListener func() *pq.Listener) Router {
	log = log.WithField("service", "participations_router")
	r := &notifier{
		log:               log,
		storage:           storage,
		router:            router.New(ctx, log),
		ctx:               ctx,
		lastUpdatedAt:     time.Now().UTC(),
		getNewPqListener:  getNewPqListener,
		participantStates: []kafka.Participant_State{kafka.Participant_STATE_DEPOSIT_PENDING, kafka.Participant_STATE_PENDING},
	}

	go running.WithBackOff(r.ctx, r.log, "participations_router", r.run, time.Second, time.Second*10, time.Second*30)

	return r
}

func (r *notifier) run(ctx context.Context) error {
	pqListener := r.getNewPqListener()
	err := pqListener.Listen("participant_updated")
	if err != nil {
		return errors.Wrap(err, "failed to listen to participant_updated")
	}

	r.log.Info("listening")
	defer pqListener.Close()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case _, ok := <-pqListener.NotificationChannel():
			if !ok {
				return errors.New("notification channel has been closed")
			}
		}

		err = r.handleParticipantUpdate()
		if err != nil {
			return errors.Wrap(err, "failed to handle participant update")
		}

	}
}

func (r *notifier) handleParticipantUpdate() error {
	lastUpdatedAt, participantsByAcc, lots, err := r.getParticipantsWithLots(data.ParticipantsSelector{
		UpdatedAfter: &r.lastUpdatedAt,
	})
	if err != nil {
		return errors.Wrap(err, "failed get participants with lots")
	}

	if lastUpdatedAt != nil {
		r.lastUpdatedAt = *lastUpdatedAt
	}

	for accountID, participants := range participantsByAcc {
		env := responses.PopulateEnvelopWithParticipations(participants, lots)
		var msg []byte
		msg, err = json.Marshal(env)
		if err != nil {
			panic(errors.Wrap(err, "failed to marshal envelope"))
		}
		r.router.NotifyResourceListeners(accountID, [][]byte{msg})
	}

	return nil
}

func (r *notifier) Subscribe(ctx context.Context, connectionID, accountID string) <-chan []byte {
	q := r.router.AddNewListener(ctx, connectionID)
	r.router.Subscribe(connectionID, accountID)
	go func() {
		err := r.notifyWithInitial(connectionID, accountID)
		if err != nil {
			r.log.WithError(err).WithField("connection_id", connectionID).Error("failed to notify with initial - removing")
			r.router.CloseListener(connectionID)
		}
	}()

	return q
}

func (r *notifier) notifyWithInitial(connectionID, accountID string) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.Wrap(errors.WithStack(errors.FromPanic(rec)), "notify with initial panic")
		}
	}()
	var participants map[string][]data.Participant
	var lots map[string]data.Lot
	_, participants, lots, err = r.getParticipantsWithLots(data.ParticipantsSelector{
		State:     r.participantStates,
		AccountID: &accountID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get initial participants")
	}

	resp := responses.PopulateEnvelopWithParticipations(participants[accountID], lots)
	var msg []byte
	msg, err = json.Marshal(resp)
	if err != nil {
		return errors.Wrap(err, "failed to marshal participants env")
	}
	r.router.Notify(connectionID, [][]byte{msg})
	return nil
}

func (r *notifier) getParticipantsWithLots(selector data.ParticipantsSelector) (*time.Time, map[string][]data.Participant,
	map[string]data.Lot, error) {

	rawParticipants, err := r.storage.Participants().Select(selector)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to get participants")
	}

	lotIDs := make([]string, len(rawParticipants))
	participants := map[string][]data.Participant{}
	var lastUpdatedAt *time.Time
	for i, p := range rawParticipants {
		lotIDs[i] = p.LotID
		participants[p.AccountAddress] = append(participants[p.AccountAddress], p)
		if lastUpdatedAt == nil || lastUpdatedAt.Before(p.UpdatedAt) {
			lastUpdatedAt = new(time.Time)
			*lastUpdatedAt = p.UpdatedAt
		}
	}
	lots, err := r.storage.Lots().ByIDs(lotIDs)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to get lots")
	}

	return lastUpdatedAt, participants, lots, nil
}