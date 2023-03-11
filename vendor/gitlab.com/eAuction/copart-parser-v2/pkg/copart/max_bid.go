package copart

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type MaxBiddErr struct {
	msg string
}

func NewMaxBidErr(msg string) MaxBiddErr {
	return MaxBiddErr{msg:msg}
}

func (m MaxBiddErr) Error() string {
	return m.msg
}

//ErrOutBid - signals that max bid has been outbidded
var ErrOutBid = NewMaxBidErr("max bid has been outbidded")
var ErrNotAssigned = NewMaxBidErr("lot has not been assigned to auction")
var ErrPreliminaryBidEnded = NewMaxBidErr("Preliminary bidding has ended")
var ErrBuyerIsNotLicesed = NewMaxBidErr("buyer is not licensed")

type MaxBidLotResult struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorText    string `json:"errorText"`
}

type MaxBidResult struct {
	Lots map[string]MaxBidLotResult `json:"lots"`
}

type MaxBidDataResult struct {
	Result MaxBidResult `json:"result"`
}

type maxBidResponse struct {
	responseMeta
	Data MaxBidDataResult `json:"data"`
}

type MaxBidRequest struct {
	LotID            string `json:"lotId"`
	BidAmount        string `json:"bidAmount"`
	StartingBid      string `json:"startingBid"`
	OneMoreIncrement string `json:"oneMoreIncrement"`
	LotCountry       string `json:"lotCountry"`
	YardNumber       int    `json:"yardNumber"`
	MaxBid           string `json:"maxBid"`
	CurrentBid       int64  `json:"currentBid,omitempty"`
	CurrentMaxBid    int64  `json:"currentMaxBid,omitempty"`
}

func NewMaxBidRequest(lotID, bidAmount, maxBid, lotCountry string, yardNumber int, currentBid, currentMaxBid int64) MaxBidRequest {
	return MaxBidRequest{
		LotID:            lotID,
		BidAmount:        bidAmount,
		StartingBid:      bidAmount,
		OneMoreIncrement: "N",
		LotCountry:       lotCountry,
		YardNumber:       yardNumber,
		MaxBid:           maxBid,
		CurrentBid:       currentBid,
		CurrentMaxBid:    currentMaxBid,
	}
}

func (c *connector) PostMaxBid(request MaxBidRequest) error {
	body, err := json.Marshal([]MaxBidRequest{request})
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}
	req, err := http.NewRequest(http.MethodPost, "https://www.copart.com/data/lots/prelim-bid", bytes.NewReader(body))
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	c.populateHeaders(req)

	var result maxBidResponse
	responseBody, err := c.doJsonWithRawBody(req, &result)
	if err != nil {
		return errors.Wrap(err, "failed to perform request", logan.F{
			"max_bid_request": string(body),
		})
	}

	fields := logan.F{
		"max_bid_result":  string(responseBody),
		"max_bid_request": string(body),
	}

	lotResult, ok := result.Data.Result.Lots[request.LotID]
	if !ok {
		return nil
	}

	err = getTypedError(lotResult.ErrorCode)
	if err != nil {
		return errors.From(err, fields)
	}

	return nil
}

func getTypedError(errorCode string) error {
	switch errorCode {
	case "":
		return nil
	case "5009":
		return ErrOutBid
	case "5002":
		return ErrNotAssigned
	case "5026":
		return ErrPreliminaryBidEnded
	case "C218502":
		return ErrBuyerIsNotLicesed
	default:
		return NewMaxBidErr(fmt.Sprintf("unexpected max bid error: %s", errorCode))
	}
}
