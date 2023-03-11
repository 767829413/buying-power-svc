package bouncer

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"strings"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gitlab.com/eAuction/bouncer/jwt"
)

var (
	ErrNotAllowed   = &notallowed{"not allowed"}
	ErrInvalidToken = &badrequest{"invalid token"}
	ErrNotSigned    = &notallowed{"not signed"}
)

const (
	BlacklistHeader = "X-Blacklist"
	GatewayHeader   = "X-Gateway"
)

const BlacklistStatusAccountPaused = "PAUSED"

func GatewayPassed(r *http.Request) bool {
	return r.Header.Get(GatewayHeader) != ""
}

// The RuleFunc type is an adapter to allow the use of
// ordinary functions as Rule
type RuleFunc func(jwt.Claims) error

func (r RuleFunc) Check(claims jwt.Claims) error {
	return r(claims)
}

type Rule interface {
	Check(jwt.Claims) error
}

type Bouncer interface {
	Check(r *http.Request, rules ...Rule) (*jwt.Claims, error)
}

type bouncer struct {
	opts Opts
}

type Opts struct {
	// SkipChecks make any request with valid or missing token pass
	SkipChecks  bool
	AllowPaused bool
}

func New(opts Opts) Bouncer {
	return &bouncer{opts}
}

func getToken(r *http.Request) (*string, error) {
	v := r.Header.Get("Authorization")
	if v == "" {
		cookie, err := r.Cookie("Authorization")
		// it is the only error returned from the method above
		if err == http.ErrNoCookie {
			return nil, nil
		}

		v = cookie.Value
	}

	token := strings.TrimPrefix(v, "Bearer ")
	if token == "" {
		return nil, ErrInvalidToken
	}

	return &token, nil
}

func ParseWithClaims(r *http.Request, claims *jwt.Claims) error {
	rawToken, err := getToken(r)
	if err != nil {
		return err
	}

	if rawToken == nil {
		return ErrNotSigned
	}

	var parser jwtgo.Parser
	_, _, err = parser.ParseUnverified(*rawToken, claims)
	if err != nil {
		return ErrInvalidToken
	}

	if r.Header.Get(BlacklistHeader) == BlacklistStatusAccountPaused {
		claims.AccountPaused = true
	}

	return nil
}

func ValidateWithClaims(r *http.Request, key *ecdsa.PublicKey, claims jwtgo.Claims) error {
	rawToken, err := getToken(r)
	if err != nil {
		return err
	}

	if rawToken == nil {
		return ErrNotSigned
	}

	_, err = jwtgo.ParseWithClaims(*rawToken, claims, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return ErrInvalidToken
	}

	return nil
}

func (c bouncer) Check(r *http.Request, rules ...Rule) (*jwt.Claims, error) {
	var claims jwt.Claims

	err := ParseWithClaims(r, &claims)
	if err != nil {
		if err == ErrNotSigned && c.opts.SkipChecks {
			// unsigned requests are allowed with SkipChecks,
			// beware of nil pointers tho'
			return nil, nil
		}

		return nil, err
	}

	if claims.AccountPaused && !c.opts.AllowPaused {
		return nil, ErrNotAllowed
	}

	// authorize any request with token if no rules specified, or told to do so
	if len(rules) == 0 || c.opts.SkipChecks {
		return &claims, nil
	}

	for _, rule := range rules {
		switch err := rule.Check(claims); err {
		case nil:
			return &claims, nil
		case ErrNotAllowed:
			continue
		default:
			return nil, errors.Wrap(err, "check failed")
		}
	}

	return nil, ErrNotAllowed
}
