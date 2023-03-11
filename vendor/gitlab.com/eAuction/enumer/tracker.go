package enumer

import (
	"errors"
	"os"
	"strings"
	"sync"

	"gitlab.com/eAuction/enumer/pkg/state"
)

type stateTracker struct {
	lock *sync.Mutex

	getter        state.Getter
	originalState *state.State
	extendedState *extendedState
}

func (s *stateTracker) injectGetter(getter state.Getter) error {
	s.setGetter(getter)
	_, err := s.getState()
	return err
}

func isInTests() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") || strings.HasSuffix(arg, ".test") {
			return true
		}
	}
	return false
}

func (s *stateTracker) setGetter(getter state.Getter) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.getter == nil {
		s.getter = getter
	} else if !isInTests() {
		panic(errors.New("getter should be injected once"))
	}
}

func (s *stateTracker) getState() (*extendedState, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.getter == nil {
		panic("getter is null")
	}
	newState, err := s.getter.GetState()
	if err != nil {
		return nil, err
	}
	// TODO: come up with better caching strategy
	if newState == s.originalState {
		return s.extendedState, nil
	}
	newExtState, err := newExtendedState(newState)
	if err != nil {
		return nil, err
	}

	s.extendedState = newExtState
	s.originalState = newState
	return newExtState, nil
}
