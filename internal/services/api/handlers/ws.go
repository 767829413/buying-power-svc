package handlers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"strings"
	"time"

	"gitlab.com/distributed_lab/running"
	"gitlab.com/eAuction/bouncer/allow"
	"gitlab.com/eAuction/bouncer/jwt"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/resources"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/responses"
	"gitlab.com/eAuction/buying-power-svc/internal/services/api/ws"

	"github.com/gorilla/websocket"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//WS - allows to establish socket connection
func WS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		HandshakeTimeout: 30 * time.Second,
		// TODO: check if it's ok
		CheckOrigin:       func(*http.Request) bool { return true },
		EnableCompression: true,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// upgrader has already responded with error code
		Log(r).WithError(err).Warn("failed to upgrade connection to ws")
		return
	}

	err = conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		Log(r).WithError(err).Warn("failed to set write deadline")
		_ = conn.Close()
		return
	}
	requestID, claims, err := authenticate(r, conn)
	if err != nil {
		Log(r).WithError(err).Debug("authentication failed")
		_ = conn.WriteJSON(responses.PopulateError(requestID, http.StatusUnauthorized, err))
		_ = conn.Close()
		return
	}

	err = conn.WriteJSON(responses.PopulateOk(requestID))
	if err != nil {
		Log(r).WithError(err).Warn("failed to write ok on auth")
		_ = conn.Close()
		return
	}

	GetRunTimeInfo(r).AddConnection(claims.AccountID)
	defer func() {
		GetRunTimeInfo(r).RemoveConnection(claims.AccountID)
	}()

	ws.RunClient(r.Context(), conn, claims, Config(r), !strings.Contains(requestID, "production"))
}

func authenticate(r *http.Request, conn *websocket.Conn) (string, jwt.Claims, error) {
	var claims *jwt.Claims
	var requestID string
	err := running.RunSafely(context.Background(), "ws-auth", func(ctx context.Context) error {
		err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			return errors.Wrap(err, "failed to set read deadline")
		}

		var request resources.EnvelopeRequestResponse
		err = conn.ReadJSON(&request)
		if err != nil {
			return errors.Wrap(err, "failed to read auth request")
		}

		requestID = request.Data.ID

		tokenKey := request.Data.Relationships.Authorize
		if tokenKey == nil || tokenKey.Data == nil {
			return errors.New("expected token as first message")
		}

		token := request.Included.MustToken(tokenKey.Data.GetKey())
		if token == nil {
			return errors.New("expected token to be included")
		}

		r.Header.Add("Authorization", token.Attributes.Token)
		claims, err = Check(r, allow.Account{})
		if err != nil {
			return errors.Wrap(err, "failed to check claims")
		}

		if claims == nil {
			return errors.New("expected claims to be present")
		}

		platform, err := Config(r).Platformer().GetPlatform(claims.PlatformID)
		if err != nil {
			return errors.Wrap(err, "failed to get platform")
		}

		if platform == nil {
			return errors.From(errors.New("platform does not exist"), logan.F{
				"platform_id": claims.PlatformID,
			})
		}

		return nil
	})
	if err != nil {
		return requestID, jwt.Claims{}, errors.Wrap(err, "authentication failed")
	}

	return requestID, *claims, nil
}
