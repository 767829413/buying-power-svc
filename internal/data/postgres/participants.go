package postgres

import (
	"database/sql"
	"time"

	"gitlab.com/eAuction/events/go/kafka"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

type participants struct {
	db *pgdb.DB
}

//Participants - returns new instance of participants
func (s *storage) Participants() data.Participants {
	return &participants{db: s.db}
}

//Create - creates new participant
func (q *participants) Create(identityID, lotID string) error {
	stmt := squirrel.Insert("participants").SetMap(map[string]interface{}{
		"account_address":   identityID,
		"lot_id":            lotID,
		"state":             int32(kafka.Participant_STATE_DEPOSIT_PENDING),
		"requested_buy_now": false,
		"current_bid":       0,
		"bid_limit":         0,
		"auto_bid_limit":    0,
		"is_local":          true,
		"is_synced":         false,
		"updated_at":        time.Now().UTC(),
		"requested_bid":     0,
	})
	return q.db.Exec(stmt)
}

// Store saves participant into database. If participant with such lot and account
// already exists, Store updates the existing one.
// NOTE: Only use for ingest
func (q *participants) Store(participant data.Participant) error {
	stmt := squirrel.Insert("participants").SetMap(map[string]interface{}{
		"account_address":   participant.AccountAddress,
		"lot_id":            participant.LotID,
		"state":             participant.State,
		"requested_buy_now": participant.RequestedBuyNow,
		"current_bid":       participant.CurrentBid,
		"bid_limit":         participant.BidLimit,
		"auto_bid_limit":    participant.AutoBidLimit,
		"is_local":          participant.IsLocal,
		"is_synced":         true,
		"updated_at":        time.Now().UTC(),
		"requested_bid":     participant.RequestedBid,
	}).Suffix(`
		on conflict on constraint participant_lot_account_unique
		do update
		set (
			account_address,
			lot_id,
			state,
			requested_buy_now,
			current_bid,
			bid_limit,
			auto_bid_limit,
			is_local,
			updated_at
		) = (
			excluded.account_address,
			excluded.lot_id,
			excluded.state,
			excluded.requested_buy_now,
			excluded.current_bid,
			excluded.bid_limit,
			excluded.auto_bid_limit,
			excluded.is_local,
			excluded.updated_at
		)
	`)

	return q.db.Exec(stmt)
}

// ByLotIDAccountAddress returns a participant by it's lot ID and account address.
// Returns nil, nil if entry doesn't exist.
func (q *participants) ByLotIDAccountAddress(lotID, accountAddress string) (*data.Participant, error) {
	return q.get(lotID, accountAddress, false)
}

func (q *participants) Exists(lotID, accountAddress string) (bool, error) {
	p, err := q.get(lotID, accountAddress, false)
	return p != nil, err
}

func (q *participants) get(lotID, accountAddress string, forUpdate bool) (*data.Participant, error) {
	stmt := squirrel.Select("*").From("participants").Where(squirrel.Eq{
		"lot_id":          lotID,
		"account_address": accountAddress,
	})
	if forUpdate {
		stmt = stmt.Suffix("for update")
	}

	var result data.Participant
	err := q.db.Get(&result, stmt)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to get participant by address")
	}

	return &result, nil
}

func (q *participants) GetForUpdate(lotID, accountAddress string) (*data.Participant, error) {
	return q.get(lotID, accountAddress, true)
}

//GetClosestForUpdate - returns pending participant with bid closest to the specified
func (q *participants) GetClosestForUpdate(lotID string, bid int64, platforms []string) (*data.Participant, error) {
	stmt := squirrel.Select("p.*").
		From("participants p").
		Join("identities i on p.account_address = i.id").
		Where(
			squirrel.Eq{
				"p.lot_id":   lotID,
				"p.state":    int32(kafka.Participant_STATE_PENDING),
				"i.platform": platforms,
			}).
		Where("p.bid_limit >= ? AND p.requested_bid >= ?", bid, bid).
		OrderBy("p.requested_bid asc").
		Limit(1).
		Suffix("for update")

	var result data.Participant
	err := q.db.Get(&result, stmt)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to get participant by closest bid")
	}

	return &result, nil
}

//Select - returns list of participants for specified selector
func (q *participants) Select(selector data.ParticipantsSelector) ([]data.Participant, error) {
	var s squirrel.SelectBuilder
	if len(selector.DepositStates) != 0 {
		s = squirrel.Select("distinct(p.*)")
	} else {
		s = squirrel.Select("p.*")
	}
	s = s.From("participants p")
	if len(selector.State) != 0 {
		s = s.Where(squirrel.Eq{
			"p.state": selector.State,
		})
	}

	if len(selector.DepositStates) != 0 {
		s = s.Join("deposits d on d.lot_id is not null and p.lot_id = d.lot_id AND d.depositor = p.account_address").Where(
			squirrel.Eq{
				"d.state": selector.DepositStates,
			})
	}

	if selector.LotID != nil {
		s = s.Where(squirrel.Eq{"p.lot_id": selector.LotID})
	}

	if selector.AccountID != nil {
		s = s.Where(squirrel.Eq{"p.account_address": *selector.AccountID})
	}

	if selector.UpdatedAfter != nil {
		s = s.Where(squirrel.GtOrEq{"p.updated_at": *selector.UpdatedAfter})
	}

	if selector.MinLotsEndsAt != nil {
		s = s.Join("lots l on l.id = p.lot_id").Where(squirrel.Lt{
			"l.ends_at": *selector.MinLotsEndsAt,
		})
	}

	if len(selector.Platforms) != 0 {
		s = s.Join("identities i on p.account_address = i.id").Where(squirrel.Eq{
			"i.platform": selector.Platforms,
		})
	}

	if selector.ForUpdate {
		s = s.Suffix("for update")
	}

	var result []data.Participant
	err := q.db.Select(&result, s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select participants")
	}

	return result, nil
}

//SelectNotSynced - returns slice of participants
func (q *participants) SelectNotSynced() ([]data.Participant, error) {
	var result []data.Participant
	err := q.db.Select(&result, squirrel.Select("*").From("participants").Where("NOT is_synced"))
	return result, err
}

//MarkSynced - marks participant as synced
func (q *participants) MarkSynced(p data.Participant) error {
	return q.db.Exec(squirrel.Update("participants").SetMap(map[string]interface{}{
		"is_synced":          true,
		"published_creation": true,
	}).Where(squirrel.Eq{
		"id":      p.ID,
		"version": p.Version,
	}))
}

//SetBidLimit - sets bid limit for participation
func (q *participants) SetBidLimit(id, bidLimit int64) error {
	return q.db.Exec(squirrel.Update("participants").SetMap(q.withDefaults(map[string]interface{}{
		"state":     int32(kafka.Participant_STATE_PENDING),
		"bid_limit": bidLimit,
	})).Where("id = ?", id))
}

//SetAutoBidLimit - sets auto bid limit for participation
func (q *participants) SetAutoBidLimit(id, bidLimit int64) error {
	return q.db.Exec(squirrel.Update("participants").SetMap(q.withDefaults(map[string]interface{}{
		"auto_bid_limit": bidLimit,
	})).Where("id = ?", id))
}

//SetBid - sets auto bid limit for participation
func (q *participants) SetBid(id, bid int64) error {
	return q.db.Exec(squirrel.Update("participants").SetMap(q.withDefaults(map[string]interface{}{
		"current_bid": bid,
	})).Where("id = ?", id))
}

//SetRequestedBid - sets requested bid for participant
func (q *participants) SetRequestedBid(id, bid int64) error {
	// we do not mark participant for update on purpose as requested_bid is not present in the participant event
	return q.db.Exec(squirrel.Update("participants").SetMap(map[string]interface{}{
		"requested_bid": bid,
		"updated_at":    time.Now().UTC(),
		"is_local":      true,
	}).Where("id = ?", id))
}

func (q *participants) withDefaults(newValues map[string]interface{}) map[string]interface{} {
	newValues["is_synced"] = false
	newValues["updated_at"] = time.Now().UTC()
	newValues["version"] = squirrel.Expr("version+1")
	newValues["is_local"] = true
	return newValues
}

//SetBuyNow - sets buy now
func (q *participants) SetBuyNow(id int64) error {
	return q.db.Exec(squirrel.Update("participants").SetMap(q.withDefaults(map[string]interface{}{
		"requested_buy_now": true,
	})).Where("id = ?", id))
}

//SetState - sets state for the participant
func (q *participants) SetState(id int64, state kafka.Participant_State) error {
	return q.db.Exec(squirrel.Update("participants").SetMap(q.withDefaults(map[string]interface{}{
		"state": int32(state),
	})).Where("id = ?", id))
}

func (q *participants) SetOutbidedWith(id, amount int64) error {
	return q.db.Exec(squirrel.Update("participants").SetMap(map[string]interface{}{
		"outbided_with": amount,
	}).Where("id = ?", id))
}
