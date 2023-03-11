package copart

//responseMeta - defines details of the generic response
type responseMeta struct {
	ReturnCode     int    `json:"returnCode"`
	ReturnCodeDesc string `json:"returnCodeDesc"`
}

//IsSuccess - returns true if copart internal code signals success
func (m responseMeta) IsSuccess() bool {
	return m.ReturnCode == 1
}