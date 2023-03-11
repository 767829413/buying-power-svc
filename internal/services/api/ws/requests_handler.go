package ws

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/ws/copart"
	"net/http"
	"time"
)

//CopartWriter - helper interface to scheduleWrite requests to copart
type CopartWriter interface {
	MakeBid(ctx context.Context, requestID string, saleID, lotID string, amount int64) error
	UnsubscribeFromSale(connectionID, saleID string)
	SubscribeForSale(ctx context.Context, connectionID, requestID string, sale copart.Sale)
}

type requestsHandler struct {
	connectionID  string
	getLiveBidder func() CopartWriter
	log           *logan.Entry
	accountID     string
	storage       data.Storage
	isSimulation  bool
}

func newRequestsHandler(log *logan.Entry, connectionID string, accountID string, isSimulation bool,
	storage data.Storage, getLiveBidder func() CopartWriter) *requestsHandler {
	return &requestsHandler{
		getLiveBidder: getLiveBidder,
		log:           log.WithField("handler", "requests_handler"),
		accountID:     accountID,
		connectionID:  connectionID,
		storage:       storage,
		isSimulation:  isSimulation,
	}
}

//HandleRequest - handles request. Returns error in case of non fatal issue. On panic - client connection should be closed
func (h *requestsHandler) HandleRequest(ctx context.Context, request resources.EnvelopeRequestResponse) (resp *resources.EnvelopeResponse) {
	log := h.log.WithField("request_id", request.Data.ID)
	log.Debug("got new request to handle")
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	resp, err := h.handleRequest(ctx, request)
	if err != nil {
		log.WithError(err).Error("failed to handle request")
		return responses.PopulateError(request.Data.ID, http.StatusInternalServerError, errors.New("server error"))
	}

	log.Debug("handled request")
	return resp
}

func (h *requestsHandler) handleRequest(ctx context.Context, request resources.EnvelopeRequestResponse) (*resources.EnvelopeResponse, error) {
	if request.Data.Relationships.Subscribe != nil {
		return h.handleSubscribe(ctx, request)
	}

	if request.Data.Relationships.Authorize != nil {
		return responses.PopulateError(request.Data.ID, http.StatusConflict, errors.New("already authorized")), nil
	}

	if request.Data.Relationships.Bid != nil {
		return h.handleBid(ctx, request)
	}

	if request.Data.Relationships.Unsubscribe != nil {
		return h.handleUnsubscribe(ctx, request)
	}

	return responses.PopulateError(request.Data.ID, http.StatusMethodNotAllowed, errors.New("unknown command")), nil
}

func (h *requestsHandler) handleUnsubscribe(_ context.Context, request resources.EnvelopeRequestResponse) (*resources.EnvelopeResponse, error) {
	obj := request.Data.Relationships.Unsubscribe.Data
	if obj == nil {
		return responses.PopulateError(request.Data.ID, http.StatusBadRequest, errors.New("expected non nil UnSubscribe data")), nil
	}
	saleID, err := data.ExtractCopartSaleID(obj.ID)
	if err != nil {
		return responses.PopulateError(request.Data.ID, http.StatusBadRequest, err), nil
	}
	h.getLiveBidder().UnsubscribeFromSale(h.connectionID, saleID)
	return responses.PopulateOk(request.Data.ID), nil
}

func (h *requestsHandler) handleSubscribe(ctx context.Context, request resources.EnvelopeRequestResponse) (*resources.EnvelopeResponse, error) {
	obj := request.Data.Relationships.Subscribe.Data
	if obj == nil {
		panic(errors.New("expected subscribe data validation to be performed by caller of the method"))
	}

	log := h.log.WithField("request_id", request.Data.ID)

	sale, err := h.storage.Sales().GetSaleByID(obj.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sale")
	}

	if sale == nil {
		resp := responses.PopulateSaleEnd(obj.ID, "not found")
		return &resp, nil
	}

	if h.isSimulation {
		resp := responses.PopulateSaleEnd(obj.ID, "simualtion is disabled for now")
		return &resp, nil
	}

	if !h.isSimulation && time.Now().Add(time.Minute * 40).UTC().Before(sale.StartsAt) {
		resp := responses.PopulateSaleEnd(obj.ID, "too early")
		return &resp, nil
	}

	if !h.isSimulation &&  sale.StartsAt.Add(4 * time.Hour).Before(time.Now().UTC()) {
		resp := responses.PopulateSaleEnd(obj.ID, "too late")
		return &resp, nil
	}

	copartSaleID, err := data.ExtractCopartSaleID(obj.ID)
	if err != nil {
		resp := responses.PopulateSaleEnd(obj.ID, err.Error())
		return &resp, nil
	}

	// subscribe might take some time to complete, to ensure that other requests are processed need to run it in separate go routine
	go func() {
		// TODO: find better way to handle issue with context been canceled too early
		subscribeCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err= running.RunSafely(subscribeCtx, "sale_subscribed", func(ctx context.Context) error {
			log.Debug("sending subscribe request")
			h.getLiveBidder().SubscribeForSale(ctx, h.connectionID, request.Data.ID, copart.Sale{
				ShelfSaleID:  sale.ID,
				CopartSaleID: copartSaleID,
			})
			log.Debug("subscribed for sale")
			return nil
		})
		if err != nil {
			log.WithError(err).Error("failed to subscribe to sale")
		}
	}()

	return nil, nil
}
