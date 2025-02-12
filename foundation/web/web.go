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

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
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

		v := handler(ctx, r)

		data, typ, err := v.Encode()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", typ)
		w.WriteHeader(http.StatusOK)
		w.Write(data)

		// I CAN DO ANYTHING HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
