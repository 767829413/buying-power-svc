package allow

import (
	"errors"
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/jwt"
)

//And - defines AND operator
type And []bouncer.Rule

//Check - checks if all rules in the array are satisfied
func (a And) Check(c jwt.Claims) error {
	if len(a) == 0 {
		return errors.New("expected AND operator to contain at least one rule")
	}

	for _, r := range a {
		err := r.Check(c)
		if err != nil {
			return err
		}
	}

	return nil
}

//Or - defines OR operator
type Or []bouncer.Rule

//Check - checks if all rules in the array are satisfied
func (a Or) Check(c jwt.Claims) error {
	if len(a) == 0 {
		return errors.New("expected OR operator to contain at least one rule")
	}

	for _, r := range a {
		err := r.Check(c)
		if err == nil {
			return nil
		}
	}

	return bouncer.ErrNotAllowed
}

//True - check always returns nil
type True struct {}

func (t True) Check(c jwt.Claims) error {
	return nil
}

//False - always returns not allowed
type False struct {}
func (f False) Check(c jwt.Claims) error {
	return bouncer.ErrNotAllowed
}