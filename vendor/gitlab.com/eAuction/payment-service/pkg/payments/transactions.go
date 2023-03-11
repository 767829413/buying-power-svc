package payments

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/payment-service/pkg/payments/resources"
)

//GetLots - returns list of the transactions
func (c *paymentService) GetTransactions(params GetTransactionsParams) (resources.TransactionListResponse, error) {
	resp, err := c.get([]string{"transactions"}, params)

	if resp != nil {
		defer func() { _ = resp.Body.Close() }()
	}
	if err != nil {
		return resources.TransactionListResponse{}, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var response resources.TransactionListResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return resources.TransactionListResponse{}, errors.Wrap(err, "malformed response")
		}

		return response, nil
	default:
		respBB, _ := ioutil.ReadAll(resp.Body)
		return resources.TransactionListResponse{}, errors.From(errors.New("unexpected status code"), logan.F{
			"url":          resp.Request.RequestURI,
			"raw_response": string(respBB),
			"status_code":  resp.StatusCode,
		})
	}
}

func (c *paymentService) GetTransaction(id string) (resources.TransactionResponse, error) {
	resp, err := c.get([]string{"transactions", id}, nil)

	if resp != nil {
		defer func() { _ = resp.Body.Close() }()
	}
	if err != nil {
		return resources.TransactionResponse{}, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var response resources.TransactionResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return resources.TransactionResponse{}, errors.Wrap(err, "malformed response")
		}

		return response, nil
	case http.StatusNotFound:
		return resources.TransactionResponse{}, ErrNotFound
	default:
		respBB, _ := ioutil.ReadAll(resp.Body)
		return resources.TransactionResponse{}, errors.From(errors.New("unexpected status code"), logan.F{
			"url":          resp.Request.RequestURI,
			"raw_response": string(respBB),
			"status_code":  resp.StatusCode,
		})
	}
}
