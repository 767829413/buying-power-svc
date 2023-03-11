package postgres

import (
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

//identitiesToBrokers - helper struct to access identities to brokers in db
type identityToBrokers struct {
	db *pgdb.DB
}

//IdentitiesToBrokers - creates new instance of IdentityToBrokers
func (s *storage) IdentitiesToBrokers() data.IdentityToBrokers {
	return &identityToBrokers{db: s.db}
}

func (s *identityToBrokers) Insert(entity data.IdentityToBroker) error {
	stmt := squirrel.Insert("identity_to_brokers").Columns("broker", "identity").Values(entity.Broker, entity.Identity)
	return s.db.Exec(stmt)
}
func (s *identityToBrokers) Delete(identity string) error {
	return s.db.Exec(squirrel.Delete("identity_to_brokers").Where("identity = ?", identity))
}

func (s identityToBrokers) GetBrokersForIdentity(identity string) ([]string, error) {
	var result []string
	return result, s.db.Select(&result, squirrel.Select("broker").From("identity_to_brokers").Where(squirrel.Eq{
		"identity": identity,
	}))
}
