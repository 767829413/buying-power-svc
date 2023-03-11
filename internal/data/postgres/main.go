package postgres

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

type storage struct {
	db *pgdb.DB
}

//New - returns new instance of storage
func New(db *pgdb.DB) data.Storage {
	return &storage{db: db}
}

func (s *storage) Transaction(fn pgdb.TransactionFunc) error {
	return s.db.Transaction(fn)
}

func (s *storage) Clone() data.Storage {
	return New(s.db.Clone())
}
