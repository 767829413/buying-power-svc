package alias

import (
	"bytes"
	"encoding/csv"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/enums/go/internal/assets"
)

func readBoxData(name string) ([][]string, error) {
	input, err := assets.Assets.Find(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find box", logan.F{
			"name": name,
		})
	}
	file := bytes.NewBuffer(input)
	return csv.NewReader(file).ReadAll()
}
