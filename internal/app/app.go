package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/sixojke/internal/server"
	"github.com/sixojke/pkg/database"
	"github.com/sixojke/pkg/migrations"

	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/delivery"
	"github.com/sixojke/internal/repository"
	"github.com/sixojke/internal/service"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("config error: %v", err))
	}

	postgres, err := database.NewPostgresDB(cfg.Postgres)
	if err != nil {
		log.Fatal(fmt.Sprintf("postgres connection error: %v", err))
	}
	defer postgres.Close()
	log.Info("[POSTGRES] Connection successful")

	if err := migrations.MigratePostgres(cfg.Postgres); err != nil {
		log.Error(fmt.Sprintf("postgres migrate error: %v", err))
	}
	log.Info("[POSTGRES] Migrate successful")

	redis, err := database.NewRedisDB(cfg.Redis)
	if err != nil {
		log.Fatal(fmt.Sprintf("redis connection error: %v", err))
	}
	defer redis.Close()
	log.Info("[REDIS] Connection successful")

	repo := repository.NewRepository(&repository.Deps{
		Postgres: postgres,
		Redis:    redis,
	})
	services := service.NewService(&service.Deps{
		Repo: repo,
	})
	handler := delivery.NewHandler(services)

	srv := server.NewServer(cfg.HTTPServer, handler.Init())
	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %v\n", err)
		}
	}()
	log.Info(fmt.Sprintf("[SERVER] Started :%v", cfg.HTTPServer.Port))

	shutdown(srv, services, postgres, redis)
}

func shutdown(srv *server.Server, services *service.Service, postgres *sqlx.DB, redis *redis.Client) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 3 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}

	postgres.Close()
	redis.Close()

}
