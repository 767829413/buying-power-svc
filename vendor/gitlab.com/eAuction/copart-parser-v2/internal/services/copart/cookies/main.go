package cookies

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/copart-parser-v2/pkg/resources"
)

//Provider - helper interface to access cookies provided by copart
type Provider interface {
	GetNewCookies(copartLogin, copartPass *string) (*CookieJar, error)
}

// cookiesProvider implements CookiesProvider interface.
type cookiesProvider struct {
	client *http.Client
	url    url.URL
}

//NewProvider - creates new cookies provider
func NewProvider(url url.URL) Provider {
	return &cookiesProvider{
		client: &http.Client{
			Timeout: 2 * time.Minute,
		},
		url: url,
	}
}

// Proxy holds proxy credentials
type Proxy struct {
	URL                *url.URL
	Username, Password *string
}

// CookieJar holds cookies and request params
type CookieJar struct {
	UserAgent string
	CsrfToken string
	Cookies   []*http.Cookie
	Proxy     *Proxy
}

//GetNewCookies - returns new cookies and user agent
func (c *cookiesProvider) GetNewCookies(
	copartLogin,
	copartPass *string,
) (*CookieJar, error) {
	requestBody := resources.SessionRequest{
		Key: resources.Key{
			Type: resources.SESSIONS,
		},
		Attributes: resources.SessionRequestAttributes{
			CopartPassword: copartPass,
			CopartUsername: copartLogin,
		},
	}

	requestBodyAsJSON, err := json.Marshal(resources.SessionRequestResponse{
		Data: requestBody,
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal request body"))
	}

	u := c.url
	u.Path = path.Join(u.Path, "sessions")
	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(requestBodyAsJSON))
	if err != nil {
		panic(errors.Wrap(err, "failed to create request"))
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.From(errors.New("unexpected status code"), logan.F{
			"status_code": resp.StatusCode,
			"body":        string(respBody),
		})
	}

	var response resources.SessionResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body", logan.F{
			"body": string(respBody),
		})
	}

	requiredCookies := map[string]struct{}{
		"copartg2auction": {},
		"incap_ses":       {},
	}
	var cookies []*http.Cookie
	for _, cookieRel := range response.Data.Relationships.Cookies.Data {
		cookie := response.Included.MustCookie(cookieRel.GetKey())
		if cookie == nil {
			return nil, errors.New("expected cookie to be included")
		}
		cookies = append(cookies, &http.Cookie{
			Name:     cookie.Attributes.Name,
			Value:    cookie.Attributes.Value,
			Path:     cookie.Attributes.Path,
			Domain:   cookie.Attributes.Domain,
			Secure:   cookie.Attributes.Secure,
			HttpOnly: cookie.Attributes.HttpOnly,
		})

		for key := range requiredCookies {
			if strings.HasPrefix(cookie.Attributes.Name, key) {
				delete(requiredCookies, key)
			}
		}
	}

	if len(requiredCookies) != 0 {
		return nil, errors.From(errors.New("failed to get all required cookies"), logan.F{
			"missing": requiredCookies,
		})
	}

	proxyInclude := response.Included.MustProxy(response.Data.Relationships.Proxy.Data.GetKey())
	if proxyInclude == nil {
		return nil, errors.New("expected proxy to be included")
	}
	proxyURL, err := url.Parse(proxyInclude.Attributes.Url)
	if err != nil {
		return nil, errors.Wrap(err, "invalid proxy url", logan.F{
			"url": proxyInclude.Attributes.Url,
		})
	}

	proxy := Proxy{
		URL:      proxyURL,
		Username: proxyInclude.Attributes.Username,
		Password: proxyInclude.Attributes.Password,
	}

	jar := CookieJar{
		UserAgent: response.Data.Attributes.UserAgent,
		Cookies:   cookies,
		Proxy:     &proxy,
		CsrfToken: response.Data.Attributes.CsrfToken,
	}

	return &jar, nil

}
