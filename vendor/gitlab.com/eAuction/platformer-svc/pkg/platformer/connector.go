package platformer

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/platformer-svc/resources"
)

// connector - helper struct to communicate with platformer-svc.
type connector struct {
	client  *http.Client
	baseURL *url.URL

	cachedResponse resources.PlatformListResponse
	createdAt      time.Time
}

// NewPlatformer - create new instance of thread safe connector for platfomer-svc
func NewPlatformer(baseURL *url.URL) *connector {
	return &connector{
		client: &http.Client{
			Timeout: 40 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (c *connector) get(u string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	return c.client.Do(request)
}

// GetPlatformList returns a list of all platforms service is able to provide.
func (c *connector) GetPlatformList() (*resources.PlatformListResponse, error) {
	if c.createdAt.Add(10*time.Minute).After(time.Now()) {
		return &c.cachedResponse, nil
	}

	uri := *c.baseURL
	uri.Path += "/platforms"
	uri.RawQuery = "include=fees,shelf_fee,transportation_fee,indirect_fees,liontrans_fees"

	resp, err := c.get(uri.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.From(errors.New("expected 200 status code"), logan.F{
			"status_code": resp.StatusCode,
		})
	}

	var respBody resources.PlatformListResponse
	if err = json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	c.cachedResponse = respBody
	c.createdAt = time.Now()

	return &c.cachedResponse, nil
}

// MustPlatformList returns a list of all platforms service is able to provide. Panics if got error.
func (c *connector) MustPlatformList() *resources.PlatformListResponse {
	platforms, err := c.GetPlatformList()
	if err != nil {
		panic(err)
	}

	return platforms
}

// GetPlatform returns a specific platform by it's ID. Returns nil, nil
// if platform not found.
func (c *connector) GetPlatform(id string) (*resources.PlatformResponse, error) {
	platforms, err := c.GetPlatformList()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get platforms")
	}

	for _, platform := range platforms.Data {
		if platform.ID == id {
			return &resources.PlatformResponse{
				Data:     platform,
				Included: platforms.Included,
			}, nil
		}
	}

	return nil, nil
}

// MustPlatform returns a specific platform by it's ID. Returns nil, nil
// if platform not found. Panics if got other that 404 error.
func (c *connector) MustPlatform(id string) *resources.PlatformResponse {
	platform, err := c.GetPlatform(id)
	if err != nil {
		panic(err)
	}

	return platform
}

// GetPlatformIDs returns list of just platform IDs, without any details.
func (c *connector) GetPlatformIDs() ([]string, error) {
	platforms, err := c.GetPlatformList()
	if err != nil {
		return nil, err
	}

	var result []string
	for _, p := range platforms.Data {
		result = append(result, p.ID)
	}

	return result, nil
}

// MustPlatformIDs returns list of just platform IDs, without any details. Panics if got error.
func (c *connector) MustPlatformIDs() []string {
	result, err := c.GetPlatformIDs()
	if err != nil {
		panic(err)
	}

	return result
}

//GetPlatformByCode - returns platform by code
func (c *connector) GetPlatformByCode(code string) (*resources.PlatformResponse, error) {
	platforms, err := c.GetPlatformList()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get platforms")
	}

	code = strings.ToUpper(code)
	for _, platform := range platforms.Data {
		if strings.ToUpper(platform.Attributes.Code) == code {
			return &resources.PlatformResponse{
				Data:     platform,
				Included: platforms.Included,
			}, nil
		}
	}

	return nil, nil
}
