package copart

import (
	"bytes"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type removeWatchListData struct {
	ErrorCode string `json:"errorCode"`
}

type removeWatchListResponse struct {
	responseMeta
	Data removeWatchListData `json:"data"`
}

//RemoveFromWatchlist - removes specified lots from the watch list
func (c *connector) RemoveFromWatchlist(lots []string) error {
	rawBody := map[string][]string{
		"lots": lots,
	}

	jsonBody, err := json.Marshal(rawBody)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal body"))
	}
	req, err := http.NewRequest(http.MethodPost, "https://www.copart.com/data/lots/remove/checked/watchlist", bytes.NewReader(jsonBody))
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	c.populateHeaders(req)

	var result removeWatchListResponse
	err = c.doJson(req, &result)
	if err != nil {
		return errors.Wrap(err, "failed to perform request")
	}

	if result.Data.ErrorCode != "" {
		return errors.From(errors.New("request rejected by copart"), logan.F{
			"error_code": result.Data.ErrorCode,
		})
	}

	return nil
}

func (c *connector) populateHeaders(req *http.Request) {
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("x-xsrf-token", c.csrfToken)
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("access-control-allow-headers", "Content-Type, X-XSRF-TOKEN")
	req.Header.Set("accept", "application/json, text/plain, */*")
}
