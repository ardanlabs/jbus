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
