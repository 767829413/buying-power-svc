package power

import (
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"gitlab.com/eAuction/buying-power-svc/internal/config"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/types"
	"gitlab.com/eAuction/events/go/kafka"
)

const (
	defaultCurrency = "USD"
	depositor       = "depositor"
	maxBalance      = 5000000
)

var (
	minLimit = config.Limit{
		Deposit: 300,
	}
	midLimit = config.Limit{
		Deposit: 600,
	}
	maxLimit = config.Limit{
		Deposit: 1200,
	}
	minDeposit = newDeposit("min_deposit", minLimit.Deposit)
	midDeposit = newDeposit("mid_deposit", midLimit.Deposit)
	maxDeposit = newDeposit("max_deposit", maxLimit.Deposit)
)

func newDeposit(id string, amount int64) data.Deposit {
	return data.Deposit{
		ID:        id,
		Amount:    amount,
		Depositor: depositor,
		State:     int32(kafka.DepositV2_STATE_PENDING),
		Currency:  defaultCurrency,
	}
}

type testCase struct {
	balance           int64
	attachedDeposits  []data.Deposit
	availableDeposits []data.Deposit
	limit             config.Limit

	expectedState    types.DepositOptionState
	depositsToAttach []data.Deposit
	depositToCreate  *data.Deposit
}

func (c testCase) Run(t *testing.T, calc Calculator) {
	state, depositsToAttach, depositToCreate := calc.GetDepositsOptions(depositor, c.balance, c.attachedDeposits, c.availableDeposits, c.limit)
	assert.Equal(t, state, c.expectedState)
	assert.Equal(t, depositsToAttach, c.depositsToAttach)
	assert.Equal(t, depositToCreate, c.depositToCreate)
}
// TODO: fix me
func TestBuyingPower(t *testing.T) {
	cfg := &config.MockConfig{}
	defer cfg.AssertExpectations(t)
	depositIDGenerator := &data.MockDepositIDGenerator{}
	defer depositIDGenerator.AssertExpectations(t)
	cfg.On("ExRatesConverter").Return(nil)
	cfg.On("Deposits").Return(config.Deposits{
		ExpiresIn:             25 * 24 * time.Hour,
		LotProcessingDuration: 3 * 24 * time.Hour,
		Currency:              defaultCurrency,
		MaxAmount:             maxBalance,
		// Limits:                []config.Limit{minLimit, midLimit, maxLimit},
	})
	cfg.On("Platformer").Return(nil)
	cfg.On("DepositIDGenerator").Return(depositIDGenerator)
	calc := NewCalculator(cfg)
	t.Run("option is already taken", func(t *testing.T) {
		deposit := newDeposit("taken", minLimit.Deposit)
		testCase{
			balance:           0,
			attachedDeposits:  []data.Deposit{deposit},
			availableDeposits: []data.Deposit{minDeposit, midDeposit, maxDeposit},
			limit:             minLimit,
			expectedState:     types.DepositOptionStateTaken,
			depositsToAttach:  []data.Deposit{deposit},
			depositToCreate:   nil,
		}.Run(t, *calc)
	})
	t.Run("option is not available as higher is already taken", func(t *testing.T) {
		testCase{
			balance:           0,
			attachedDeposits:  []data.Deposit{midDeposit},
			availableDeposits: []data.Deposit{minDeposit, midDeposit, maxDeposit},
			limit:             minLimit,
			expectedState:     types.DepositOptionStateNotAvailable,
			depositsToAttach:  nil,
			depositToCreate:   nil,
		}.Run(t, *calc)
	})
	t.Run("prefer getting new exact deposit rather than available", func(t *testing.T) {
		depositID := data.NewDepositIDGenerator().Generate()
		depositIDGenerator.On("Generate").Return(depositID).Once()
		deposit := newDeposit(depositID, minLimit.Deposit)
		testCase{
			balance:           0,
			attachedDeposits:  []data.Deposit{},
			availableDeposits: []data.Deposit{midDeposit, maxDeposit},
			limit:             minLimit,
			expectedState:     types.DepositOptionStateAvailable,
			depositsToAttach:  []data.Deposit{},
			depositToCreate:   &deposit,
		}.Run(t, *calc)
	})
	t.Run("prefer getting new exact deposit with attached deposits rather than available with attached deposit", func(t *testing.T) {
		depositID := data.NewDepositIDGenerator().Generate()
		depositIDGenerator.On("Generate").Return(depositID).Once()
		deposit := newDeposit(depositID, minLimit.Deposit)
		testCase{
			balance:           0,
			attachedDeposits:  []data.Deposit{minDeposit},
			availableDeposits: []data.Deposit{midDeposit, maxDeposit},
			limit:             midLimit,
			expectedState:     types.DepositOptionStateAvailable,
			depositsToAttach:  []data.Deposit{minDeposit},
			depositToCreate:   &deposit,
		}.Run(t, *calc)
	})
	t.Run("exact match from available", func(t *testing.T) {
		testCase{
			balance:           0,
			attachedDeposits:  []data.Deposit{},
			availableDeposits: []data.Deposit{minDeposit, midDeposit, maxDeposit},
			limit:             midLimit,
			expectedState:     types.DepositOptionStateAvailable,
			depositsToAttach:  []data.Deposit{midDeposit},
			depositToCreate:   nil,
		}.Run(t, *calc)
	})
	t.Run("attached + available", func(t *testing.T) {
		testCase{
			balance:           0,
			attachedDeposits:  []data.Deposit{minDeposit},
			availableDeposits: []data.Deposit{minDeposit},
			limit:             midLimit,
			expectedState:     types.DepositOptionStateAvailable,
			depositsToAttach:  []data.Deposit{minDeposit, minDeposit},
			depositToCreate:   nil,
		}.Run(t, *calc)
	})
	t.Run("attached + new deposit", func(t *testing.T) {
		depositID := data.NewDepositIDGenerator().Generate()
		depositIDGenerator.On("Generate").Return(depositID).Once()
		deposit := newDeposit(depositID, minLimit.Deposit)
		testCase{
			balance:          0,
			attachedDeposits: []data.Deposit{minDeposit},
			limit:            midLimit,
			expectedState:    types.DepositOptionStateAvailable,
			depositsToAttach: []data.Deposit{minDeposit},
			depositToCreate:  &deposit,
		}.Run(t, *calc)
	})
	t.Run("not able to create new deposit", func(t *testing.T) {
		testCase{
			balance:           maxBalance,
			attachedDeposits:  []data.Deposit{},
			availableDeposits: []data.Deposit{minDeposit},
			limit:             midLimit,
			expectedState:     types.DepositOptionStateNotAvailable,
			depositsToAttach:  nil,
			depositToCreate:   nil,
		}.Run(t, *calc)
	})
	t.Run("not able to make match exact as it's not possible to create new deposit", func(t *testing.T) {
		testCase{
			balance:           maxBalance,
			attachedDeposits:  []data.Deposit{minDeposit},
			availableDeposits: []data.Deposit{midDeposit},
			limit:             midLimit,
			expectedState:     types.DepositOptionStateAvailable,
			depositsToAttach:  []data.Deposit{minDeposit, midDeposit},
			depositToCreate:   nil,
		}.Run(t, *calc)
	})
}
