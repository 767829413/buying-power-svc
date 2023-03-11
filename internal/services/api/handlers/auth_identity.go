package handlers

import (
	"gitlab.com/eAuction/events/go/kafka"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/eAuction/bouncer"
	"gitlab.com/eAuction/bouncer/allow"
	"gitlab.com/eAuction/buying-power-svc/internal/data"
)

//GetAuthorizedAccount - checks if requestor is authorized to perform request in the name of identityID
// renders error and returns nil, if not
func GetAuthorizedAccount(r *http.Request, w http.ResponseWriter, identityID string, onlyForAdmin bool) (*data.Identity, bool) {
	identity, err := Storage(r).Identities().Get(identityID)
	if err != nil {
		panic(errors.Wrap(err, "failed to load identity"))
	}

	if identity == nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"account": errors.New("now found"),
		})...)
		return nil, false
	}

	rules := []bouncer.Rule{
		allow.Admin{
			Platform: identity.Platform,
		},
	}

	if !onlyForAdmin {
		brokers, err := Storage(r).IdentitiesToBrokers().GetBrokersForIdentity(identity.ID)
		if err != nil {
			panic(errors.Wrap(err, "failed to get brokers for identity"))
		}

		for _, broker := range brokers {
			rules = append(rules, allow.Broker{
				Platform: identity.Platform,
				Address:  broker,
			})
		}

		rules = append(rules, allow.Account{Address: identity.ID})
	}

	claims, err := Check(r, rules...)
	// non admin accounts must have verified phone number
	if err != nil || claims == nil || (!(allow.IsAdmin(claims) || allow.IsBroker(claims))  && !claims.PhoneVerified) {
		ape.RenderErr(w, problems.NotAllowed())
		return nil, false
	}
	return identity, claims.AccountType == int(kafka.Identity_TYPE_ADMIN) || claims.AccountType == int(kafka.Identity_TYPE_BROKER)
}
