package payments

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/eAuction/go/amount"
	"gitlab.com/eAuction/payment-service/pkg/payments/resources"
)

type GetTransactionsParams struct {
	Status     *string `url:"status"`
	InvoiceID  *string `url:"invoice_id"`
	AccountID  *string `url:"account_id"`
	Purpose    *string `url:"purpose"`
	Sort       *string `url:"sort"`
	PageCursor *string `page:"cursor"`
	PageLimit  *int64  `page:"limit"`
}

//go:generate mockery --case underscore --name PaymentService --inpackage
type PaymentService interface {
	GetTransactions(params GetTransactionsParams) (resources.TransactionListResponse, error)
	GetTransaction(id string) (resources.TransactionResponse, error)
	ConfirmTransaction(id string, amount *amount.Fiat) (resources.TransactionResponse, error)
	ConfirmTransactionByInvoiceID(purpose, invoiceID string, amount *amount.Fiat) error
	RefundTransaction(id string) (resources.TransactionResponse, error)
	RefundTransactionByInvoiceID(purpose, invoiceID string) error
}

//ErrNotFound - resource not found
var ErrNotFound = errors.New("not found")
var ErrForbidden = errors.New("forbidden")
var ErrNotAuthorized = errors.New("not authorized")

//paymentService - provides access to payment-service REST API
type paymentService struct {
	client  *http.Client
	log     *logan.Entry
	baseURL url.URL
}

const requestsDelay = time.Second * 7

//NewPaymentService - creates new paymentService
func NewPaymentService(baseURL url.URL, log *logan.Entry) PaymentService {

	result := paymentService{
		client:  &http.Client{Timeout: time.Second * 40},
		baseURL: baseURL,
		log:     log,
	}

	return &result
}

func (c *paymentService) get(endpointPath []string, params interface{}) (*http.Response, error) {
	uri := c.baseURL
	uri.Path = path.Join(uri.Path, path.Join(endpointPath...))

	if params != nil {
		query, err := urlval.Encode(params)
		if err != nil {
			panic(errors.Wrap(err, "failed to encode url query"))
		}
		uri.RawQuery = query
	}

	resp, err := c.client.Get(uri.String())
	if err != nil {
		return resp, errors.Wrap(err, "failed to send request")
	}

	return resp, nil
}

func (c *paymentService) post(endpointPath []string, body interface{}) (*http.Response, error) {
	uri := c.baseURL
	uri.Path = path.Join(uri.Path, path.Join(endpointPath...))

	rawBody, err := json.Marshal(body)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode url query"))
	}

	resp, err := c.client.Post(uri.String(), "application/json", bytes.NewReader(rawBody))
	if err != nil {
		return resp, errors.Wrap(err, "failed to send request")
	}

	return resp, nil
}
