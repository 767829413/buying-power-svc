package payments

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/go/amount"
	"gitlab.com/eAuction/payment-service/pkg/payments/resources"
)

//ConfirmTransaction and charge the money
func (c *paymentService) ConfirmTransaction(id string, amount *amount.Fiat) (resources.TransactionResponse, error) {
	body := resources.PostConfirmationResponse{Data: resources.PostConfirmation{
		Key: resources.Key{ID: id, Type: resources.TRANSACTIONS},
	}}
	if amount != nil {
		body.Data.Attributes = &resources.PostConfirmationAttributes{
			Amount: *amount,
		}
	}

	resp, err := c.post([]string{"confirmations"}, body)
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
	case http.StatusForbidden:
		return resources.TransactionResponse{}, ErrForbidden
	default:
		respBB, _ := ioutil.ReadAll(resp.Body)
		return resources.TransactionResponse{}, errors.From(fmt.Errorf("unexpected status code: %d", resp.StatusCode), logan.F{
			"url":          resp.Request.RequestURI,
			"raw_response": string(respBB),
			"status_code":  resp.StatusCode,
		})
	}
}

//RefundTransaction
func (c *paymentService) RefundTransaction(id string) (resources.TransactionResponse, error) {
	body := resources.PostRefundResponse{Data: resources.PostRefund{
		Key: resources.Key{ID: id, Type: resources.TRANSACTIONS},
	}}
	resp, err := c.post([]string{"refunds"}, body)

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
	case http.StatusForbidden:
		return resources.TransactionResponse{}, ErrForbidden
	default:
		respBB, _ := ioutil.ReadAll(resp.Body)
		return resources.TransactionResponse{}, errors.From(fmt.Errorf("unexpected status code: %d", resp.StatusCode), logan.F{
			"url":          resp.Request.RequestURI,
			"raw_response": string(respBB),
			"status_code":  resp.StatusCode,
		})
	}
}

//ConfirmTransactionByInvoiceID tries to find a transaction by invoice ID and charge it
func (p *paymentService) ConfirmTransactionByInvoiceID(purpose, invoiceID string, amount *amount.Fiat) error {
	statuses := strings.Join([]string{TxStatusOk, TxStatusConfirmed}, ",")
	transactions, err := p.GetTransactions(GetTransactionsParams{
		InvoiceID: &invoiceID,
		Purpose:   &purpose,
		Status:    &statuses,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get tx id")
	}
	if len(transactions.Data) == 0 {
		return errors.Wrap(ErrNotFound, fmt.Sprintf("matching transaction is not found for invoice_id = %s", invoiceID))
	}

	var activeTx *resources.Transaction
	for _, tx := range transactions.Data {
		if tx.Attributes.Status == TxStatusOk {
			if activeTx != nil {
				p.log.WithFields(logan.F{"invoice_id": invoiceID, "tx_id": tx.ID}).Warnf("found multiple active payments")
			} else {
				activeTx = &tx
			}
		} else {
			p.log.WithFields(logan.F{"invoice_id": invoiceID, "tx_id": tx.ID}).Warnf("found an already confirmed payment")
		}
	}

	if activeTx == nil {
		p.log.WithFields(logan.F{"invoice_id": invoiceID}).Warn("payment was processed already")
		return nil
	}

	_, err = p.ConfirmTransaction(activeTx.ID, amount)
	if err == ErrForbidden {
		p.log.WithFields(logan.F{"invoice_id": invoiceID}).Warn("payment was processed already")
	} else if err != nil {
		return errors.Wrap(err, "failed to confirm tx", logan.F{"tx_id": activeTx.ID, "invoice_id": invoiceID})
	}

	return nil
}

//ConfirmTransactionByInvoiceID tries to find a transaction by invoice ID and refund it
func (p *paymentService) RefundTransactionByInvoiceID(purpose, invoiceID string) error {
	statuses := strings.Join([]string{TxStatusOk, TxStatusConfirmed, TxStatusReversed, TxStatusRefunded}, ",")
	transactions, err := p.GetTransactions(GetTransactionsParams{
		InvoiceID: &invoiceID,
		Purpose:   &purpose,
		Status:    &statuses,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get tx id")
	}
	if len(transactions.Data) == 0 {
		return errors.Wrap(ErrNotFound, fmt.Sprintf("matching transaction is not found for invoice_id = %s", invoiceID))
	}

	var activeTx *resources.Transaction
	for _, tx := range transactions.Data {
		if tx.Attributes.Status == TxStatusReversed || tx.Attributes.Status == TxStatusRefunded {
			p.log.WithFields(logan.F{"invoice_id": invoiceID, "tx_id": tx.ID}).Warnf("found an already refunded transaction")
		} else {
			if activeTx != nil {
				p.log.WithFields(logan.F{"invoice_id": invoiceID, "tx_id": tx.ID}).Warnf("found multiple active transactions")
			} else {
				activeTx = &tx
			}
		}
	}

	if activeTx == nil {
		p.log.WithFields(logan.F{"invoice_id": invoiceID}).Warn("transaction was processed already")
		return nil
	}

	_, err = p.RefundTransaction(activeTx.ID)
	if err == ErrForbidden {
		p.log.WithFields(logan.F{"invoice_id": invoiceID}).Warn("transaction was processed already")
	} else if err != nil {
		return errors.Wrap(err, "failed to confirm tx", logan.F{"tx_id": activeTx.ID, "invoice_id": invoiceID})
	}

	return nil
}
