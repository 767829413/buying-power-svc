package postgres

import (
	"database/sql"
	"gitlab.com/eAuction/events/go/kafka"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

//identities - helper struct to access identities in db
type identities struct {
	db *pgdb.DB
}

//Identities - creates new instance of identities
func (s *storage) Identities() data.Identities {
	return &identities{db: s.db}
}

//Store - stores identity into db, on conflict does nothing
func (q *identities) Store(identity data.Identity) error {
	return errors.Wrap(q.db.Exec(squirrel.Insert("identities").SetMap(map[string]interface{}{
		"id":       identity.ID,
		"platform": identity.Platform,
		"type":     identity.Type,
	}).Suffix(`on conflict on constraint identities_pkey do update
		set (
			platform,
			type
		) = (
			excluded.platform,
			excluded.type
		)`)), "failed to store identity")
}

//MarkDeleted - marks identity as deleted
func (q *identities) MarkDeleted(address string) error {
	return errors.Wrap(q.db.Exec(squirrel.Update("identities").Set("is_deleted", true).Where(squirrel.Eq{
		"id": address,
	})), "failed to mark deleted")
}

//GetForUpdate - returns identity with lock for update
func (q *identities) GetForUpdate(address string) (*data.Identity, error) {
	return q.get(address, true)
}

//Get - returns identity
func (q *identities) Get(address string) (*data.Identity, error) {
	return q.get(address, false)
}

func (q *identities) get(address string, forUpdate bool) (*data.Identity, error) {
	stmt := squirrel.Select("*").From("identities").Where(squirrel.Eq{
		"id": address,
	})
	if forUpdate {
		stmt = stmt.Suffix("for update")
	}
	var result data.Identity
	err := q.db.Get(&result, stmt)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to get identity for update")
	}

	return &result, nil
}

func (q *identities) SelectNeedRecharge(types []kafka.Identity_Type, targetBalance int64) ([]data.IdentityWithDeposiAmount, error) {
	stmt := squirrel.Select("i.*, sum(COALESCE(d.amount,0)) as deposit_amount").From("identities i").
		LeftJoin("deposits d on i.id = d.depositor AND d.state in (?,?,?,?)", kafka.DepositV2_STATE_PAID,
			kafka.DepositV2_STATE_LOCKED,
			kafka.DepositV2_STATE_RETURNING,
			kafka.DepositV2_STATE_REQUESTED_WITHDRAWAL).
		Where(
			squirrel.Eq{
				"i.type": types,
			}).
		Where("NOT is_deleted").
		GroupBy("i.id").Having("sum(COALESCE(d.amount,0)) < ?", targetBalance)
	var result []data.IdentityWithDeposiAmount
	err := q.db.Select(&result, stmt)
	return result, errors.Wrap(err, "failed to exec select")
}
