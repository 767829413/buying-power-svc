package data

//IdentityToBroker - tracks identities to brokers connections
type IdentityToBroker struct {
	Identity string `db:"identity"`
	Broker   string `db:"broker"`
}

//IdentityToBrokers - allows access identity to broker entities
type IdentityToBrokers interface {
	Insert(entity IdentityToBroker) error
	Delete(identity string) error
	GetBrokersForIdentity(identity string) ([]string, error)
}