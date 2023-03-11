package postgres

import (
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/lorem"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

//withdrawals - helper struct to access withdrawals stored in db
type withdrawals struct {
	db *pgdb.DB
}

//Withdrawals - creates new instance of withdrawals
func (s *storage) Withdrawals() data.Withdrawals {
	return &withdrawals{db: s.db}
}

//Create - creates new deposit instance
func (q *withdrawals) Create(depositID string) (data.Withdrawal, error) {
	now := time.Now().UTC()

	s, args, err := squirrel.Insert("withdrawals").SetMap(map[string]interface{}{
		"id":         "withdrawal:" + lorem.ULID(),
		"deposit_id": depositID,
		"state":      int32(kafka.Withdrawal_STATE_PENDING),
		"created_at": now,
		"updated_at": now,
	}).Suffix("returning *").ToSql()

	if err != nil {
		return data.Withdrawal{}, errors.Wrap(err, "failed to build query")
	}

	var result data.Withdrawal
	err = q.db.GetRaw(&result, s, args...)
	if err != nil {
		return data.Withdrawal{}, errors.Wrap(err, "failed to exec statement")
	}

	return result, nil

}

/*
Create(depositID string) (Withdrawal, error)
	GetForUpdate(id string) (*Withdrawal, error)
	SetState(id string, state int32) error
	Select(selector WithdrawalsSelector) ([]Withdrawal, error)
	MarkSynced(ids []string) error
*/

//GetForUpdate - gets withdrawal by id, returns nil, nil, if not found
func (q *withdrawals) GetForUpdate(id string) (*data.Withdrawal, error) {
	var result data.Withdrawal
	err := q.db.Get(&result, squirrel.Select("*").Suffix("for update").From("withdrawals").Where(squirrel.Eq{"id": id}))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

//SetState - sets state for specified withdrawal
func (q *withdrawals) SetState(id string, state int32) error {
	err := q.db.Exec(squirrel.Update("withdrawals").SetMap(map[string]interface{}{
		"state":      state,
		"updated_at": time.Now().UTC(),
		"synced":     false,
	}).Where(squirrel.Eq{"id": id}))
	return errors.Wrap(err, "failed to set state for withdrawal")
}

//Select - selects deposits for specified filter
func (q *withdrawals) Select(selector data.WithdrawalsSelector) ([]data.Withdrawal, error) {
	s := squirrel.Select("*").From("withdrawals").OrderBy("id asc")
	if selector.ForUpdate {
		s = s.Suffix("for update")
	}

	if len(selector.States) != 0 {
		s = s.Where(squirrel.Eq{
			"state": selector.States,
		})
	}

	if selector.IsSynced != nil {
		s = s.Where(squirrel.Eq{
			"synced": selector.IsSynced,
		})
	}

	if len(selector.Deposits) != 0 {
		s = s.Where(squirrel.Eq{
			"deposit_id": selector.Deposits,
		})
	}

	var result []data.Withdrawal
	err := q.db.Select(&result, s)
	return result, errors.Wrap(err, "failed to select withdrawals")
}

//MarkSynced - marks deposits as synced
func (q *withdrawals) MarkSynced(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	return q.db.Exec(squirrel.Update("withdrawals").SetMap(map[string]interface{}{
		"synced": true,
	}).Where(squirrel.Eq{
		"id": ids,
	}))
}
