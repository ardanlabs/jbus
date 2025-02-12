package mux

import (
	"context"

	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI(log *logger.Logger) *web.App {
	l := func(ctx context.Context, msg string, args ...any) {
		log.Info(ctx, msg, args...)
	}

	app := web.NewApp(l, mid.Logger(log), mid.Error(log))

	testapp.Routes(app)

	return app
}
