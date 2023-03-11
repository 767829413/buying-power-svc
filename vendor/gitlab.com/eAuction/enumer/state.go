package enumer

import (
	"sync"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/enumer/pkg/maker_model"
	"gitlab.com/eAuction/enumer/pkg/state"
)

var tracker = stateTracker{
	lock: &sync.Mutex{},
}

// InjectStateGetter atomically substitute global state getter by provided one.
// Panics if ModelFinder instantiation fails.
func InjectStateGetter(f state.Getter) {
	err := InjectStateGetterE(f)
	if err != nil {
		panic(err)
	}
}

// InjectStateGetterE does the same as InjectStateGetter but returns error instead of panic.
func InjectStateGetterE(f state.Getter) error {
	return tracker.injectGetter(f)
}

func getState() *extendedState {
	s, err := getStateE()
	if err != nil {
		panic(err)
	}
	return s
}

func getStateE() (*extendedState, error) {
	return tracker.getState()
}

type extendedState struct {
	*state.State
	*maker_model.ModelFinder
}

func newExtendedState(s *state.State) (*extendedState, error) {
	finder, err := maker_model.NewModelFinder(s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate ModelFinder")
	}
	return &extendedState{
		State:       s,
		ModelFinder: finder,
	}, nil
}
