package copart

import (
	"compress/gzip"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type baseResponse interface {
	IsSuccess() bool
}

func (c *connector) do(req *http.Request) (io.ReadCloser, error) {
	select {
	case <-c.throttle.C:
	case <-c.Done():
		return nil, errors.Wrap(c.ctx.Err(), "connector context has been closed")
	}

	defer c.throttle.Reset(requestsDelay)
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	if resp.StatusCode == http.StatusForbidden {
		c.Close()
		return nil, errors.New("received 403 error from copart - closing connector")
	}

	if resp.StatusCode >= 400 {
		return nil, errors.Errorf("request failed with code %d", resp.StatusCode)
	}

	var bodyReader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		bodyReader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get gzip reader")
		}
	default:
		bodyReader = resp.Body
	}

	return bodyReader, nil
}

func (c *connector) doJsonWithRawBody(req *http.Request, response baseResponse) ([]byte, error) {
	bodyReader, err := c.do(req)
	defer func() {
		if bodyReader != nil {
			_ = bodyReader.Close()
		}
	}()
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}

	fields := logan.F{
		"raw_body": string(body),
	}

	if response == nil {
		return body, nil
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		if strings.Contains(string(body), "Incapsula") {
			c.log.WithFields(fields).Error("received incapsula error - closing connector")
			c.Close()
		}

		if strings.Contains(string(body), "<!DOCTYPE html>") {
			c.log.WithFields(fields).Error("received html while expecting json - closing connector")
			c.Close()
		}
		return body, errors.Wrap(err, "failed to parse body", fields)
	}

	if !response.IsSuccess() {
		return body, errors.From(errors.New("request failed"), fields)
	}

	return body, nil
}

//doJson - performs request
func (c *connector) doJson(req *http.Request, response baseResponse) error {
	_, err := c.doJsonWithRawBody(req, response)
	return err
}
