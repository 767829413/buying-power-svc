package remote

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/enumer/pkg/state"
)

// NewGetter returns remote server state getter.
func NewGetter(endpoint url.URL, failOnUpdate bool) (state.Getter, error) {
	enumer := &getter{
		endpoint:     endpoint,
		failOnUpdate: failOnUpdate,
		httpclient: &http.Client{
			Timeout: 2 * time.Minute,
		},
	}
	return enumer, enumer.refresh()
}

type getter struct {
	endpoint     url.URL
	failOnUpdate bool
	httpclient   *http.Client

	// TODO: add hash of the current state to not reparse and
	//  so not rebuild enums if the state is actually unchanged.
	//  Or add timestamp for server response and compare it to local one.
	cached   *state.State
	expireAt time.Time
	lock     sync.RWMutex
}

// GetState returns cached or newly updated state.
// Period between updates is 1 hour for both normal and failed update.
func (g *getter) GetState() (*state.State, error) {
	err := g.refresh()
	if err != nil && (g.failOnUpdate || g.cached == nil) {
		return nil, err
	}

	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.cached, nil
}

func (g *getter) refresh() error {
	// Need an exclusive lock to not let other routines a chance to start a refresh.
	g.lock.Lock()
	defer g.lock.Unlock()

	expired := g.expireAt.Before(time.Now())
	if !expired {
		return nil
	}

	cached, err := g.loadState()
	if err != nil {
		return errors.Wrap(err, "failed to load state")
	}

	g.cached = cached
	g.expireAt = time.Now().Add(1 * time.Hour)
	return nil
}

func (g *getter) loadState() (*state.State, error) {
	u := g.endpoint
	u.Path = path.Join(u.Path, "state")

	response, err := g.httpclient.Get(u.String())
	if response != nil {
		defer response.Body.Close()
		defer io.Copy(ioutil.Discard, response.Body)
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get state")
	}

	responseBB, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var state state.State
	if err = json.Unmarshal(responseBB, &state); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	return &state, nil
}
