package ws

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/events/go/kafka"
	"net/http"
	"strconv"
	"strings"
)

func (h *requestsHandler) handleBid(ctx context.Context, request resources.EnvelopeRequestResponse) (*resources.EnvelopeResponse, error) {
	obj := request.Data.Relationships.Bid.Data
	if obj == nil {
		return responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("expected non nil Bid data")), nil
	}

	bid := request.Included.MustLiveBid(obj.GetKey())
	if bid == nil {
		return responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("expected live bid to be included")), nil
	}

	if bid.Relationships.Lot.Data == nil {
		return responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("expected lot id to be specified")), nil
	}

	lot, err := h.storage.Lots().ByID(bid.Relationships.Lot.Data.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get lot by id")
	}

	if lot == nil {
		return responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("lot not found")), nil
	}

	amount := int64(bid.Attributes.Amount)

	txStorage := h.storage.Clone()
	var resp *resources.EnvelopeResponse
	err = txStorage.Transaction(func() error {
		participant, err := txStorage.Participants().GetForUpdate(bid.Relationships.Lot.Data.ID, h.accountID)
		if err != nil {
			return errors.Wrap(err, "failed to get participant")
		}

		if participant == nil || participant.State != int32(kafka.Participant_STATE_PENDING) {
			resp = responses.PopulateError(request.Data.ID, http.StatusMethodNotAllowed, errors.New("participant must exist in pending state"))
			return nil
		}

		if participant.CurrentBid > amount || participant.RequestedBid > amount {
			resp = responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("amount must be greater current bid and requested bid"))
			return nil
		}

		if participant.BidLimit < amount {
			resp = responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("amount must not exceed bidding limit"))
			return nil
		}

		err = txStorage.Participants().SetRequestedBid(participant.ID, amount)
		if err != nil {
			return errors.Wrap(err, "failed to set requested bid")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to commit tx")
	}

	if resp != nil {
		return resp, nil
	}

	log := h.log.WithFields(logan.F{
		"shelf_request_id":  request.Data.ID,
	})
	log.Debug("sending bid to copart")
	copartSaleID, err := data.ExtractCopartSaleID(lot.SaleID.String)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract copart sale id")
	}

	copartLotID, err := getLotID(*lot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get lot id")
	}

	err = h.getLiveBidder().MakeBid(ctx, request.Data.ID, copartSaleID, strconv.FormatInt(copartLotID, 10),amount/100 )
	if err != nil {
		return nil, errors.Wrap(err, "failed to make bid")
	}
	log.Debug("bid has been accepted")

	return responses.PopulateOk(request.Data.ID), nil
}

func getLotID(lot data.Lot) (int64, error) {
	parts := strings.Split(lot.ExternalID.String, "-")
	if len(parts) != 3 {
		return 0, errors.From(errors.New("unexpected format of external ID"), logan.F{
			"ex_id": lot.ExternalID,
		})
	}

	lotID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse lot id", logan.F{
			"ex_id": lot.ExternalID,
		})
	}

	return lotID, nil
}
