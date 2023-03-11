package config

import (
	"reflect"
	"strings"
	"time"

	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//Limits - defines deposit-limit pairs
type Limits struct {
	Limits                  []Limit `fig:"limits"`
	supportedDepositAmounts comfig.Once
	maxLimit                comfig.Once
}

//SupportedDepositAmounts - returns list of supported deposit amount suitable for validation
func (d Limits) SupportedDepositAmounts() []interface{} {
	return d.supportedDepositAmounts.Do(func() interface{} {
		result := make([]interface{}, len(d.Limits))
		for i := range d.Limits {
			result[i] = d.Limits[i].Deposit
		}

		return result
	}).([]interface{})
}

//GetMaxLimit - returns max limit
func (d Limits) GetMaxLimit() Limit {
	return d.maxLimit.Do(func() interface{} {
		var maxLimit Limit
		for _, limit := range d.Limits {
			if maxLimit.Limit < limit.Limit {
				maxLimit = limit
			}
		}

		if maxLimit.Limit == 0 || maxLimit.Deposit == 0 {
			panic(errors.New("expected at least one limit to be available - try setting limits or disabling buying power recharger"))
		}

		return maxLimit
	}).(Limit)
}

//FindMatchingLimit - tries to find highest limit that is satisfied with the deposit requirements
func (d Limits) FindMatchingLimit(depositAmount int64) int64 {
	var result int64
	for _, limit := range d.Limits {
		if limit.Deposit <= depositAmount && limit.Limit > result {
			result = limit.Limit
		}
	}

	return result
}

//Deposits -
type Deposits struct {
	ExpiresIn             time.Duration     `fig:"expires_in"`
	LotProcessingDuration time.Duration     `fig:"lot_processing_duration"`
	Currency              string            `fig:"currency,required"`
	MaxAmount             int64             `fig:"max_amount,required"`
	DefaultLimits         Limits            `fig:"default_limits"`
	PlatformLimits        map[string]Limits `fig:"platform_limits"`
	// percentage for computing deposit amount for lot
	DepositFraction *float64 `fig:"deposit_fraction"`
	//
	ReturnDepositsToCard  bool `fig:"return_deposits_to_card"`
	AutoApproveWithdrawal bool `fig:"auto_approve_withdrawal"`
}

//GetLimits - returns limits for specified platform
func (d Deposits) GetLimits(platformID string) Limits {
	result, ok := d.PlatformLimits[platformID]
	if ok {
		return result
	}

	return d.DefaultLimits
}

//Limit - defines bidding limit for specified amount of deposit
type Limit struct {
	Deposit         int64 `fig:"deposit,required"`
	Limit           int64 `fig:"limit,required"`
	IsBuyNowAllowed bool  `json:"is_buy_now_allowed"`
	IsRechargerOnly bool  `json:"is_recharger_only"`
}

func (c *config) Deposits() Deposits {
	return c.deposits.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "deposits")

		var cfg Deposits

		if err := figure.Out(&cfg).With(localHooks, limitsHook, figure.BaseHooks).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out deposits config"))
		}
		if (len(cfg.DefaultLimits.Limits) == 0) == (cfg.DepositFraction == nil) {
			panic("deposits: either limits or deposit_fraction must be set, not both")
		}

		return cfg
	}).(Deposits)
}

var localHooks = figure.Hooks{
	"map[string]config.Limits": func(value interface{}) (reflect.Value, error) {
		platformLimits := map[string]Limits{}
		for platformID, v := range cast.ToStringMap(value) {
			var limits Limits
			err := figure.
				Out(&limits).
				With(limitsHook, figure.BaseHooks).
				From(cast.ToStringMap(v)).
				Please()

			if err != nil {
				panic(errors.Wrap(err, "failed to figure out platform limits", logan.F{
					"platform_id": platformID,
				}))
			}

			platformLimits[strings.ToUpper(platformID)] = limits
		}

		return reflect.ValueOf(platformLimits), nil
	},
	"*float64": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToFloat64E(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse *float64")
		}
		return reflect.ValueOf(&result), nil
	},
}

var limitsHook = figure.Hooks{
	"[]config.Limit": func(value interface{}) (reflect.Value, error) {
		rawSlice, err := cast.ToSliceE(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "expected slice of Limit")
		}

		deposits := make([]Limit, len(rawSlice))
		for idx, val := range rawSlice {
			raw, err := cast.ToStringMapE(val)
			if err != nil {
				return reflect.Value{},
					errors.Wrap(err, "expected Limit to be map[string]interface{}")
			}

			var opts Limit
			if err = figure.Out(&opts).From(raw).Please(); err != nil {
				return reflect.Value{}, errors.Wrap(err, "malformed DepositOpts")
			}

			deposits[idx] = opts
		}

		return reflect.ValueOf(deposits), nil
	},
}
