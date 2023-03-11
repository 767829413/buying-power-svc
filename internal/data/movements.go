package data

import (
	"database/sql"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/eAuction/events/go/kafka"
	"time"
)

//Movement - represents movements over balance
type Movement struct {
	ID        int64                 `db:"id"`
	Identity  string                `db:"identity"`
	Amount    int64                 `db:"amount"`
	Currency  string                `db:"currency"`
	Action    kafka.Movement_Action `db:"action"`
	CreatedAt time.Time             `db:"created_at"`
	Lot       sql.NullString        `db:"lot_id"`
}

//MovementsSelector - helper struct to describe select filters
type MovementsSelector struct {
	Identity *string
	Page     pgdb.CursorPageParams
}

//Movements - helper interface to access movements in db
type Movements interface {
	Insert(movement Movement) error
	Delete(id []int64) error
	SelectNotSynced(limit uint64) ([]Movement, error)
}
