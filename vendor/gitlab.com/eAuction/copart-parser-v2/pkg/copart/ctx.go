package copart

import "context"

type ctxKey int

const (
	requestIDCtxKey ctxKey = iota
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDCtxKey, requestID)
}

func RequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDCtxKey).(string)
	if !ok {
		requestID = "not_specified"
	}

	return requestID
}
