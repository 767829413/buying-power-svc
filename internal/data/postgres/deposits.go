package postgres

import (
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/events/go/kafka"
)

//deposits - helper struct to access deposits stored in db
type deposits struct {
	db *pgdb.DB
}

//Deposits - creates new instance of deposits
func (s *storage) Deposits() data.Deposits {
	return &deposits{db: s.db}
}

//CreateLegacy - creates legacy deposit
func (q *deposits) CreateLegacy(deposit data.Deposit) error {
	s := squirrel.Insert("deposits").SetMap(map[string]interface{}{
		"id":         deposit.ID,
		"state":      deposit.State,
		"amount":     deposit.Amount,
		"currency":   deposit.Currency,
		"depositor":  deposit.Depositor,
		"created_at": deposit.CreatedAt,
		"updated_at": deposit.UpdatedAt,
		"lot_id":     deposit.LotID,
		"synced":     deposit.Synced,
	})

	err := q.db.Exec(s)
	if err != nil {
		return errors.Wrap(err, "failed to exec statement")
	}

	return nil
}

//Create - creates new deposit instance
func (q *deposits) Create(amount int64, currency, depositor string) (data.Deposit, error) {
	now := time.Now().UTC()

	s, args, err := squirrel.Insert("deposits").SetMap(map[string]interface{}{
		"id":         data.NewDepositIDGenerator().Generate(),
		"state":      int32(kafka.DepositV2_STATE_PENDING),
		"amount":     amount,
		"currency":   currency,
		"depositor":  depositor,
		"created_at": now,
		"updated_at": now,
	}).Suffix("returning *").ToSql()

	if err != nil {
		return data.Deposit{}, errors.Wrap(err, "failed to build query")
	}

	var result data.Deposit
	err = q.db.GetRaw(&result, s, args...)
	if err != nil {
		return data.Deposit{}, errors.Wrap(err, "failed to exec statement")
	}

	return result, nil

}

//Get - gets deposit by id, returns nil, nil, if not found
func (q *deposits) Get(id string, forUpdate bool) (*data.Deposit, error) {
	stmt := squirrel.Select("*").From("deposits").Where(squirrel.Eq{"id": id})
	if forUpdate {
		stmt = stmt.Suffix("for update")
	}
	var result data.Deposit
	err := q.db.Get(&result, stmt)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

//GetForUpdate - gets deposit by id, returns nil, nil, if not found
func (q *deposits) GetForUpdate(id string) (*data.Deposit, error) {
	return q.Get(id, true)
}

//SetState - sets state for specified deposit
func (q *deposits) SetState(id string, state int32) error {
	err := q.db.Exec(squirrel.Update("deposits").SetMap(map[string]interface{}{
		"state":      state,
		"updated_at": time.Now().UTC(),
		"synced":     false,
	}).Where(squirrel.Eq{"id": id}))
	return errors.Wrap(err, "failed to set state for deposit")
}

//ManuallyApprove - marks deposit as paid and artificial
func (q *deposits) ManuallyApprove(id string) error {
	err := q.db.Exec(squirrel.Update("deposits").SetMap(map[string]interface{}{
		"state":         int32(kafka.DepositV2_STATE_PAID),
		"updated_at":    time.Now().UTC(),
		"synced":        false,
		"is_artificial": true,
	}).Where(squirrel.Eq{"id": id}))
	return errors.Wrap(err, "failed to set state for deposit")
}

//SetLot - sets lot for deposit
func (q *deposits) SetLot(id string, lotID *string) error {
	err := q.db.Exec(squirrel.Update("deposits").SetMap(map[string]interface{}{
		"lot_id":     lotID,
		"updated_at": time.Now().UTC(),
		"synced":     false,
	}).Where(squirrel.Eq{"id": id}))
	return errors.Wrap(err, "failed to set lot for deposit", logan.F{
		"lot_id":  lotID,
		"deposit": id,
	})
}

// ReturnDepositToBalance - returns deposit back to balance
func (q *deposits) ReturnDepositToBalance(id string) error {
	err := q.db.Exec(squirrel.Update("deposits").SetMap(map[string]interface{}{
		"state":      int32(kafka.DepositV2_STATE_PAID),
		"updated_at": time.Now().UTC(),
		"synced":     false,
		"lot_id":     nil,
	}).Where(squirrel.Eq{"id": id}))
	return errors.Wrap(err, "failed to update deposit")
}

// ScheduleReturnToCard - marks deposit to return to card
func (q *deposits) ScheduleReturnToCard(id string) error {
	err := q.db.Exec(squirrel.Update("deposits").SetMap(map[string]interface{}{
		"state":      int32(kafka.DepositV2_STATE_RETURNING),
		"updated_at": time.Now().UTC(),
		"synced":     false,
		"lot_id":     nil,
	}).Where(squirrel.Eq{"id": id}))
	return errors.Wrap(err, "failed to update deposit")
}

//Select - selects deposits for specified filter
func (q *deposits) Select(selector data.DepositsSelector) ([]data.Deposit, error) {
	s := squirrel.Select("*").From("deposits")
	if selector.ForUpdate {
		s = s.Suffix("for update")
	}

	if len(selector.Lots) != 0 {
		s = s.Where(squirrel.Eq{
			"lot_id": selector.Lots,
		})
	}

	if len(selector.States) != 0 {
		s = s.Where(squirrel.Eq{
			"state": selector.States,
		})
	}

	if len(selector.Depositors) != 0 {
		s = s.Where(squirrel.Eq{
			"depositor": selector.Depositors,
		})
	}

	if selector.HasLot != nil {
		if *selector.HasLot {
			s = s.Where(squirrel.NotEq{
				"lot_id": nil,
			})
		} else {
			s = s.Where(squirrel.Eq{
				"lot_id": nil,
			})
		}
	}

	if selector.MaxCreatedAt != nil {
		s = s.Where(squirrel.LtOrEq{
			"created_at": selector.MaxCreatedAt.UTC(),
		})
	}

	if selector.MinCreatedAt != nil {
		s = s.Where(squirrel.GtOrEq{
			"created_at": selector.MinCreatedAt.UTC(),
		})
	}

	if selector.Currency != nil {
		s = s.Where(squirrel.GtOrEq{
			"currency": *selector.Currency,
		})
	}

	if selector.Page != nil {
		s = selector.Page.ApplyTo(s, "id")
	} else {
		s = s.OrderBy("id asc")
	}

	if selector.IsSynced != nil {
		s = s.Where(squirrel.Eq{
			"synced": selector.IsSynced,
		})
	}

	if selector.IsArtificial != nil {
		s = s.Where(squirrel.Eq{
			"is_artificial": *selector.IsArtificial,
		})
	}

	var result []data.Deposit
	err := q.db.Select(&result, s)
	return result, errors.Wrap(err, "failed to select deposits")
}

//MarkSynced - marks deposits as synced
func (q *deposits) MarkSynced(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	return q.db.Exec(squirrel.Update("deposits").SetMap(map[string]interface{}{
		"synced": true,
	}).Where(squirrel.Eq{
		"id": ids,
	}))
}

//SelectRequireRefund - returns list of deposits that should be returned to user as it's not possible to attach them to the lot
func (q *deposits) SelectRequireRefund(minAllowedCreatedAt time.Time) ([]data.Deposit, error) {
	s := squirrel.Select("d.*").
		From("deposits d").
		Where(squirrel.And{
			squirrel.Eq{
				"d.state": []int32{int32(kafka.DepositV2_STATE_PAID)},
			},
			squirrel.LtOrEq{
				"d.created_at": minAllowedCreatedAt.UTC(),
			},
		})

	var result []data.Deposit
	err := q.db.Select(&result, s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select deposits")
	}

	return result, nil
}
