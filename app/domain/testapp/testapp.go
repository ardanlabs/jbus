package testapp

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func test(ctx context.Context, r *http.Request) web.Encoder {
	s := status{
		Status: "OK",
	}

	return s
}
