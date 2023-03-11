package exrates

import (
	"math"
	"net/url"
	"sync"
	"time"

	"github.com/shopspring/decimal"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//ErrOverflow - specifies that conversion failed due to overflow
var ErrOverflow = errors.New("overflow: result does not fit into int64")


//Converter - helper struct to convert one currency into another using cached response from exrates-svc
type Converter struct {
	connector       exRatesProvider
	minUpdatePeriod time.Duration
	ttl             time.Duration

	rates          map[string]decimal.Decimal
	ratesUpdatedAt time.Time
	triedToUpdateAt time.Time
	lock           sync.Mutex
	log            *logan.Entry

	convertThrough string
	Now            func() time.Time
}

//go:generate mockery -case underscore -name exRatesProvider -inpkg
type exRatesProvider interface {
	GetLatest(baseAsset string) (*ExchangeRateListResponse, error)
}

// NewConverter - create new instance of converter
// minUpdatePeriod - defines period at which we would try to update cache
// ttl - defines period of time after which cache is considered invalid.
func NewConverter(log *logan.Entry, baseURL url.URL, minUpdatePeriod, ttl time.Duration) *Converter {
	if minUpdatePeriod > ttl {
		panic(errors.New("expected minUpdatePeriod to <= ttl"))
	}
	return &Converter{
		connector:       NewConnector(baseURL),
		minUpdatePeriod: minUpdatePeriod,
		ttl:             ttl,
		convertThrough:  "EUR",
		Now:             time.Now,
		log:             log,
	}
}

//ConvertCents - converts amount from 'from' currency into 'to' currency.
func (c *Converter) ConvertCents(from, to string, amount int64) (int64, time.Time, error) {
	if from == to {
		return amount, c.Now(), nil
	}

	if amount == 0 {
		return 0, c.Now(), nil
	}

	actualRates, ratesUpdateTime, err := c.getLatestRates()
	if err != nil {
		return 0, time.Time{}, errors.Wrap(err, "failed to get latest rates")
	}
	rateFrom, ok := actualRates[from]
	if !ok {
		return 0, time.Time{}, errors.From(errors.New("trying to convert from not supported currency"), logan.F{
			"from": from,
		})
	}

	rateTo, ok := actualRates[to]
	if !ok {
		return 0, time.Time{}, errors.From(errors.New("trying to convert into not supported currency"), logan.F{
			"to": to,
		})
	}

	decimalAmount := decimal.New(amount, 0)
	result := decimalAmount.Mul(rateTo).DivRound(rateFrom,0)
	if result.Cmp(decimal.New(math.MaxInt64, 0)) > 1 {
		return 0, time.Time{}, errors.From(ErrOverflow, logan.F{
			"result": result.String(),
		})
	}
	return result.IntPart(), ratesUpdateTime, nil
}

func (c *Converter) getLatestRates() (map[string]decimal.Decimal, time.Time, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	now := c.Now()
	shouldUpdate := c.triedToUpdateAt.Add(c.minUpdatePeriod).Before(now)
	if !shouldUpdate {
		return c.rates, c.ratesUpdatedAt, nil
	}

	c.triedToUpdateAt = now
	mustUpdate := c.ratesUpdatedAt.Add(c.ttl).Before(now)
	latest, updatedAt, err := c.loadRates()
	if err != nil {
		if mustUpdate {
			return nil, time.Time{}, errors.Wrap(err, "failed to update exchange rates")
		}

		c.log.WithError(err).Warn("failed to updated exchange rates, but we still have valid cache")
		return c.rates, c.ratesUpdatedAt, nil
	}

	c.rates = latest
	c.ratesUpdatedAt = updatedAt
	return c.rates, c.ratesUpdatedAt, nil
}

func (c *Converter) loadRates() (map[string]decimal.Decimal, time.Time, error) {
	updatedRatesResponse, err := c.connector.GetLatest(c.convertThrough)
	if err != nil {
		c.log.WithError(err).Warn("failed to get latest exchange rates from exrates-svc")
		return nil, time.Time{}, err
	}

	var oldestRate time.Time
	result := map[string]decimal.Decimal{}
	for _, rate := range updatedRatesResponse.Data {
		base := rate.Relationships.BaseAsset.Data.ID
		if base != c.convertThrough {
			return nil, time.Time{}, errors.From(errors.New("unexpected response from exrates-svc. Received rates for not requested pair"), logan.F{
				"base":          base,
				"expected_base": c.convertThrough,
			})
		}

		quote := rate.Relationships.QuoteAsset.Data.ID
		if _, ok := result[quote]; ok {
			return nil, time.Time{}, errors.From(errors.New("received list of exchange rates with duplication"), logan.F{
				"base":  base,
				"quote": quote,
			})
		}

		result[quote] = decimal.NewFromFloat(rate.Attributes.Price)
		if oldestRate.IsZero() || oldestRate.After(rate.Attributes.UpdatedAt) {
			oldestRate = rate.Attributes.UpdatedAt
		}
	}

	result[c.convertThrough] = decimal.New(1,0)

	return result, oldestRate, nil
}
