package copart

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type ImageList struct {
	FullImages          []Image `json:"FULL_IMAGE"`
	HighResolutionImage []Image `json:"HIGH_RESOLUTION_IMAGE"`
	ThumbNailImages     []Image `json:"THUMBNAIL_IMAGE"`
}

type Image struct {
	URL string `json:"url"`
}

type imageListData struct {
	ImageList ImageList `json:"imagesList"`
}

type imageListResponse struct {
	Data imageListData `json:"data"`
	solrResponse
}

//GetImagesList - returns images for lot
func (c *connector) GetImagesList(lotID string) (*ImageList, error) {
	r, err := http.NewRequest(http.MethodGet, "https://www.copart.com/public/data/lotdetails/solr/lotImages/"+lotID+"/USA", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	var response imageListResponse
	err = c.doJson(r, &response)
	if err != nil {
		if errors.Cause(err) == ErrNotFound {
			return nil, nil
		}
	}

	return &response.Data.ImageList, nil
}