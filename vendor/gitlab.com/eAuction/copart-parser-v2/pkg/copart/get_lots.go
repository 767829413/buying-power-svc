package copart

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gocarina/gocsv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type salesLinkResponse struct {
	Data saleDataConfig `json:"saleDataConfig"`
}

func (r *salesLinkResponse) IsSuccess() bool {
	return r.Data.ErrorCode == "" && r.Data.Link != ""
}

type saleDataConfig struct {
	ErrorCode string `json:"errorCode"`
	Link      string `json:"formAction"`
}

//GetLots - returns list of the lots
func (c *connector) GetLots() ([]Lot, error) {
	if c.isAnonymous {
		panic(errors.New("method is only available for non anonymous connector"))
	}
	link, err := c.getSaleDataLink()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sale data link")
	}

	// need another client to override timeouts
	client := http.Client{
		Transport: c.client.Transport,
		Jar:       c.client.Jar,
		Timeout:   time.Minute * 10,
	}

	response, err := client.Get(link)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dowload lots data")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.From(errors.New("unexpected status code"), logan.F{
			"status_code": response.StatusCode,
		})
	}

	lots, err := readLots(c.log, response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read lots")
	}

	return lots, nil
}

// csvReader - main purpose is to ignore wrong number of fields error
type csvReader struct {
	reader *csv.Reader
	log    *logan.Entry
}

func (r *csvReader) Read() (record []string, err error) {
	panic("ReadAll should be used")
}

func (r *csvReader) ReadAll() ([][]string, error) {
	var records [][]string
	for {
		record, err := r.reader.Read()
		if err != nil {
			if err == io.EOF {
				return records, nil
			}

			r.log.WithError(err).Warn("failed to reader record")
			continue
		}

		records = append(records, record)
	}
}

func readLots(log *logan.Entry, reader io.Reader) ([]Lot, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read all reader")
	}

	data = bytes.Replace(data, []byte("High Bid =non-vix,Sealed=Vix"), []byte("High Bid"), 1)

	r := csv.NewReader(bytes.NewReader(data))
	r.LazyQuotes = true
	r.TrimLeadingSpace = true
	var result []Lot
	err = gocsv.UnmarshalCSV(&csvReader{
		reader: r,
		log:    log,
	}, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal csv")
	}

	return result, nil
}

func (c *connector) getSaleDataLink() (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.copart.com/data/sales/config", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	var response salesLinkResponse
	err = c.doJson(req, &response)
	if err != nil {
		return "", errors.Wrap(err, "failed to perform request")
	}

	return response.Data.Link, nil
}
