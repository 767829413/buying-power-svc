package exrates

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//Connector - helper struct to communicate with exrates-svc.
type Connector struct {
	client  *http.Client
	baseURL url.URL
}

//NewConnector - create new instance of thread safe connector for exrates-svc
func NewConnector(baseURL url.URL) *Connector {
	return &Connector{
		client: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
		baseURL: baseURL,
	}
}

//GetLatest - returns latest list of all available exchange rates
func (c *Connector) GetLatest(baseAsset string) (*ExchangeRateListResponse, error) {
	uri := c.baseURL
	uri.Path += "/latest/" + baseAsset

	resp, err := c.client.Get(uri.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, nil
	case http.StatusOK:
		// do nothing
	default:
		return nil, errors.Wrap(err, "expected 200 status code", logan.F{
			"status_code": resp.StatusCode,
		})
	}

	var respBody ExchangeRateListResponse
	if err = json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	return &respBody, nil
}
