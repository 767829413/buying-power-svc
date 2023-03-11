// Package amount provides utilities for converting numbers to/from
// the format used internally to stellar-core.
//
// stellar-core represents asset "amounts" as 64-bit integers, but to enable
// fractional units of an asset, horizon, the client-libraries and other built
// on top of stellar-core use a convention, encoding amounts as a string of
// decimal digits with up to seven digits of precision in the fractional
// portion. For example, an amount shown as "101.001" in horizon would be
// represented in stellar-core as 1010010.
package amount

import (
	"fmt"
	"math/big"
	"strconv"
)

// One is the value of one whole unit of currency.  Stellar uses 4 fixed digits
// for fractional values, thus One is 10 thousand (10^4)
const One = 10000

// MustParse is the panicking version of Parse
func MustParse(v string) int64 {
	ret, err := Parse(v)
	if err != nil {
		panic(err)
	}
	return ret
}

// Parse parses the provided as a stellar "amount", i.e. A 64-bit signed integer
// that represents a decimal number with 7 digits of significance in the
// fractional portion of the number.
func Parse(v string) (int64, error) {
	return ParseWithOne(v, One)
}

// String returns an "amount string" from the provided raw value `v`.
func String(v int64) string {
	return StringWithOne(v, One)
}

func stringWithOne(v *big.Int, one uint) string {
	var f, o, r big.Rat

	o.SetInt64(int64(one))
	digits := len(o.FloatString(0)) - 1
	if digits < 0 {
		panic("unexpected one")
	}

	f.SetInt(v)
	r.Quo(&f, &o)

	return r.FloatString(digits)
}

func StringWithOne(v int64, one uint) string {
	return stringWithOne(big.NewInt(v), one)
}

func multOne(v string, one uint) (string, error) {
	var f, o, r big.Rat

	_, ok := f.SetString(v)
	if !ok {
		return "", fmt.Errorf("cannot parse amount: %s", v)
	}

	o.SetInt64(int64(one))
	r.Mul(&f, &o)

	return r.FloatString(0), nil
}

// ParseWithOne converts string representation of float to int64 multiplying by 'one'
func ParseWithOne(v string, one uint) (int64, error) {
	m, err := multOne(v, one)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(m, 10, 64)
}

// MustParseWithOne will panic if ParseWithOne returns error.
func MustParseWithOne(v string, one uint) int64 {
	res, err := ParseWithOne(v, one)
	if err != nil {
		panic(err)
	}

	return res
}

// StrToUint64 converts string representation of float to uint64 multiplying by 100
func StrToUint64(s string) (uint64, error) {
	m, err := multOne(s, 100)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(m, 10, 64)
}

// StrToInt64 converts string representation of float to int64 multiplying by 100
func StrToInt64(s string) (int64, error) {
	m, err := multOne(s, 100)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(m, 10, 64)
}

func Uint64ToStr(v uint64) string {
	return stringWithOne(big.NewInt(0).SetUint64(v), 100)
}

func Int64ToStr(v int64) string {
	return stringWithOne(big.NewInt(v), 100)
}
