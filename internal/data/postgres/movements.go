package postgres

import (
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

//movements - helper struct to access movements stored in db
type movements struct {
	db *pgdb.DB
}

//Movements - creates new instance of movements
func (s *storage) Movements() data.Movements {
	return &movements{db: s.db}
}

//Insert - creates new movement
func (q *movements) Insert(movement data.Movement) error {
	s := squirrel.Insert("movements").SetMap(map[string]interface{}{
		"identity":   movement.Identity,
		"amount":     movement.Amount,
		"action":     movement.Action,
		"currency":   movement.Currency,
		"created_at": movement.CreatedAt.UTC(),
		"lot_id":     movement.Lot,
	})

	err := q.db.Exec(s)
	if err != nil {
		return errors.Wrap(err, "failed to exec statement")
	}

	return nil
}

//Delete - deletes movement
func (q *movements) Delete(id []int64) error {
	s := squirrel.Delete("movements").Where(squirrel.Eq{
		"id": id,
	})
	return q.db.Exec(s)
}

//SelectNotSynced - returns list of not synced movements
func (q *movements) SelectNotSynced(limit uint64) ([]data.Movement, error) {
	s := squirrel.Select("m.*").From("movements m").OrderBy("id asc").Limit(limit)
	var result []data.Movement
	err := q.db.Select(&result, s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select not synced movements")
	}

	return result, nil
}
