package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

type Logger func(ctx context.Context, msg string, args ...any)

type App struct {
	log Logger
	*http.ServeMux
	mw []MidFunc
}

func NewApp(log Logger, mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

func (a *App) HandleFunc(pattern string, handler HandlerFunc, mw ...MidFunc) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.New())

		resp := handler(ctx, r)

		if err := Respond(ctx, w, resp); err != nil {
			a.log(ctx, "web-respond", "ERROR", err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}
