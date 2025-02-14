package mux

import (
	"context"

	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/app/domain/userapp"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/business/domain/userbus/stores/userdb"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

func WebAPI(log *logger.Logger, db *sqlx.DB) *web.App {
	l := func(ctx context.Context, msg string, args ...any) {
		log.Info(ctx, msg, args...)
	}

	app := web.NewApp(
		l,
		mid.Logger(log),
		mid.Error(log),
		mid.Panic(),
	)

	// -------------------------------------------------------------------------

	userBus := userbus.NewBusiness(log, userdb.NewStore(log, db))

	// -------------------------------------------------------------------------

	testapp.Routes(app)

	userapp.Routes(app, userapp.Config{
		Log:     log,
		UserBus: userBus,
	})

	return app
}
