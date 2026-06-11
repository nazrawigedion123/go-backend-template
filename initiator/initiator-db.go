package initiator

import (
	"context"
	"fmt"
	"time"

	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func initDB(dbSource string, log logger.Logger) *pgxpool.Pool {
	var (
		config *pgxpool.Config
		err    error
	)

	config, err = pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Error(context.Background(), "unable to parse pgxpool config string for onepulse")
		log.Fatal(context.Background(), err.Error())
	}

	// Set idle connection timeout with default fallback
	idleConnTimeout := viper.GetDuration("database.idle_conn_timeout")
	if idleConnTimeout == 0 {
		idleConnTimeout = 4 * time.Minute
	}
	config.MaxConnIdleTime = idleConnTimeout

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("failed to connect to database (%s): %v", dbSource, err))
	}

	log.Info(context.Background(), fmt.Sprintf("connected to %s database successfully", dbSource))
	return conn
}
