package persistencedb

import (
	"github.com/OnePulseOmni/pulse-wallet/internal/constant/model/db"
	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PersistenceDB struct {
	*db.Queries
	Pool *pgxpool.Pool
	log  logger.Logger
}

type Sibling string

func New(pool *pgxpool.Pool, log logger.Logger) PersistenceDB {
	return PersistenceDB{
		Queries: db.New(pool),
		Pool:    pool,
		log:     log,
	}
}
