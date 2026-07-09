package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/nazrawigedion123/go-backend-template/internal/constant/db/dbinterface"
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

	// 1. Config Loading
	fmt.Printf("\n%s🚀 Starting %sInfrastructure...%s\n", colorBold+colorCyan, viper.GetString("app.name"), colorReset)
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("⚙️  Loading configurations... ")

	configName := "config"
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	err = InitConfig(Config{Names: []string{configName}, Path: "config", Logger: log})
	if err != nil {
		log.Fatal("unable to start config", zap.Error(err))
	}
	fmt.Printf("%s[DONE]%s\n", colorGreen, colorReset)

	// log.Info("initializing config completed")

	logger := InitLogger()
	// log.Info("initializing logger completed")

	// initailizing database connection
	// log.Info("initializing database connect")
	fmt.Printf("🔌  Connecting to database... ")

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
	// log.Info("database connection initialized")
	fmt.Printf("%s[DONE]%s\n", colorGreen, colorReset)
	// // initializing migration

	// initializing persistence layer which is responsible to communicate with the database and module layer
	// which is used as middleware between database and module layer of the application

	// logger.Info(ctx, "initializing persistence layer ")
	fmt.Printf("💾  Initializing persistence layer... ")
	persistenceDB := dbinterface.New(pgxPool, logger)
	persistence := initPersistence(&persistenceDB, logger)
	fmt.Printf("%s[DONE]%s\n", colorGreen, colorReset)
	// logger.Info(ctx, "initializing client layer")

	// initialize module
	fmt.Printf("🔧  Initializing module layer... ")
	module := initModule(persistence, logger)
	fmt.Printf("%s[DONE]%s\n", colorGreen, colorReset)

	// initializing handler layer
	// which is the layer responsible to handle http layer and validate user
	// logger.Info(ctx, "initializing handler layer ")
	fmt.Printf("🔧  Initializing handler layer... ")
	handler := initHandler(module, logger)
	// logger.Info(ctx, "done initializing handler layer")
	fmt.Printf("%s[DONE]%s\n", colorGreen, colorReset)

	fmt.Printf("🔧  Initializing http server... ")
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(middleware.GinLogger(logger))
	server.Use(middleware.CORS())
	ginsrv := server.Group("")
	// server.Use(middleware.ErrorHandler())

	// initializing route which handle route endpoints
	fmt.Printf("🔧  Initializing route... ")
	initRoute(ginsrv, handler, logger)
	fmt.Printf("%s[DONE]%s\n", colorGreen, colorReset)
	printPrettyRoutes(server)

	fmt.Printf("🔧  Initializing server... ")
	addr :=  fmt.Sprintf("%s:%d", viper.GetString("app.host"), viper.GetInt("app.port"))
	srv := &http.Server{
		Addr:            addr ,
		Handler:           server,
		ReadHeaderTimeout: viper.GetDuration("app.timeout"),
		IdleTimeout:       30 * time.Minute,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)


	// logger.Info(ctx, "server listening at port ", zap.Any("link", host))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("\n%s❌ Server crashed while starting: %v%s\n", colorRed, err, colorReset)
			// logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()
	fmt.Printf("%s✨ %s successfully bound to %s%s\n", colorBold+colorGreen, viper.GetString("app.name"), addr, colorReset)
	fmt.Println("----------------------------------------------------------------")

	

	<-quit
	fmt.Printf("\n%s🛑 Termination signal caught. Initiating graceful shutdown...%s\n", colorYellow, colorReset)


	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("%s❌ Graceful shutdown failed: %v%s\n", colorRed, err, colorReset)

	}

	fmt.Printf("%s👋 Shutdown operational sequences finalized clean.%s\n", colorGreen, colorReset)

}
