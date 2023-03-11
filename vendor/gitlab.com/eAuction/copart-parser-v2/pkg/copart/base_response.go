package copart

//g2AuctionResponseMeta - defines meta of the response
type g2AuctionResponseMeta struct {
	Meta responseMetaBody `json:"metadata"`
}

//responseMetaBody - defines body of the meta
type responseMetaBody struct {
	// fucking copart
	StatusCode      int    `json:"statusCode"`
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	StatusMessage   string `json:"statusMessage"`
}

//IsSuccess - returns true if request was successful
func (b g2AuctionResponseMeta) IsSuccess() bool {
	return b.Meta.StatusCode == 200 || b.Meta.ResponseCode == 200
}
