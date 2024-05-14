package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"project-structure-template/internal/constants/dbinstance"
	"project-structure-template/internal/handler/middleware"

	"syscall"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Initiate
//
//	@title			project setup API
//	@version		0.1
//	@host			localhost:8000
//	@BasePath		/v1
func Initiator(ctx context.Context) {

	log := InitLogger()
	log.Info(ctx, "logger initialized")

	log.Info(ctx, "initializing config")
	InitConfig("config", "config", log)
	log.Info(ctx, "config initialized")

	log.Info(ctx, "initializing database")
	Conn := InitDB(viper.GetString("database.url"), log)
	log.Info(ctx, "database initialized")

	log.Info(ctx, "initializing migration")
	m := InitiateMigration(viper.GetString("migration.path"), viper.GetString("database.url"), log)
	UpMigration(m, log)
	log.Info(ctx, "migration initialized")

	log.Info(ctx, "initializing persistence layer")
	persistence := InitPersistence(dbinstance.New(Conn), log)
	log.Info(ctx, "persistence layer initialized")

	log.Info(ctx, "initializing module")
	module := InitModule(persistence, log)
	log.Info(ctx, "module initialized")

	log.Info(ctx, "initializing handler")
	handler := InitHandler(module, log, viper.GetDuration("server.timeout"))
	log.Info(ctx, "handler initialized")

	log.Info(ctx, "initializing server")
	server := gin.New()
	server.Use(ginzap.RecoveryWithZap(log.GetZapLogger().Named("gin.recovery"), true))
	server.Use(middleware.ErrorHandler())
	log.Info(ctx, "server initialized")

	log.Info(ctx, "initializing router")
	v1 := server.Group("/v1")
	InitRouter(v1, handler)
	log.Info(ctx, "router initialized")

	srv := &http.Server{
		Addr:    viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Handler: server,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	go func() {
		log.Info(ctx, "server started",
			zap.String("host", viper.GetString("server.host")),
			zap.Int("port", viper.GetInt("server.port")))
		log.Info(ctx, fmt.Sprintf("server stopped with error %v", srv.ListenAndServe()))
	}()
	sig := <-quit
	log.Info(ctx, fmt.Sprintf("server shutting down with signal %v", sig))
	ctx, cancel := context.WithTimeout(ctx, viper.GetDuration("server.timeout"))
	defer cancel()

	log.Info(ctx, "shutting down server")
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(ctx, fmt.Sprintf("error while shutting down server: %v", err))
	} else {
		log.Info(ctx, "server shutdown complete")
	}
}
