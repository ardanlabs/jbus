package dbtest

import (
	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/business/domain/userbus/stores/userdb"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// BusDomain represents all the business domain apis needed for testing.
type BusDomain struct {
	User *userbus.Business
}

func newBusDomains(log *logger.Logger, db *sqlx.DB) BusDomain {
	userBus := userbus.NewBusiness(log, userdb.NewStore(log, db))

	return BusDomain{
		User: userBus,
	}
}
