package services

import (
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"testing"
)

func TestBuyingPowerRecharger(t *testing.T) {
	const depositAmount = 1200*1e2
	const targetDepositAmount = depositAmount*10 // 10 deposit 12,000 deposits
	cases := []struct {
		DepositAmount int64
		ExpectedNumberOfDeposits int64
	}{
		{
			DepositAmount:0,
			ExpectedNumberOfDeposits:10,
		},
		{
			DepositAmount:targetDepositAmount-depositAmount,
			ExpectedNumberOfDeposits:1,
		},
		{
			DepositAmount:targetDepositAmount-1,
			ExpectedNumberOfDeposits:1,
		},
		{
			DepositAmount:depositAmount*2,
			ExpectedNumberOfDeposits:8,
		},
	}

	for _, testCase := range cases {
		actual := getNumberOfDepositsToCreate(targetDepositAmount, depositAmount, data.IdentityWithDeposiAmount{
			Identity:      data.Identity{},
			DepositAmount: testCase.DepositAmount,
		})

		if actual != testCase.ExpectedNumberOfDeposits {
			t.Fatalf("expected %d actual %d", testCase.ExpectedNumberOfDeposits, actual)
		}
	}
}
