package data

import "gitlab.com/distributed_lab/lorem"

const DepositIDPrefix = "deposit:"

//DepositIDGenerator - generates deposit ids
//go:generate mockery -case underscore -name DepositIDGenerator -inpkg
type DepositIDGenerator interface {
	Generate() string
}

//NewDepositIDGenerator - returns new instance of geneator
func NewDepositIDGenerator() DepositIDGenerator {
	return &depositIDGenerator{}
}

type depositIDGenerator struct {
}

//Generate - generates new id for deposit
func (d *depositIDGenerator) Generate() string {
	return DepositIDPrefix + lorem.ULID()
}
