package copart

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type watchListElem struct {
	LotID string `json:"lotId"`
}

type watchlistData struct {
	Watchlist []watchListElem `json:"watchList"`
}

type watchlistResponse struct {
	responseMeta
	Data watchlistData `json:"data"`
}

func (c *connector) GetWatchlist() ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.copart.com/data/lots/watchList", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create new request"))
	}

	c.populateHeaders(req)

	var resultBody watchlistResponse
	err = c.doJson(req, &resultBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	var result []string
	for _, lotID := range resultBody.Data.Watchlist {
		result = append(result, lotID.LotID)
	}

	return result, nil
}
