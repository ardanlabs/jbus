package web

import (
	"context"
	"net/http"
)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

func (app *App) HandleFunc(pattern string, handler HandlerFunc) {

	h := func(w http.ResponseWriter, r *http.Request) {

		// I CAN DO ANYTHING HERE

		v := handler(r.Context(), r)

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

	app.ServeMux.HandleFunc(pattern, h)
}
