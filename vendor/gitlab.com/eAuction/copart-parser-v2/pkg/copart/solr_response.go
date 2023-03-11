package copart

type solrResponse struct {
	ReturnCode            int64  `json:"returnCode"`
	ReturnCodeDescription string `json:"returnCodeDesc"`
}

//IsSuccess - returns true if request was successful
func (b solrResponse) IsSuccess() bool {
	return b.ReturnCode == 1
}
