package testapp

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func test(ctx context.Context, r *http.Request) web.Encoder {
	if n := rand.Intn(100); n%2 == 0 {
		//return errs.Newf(errs.FailedPrecondition, "randon error")
		panic("OHHHH NOOOOO")
	}

	s := status{
		Status: "OK",
	}

	return s
}
