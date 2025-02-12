package web

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const (
	traceIDKey ctxKey = 1
)

func setTraceID(ctx context.Context, traceID uuid.UUID) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func GetTraceID(ctx context.Context) uuid.UUID {
	v, ok := ctx.Value(traceIDKey).(uuid.UUID)
	if !ok {
		return uuid.UUID{}
	}

	return v
}
