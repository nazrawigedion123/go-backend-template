package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	dbinterface "github.com/nazrawigedion123/go-backend-template/internal/constant/db/db_interface"
	"github.com/nazrawigedion123/go-backend-template/internal/handler/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiate() {

	// docs.SwaggerInfo.Title = "PulseWallet API"
	// docs.SwaggerInfo.Description = "API documentation for PulseWallet"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.BasePath = "/"

	ctx := context.Background()

	log, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to start logging")
	}
	configName := "config"
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}
	err = InitConfig(Config{Names: []string{configName}, Path: "config", Logger: log})
	if err != nil {
		log.Fatal("unable to start config", zap.Error(err))
	}

	log.Info("initializing config completed")

	logger := InitLogger()
	log.Info("initializing logger completed")

	// initailizing database connection
	log.Info("initializing database connect")
	dbURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	)
	pgxPool := initDB(dbURL, logger)
	log.Info("database connection initialized")

	// // initializing migration
	
	// initializing persistence layer which is responsible to communicate with the database and module layer
	// which is used as middleware between database and module layer of the application

	logger.Info(ctx, "initializing persistence layer ")
	persistenceDB := dbinterface.New(pgxPool, logger)
	persistence := initPersistence(&persistenceDB, logger)
	logger.Info(ctx, "done initializing persistence layer")
	logger.Info(ctx, "initializing client layer")

	// initialize module

	logger.Info(ctx, "initializing module layer")
	module := initModule(persistence, logger)
	logger.Info(ctx, "done initializing module layer")

	// initializing handler layer
	// which is the layer responsible to handle http layer and validate user
	logger.Info(ctx, "initializing handler layer ")
	handler := initHandler(module, logger)
	logger.Info(ctx, "done initializing handler layer")

	logger.Info(ctx, "initializing http server")
	server := gin.New()
	server.Use(middleware.GinLogger(logger))
	server.Use(middleware.CORS())
	ginsrv := server.Group("")
	// server.Use(middleware.ErrorHandler())

	// initializing route which handle route endpoints
	logger.Info(ctx, "initializing route")
	initRoute(ginsrv, handler, logger)
	logger.Info(ctx, "done initializing route")

	logger.Info(ctx, "done initializing server")

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", viper.GetString("app.host"), viper.GetInt("app.port")),
		Handler:           server,
		ReadHeaderTimeout: viper.GetDuration("app.timeout"),
		IdleTimeout:       30 * time.Minute,
	}

	host := fmt.Sprint(viper.GetString("app.host"), ":", viper.GetInt("app.port"))
	logger.Info(ctx, "server listening at port ", zap.Any("link", host))
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(ctx, fmt.Sprintf("Could not start HTTP server: %s", err))
	}

}
