package test

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

// Due to unknown reasons, when I try to run tests in any lower level packages test just freezes
func TestDepositByCreateAtAscSort(t *testing.T) {
	in := []data.Deposit{
		{
			CreatedAt: time.Unix(10, 0),
		},
		{
			CreatedAt: time.Unix(100, 0),
		},
		{
			CreatedAt: time.Unix(20, 0),
		},
	}

	sort.Sort(data.DepositsByCreatedAtAsc(in))
	assert.Equal(t, int64(10), in[0].CreatedAt.Unix())
	assert.Equal(t, int64(20), in[1].CreatedAt.Unix())
	assert.Equal(t, int64(100), in[2].CreatedAt.Unix())

}
