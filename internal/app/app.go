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

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/sixojke/internal/server"
	"github.com/sixojke/pkg/2fa/authenticator"
	"github.com/sixojke/pkg/auth"
	"github.com/sixojke/pkg/database"
	email "github.com/sixojke/pkg/email/smpt"
	"github.com/sixojke/pkg/hash"
	"github.com/sixojke/pkg/migrations"
	"github.com/sixojke/pkg/otp"
	"github.com/sixojke/pkg/payments/payok"

	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/delivery"
	"github.com/sixojke/internal/repository"
	"github.com/sixojke/internal/service"
)

// @title All social networks shop API
// @version 1.0
// @description REST API for shop

// @host localhost:8009
// @BasePath /api/v1

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey TelegramAuth
// @in header
// @name Authorization
func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("config error: %v", err))
	}
	fmt.Println(cfg.Authenticator.AppName)

	payokClient := payok.NewClient(cfg.Payok.ShopId, cfg.Payok.SuccessUrl, cfg.Payok.SecretKey)

	hasher := hash.NewSHA1Hasher("my-salt")

	otpGenerator := otp.NewGOTPGenerator()

	authenticatorManager := authenticator.NewManager(cfg.Authenticator.AppName)

	emaildSender, err := email.NewSMTPSender(cfg.EmailSender.From, cfg.EmailSender.Password, cfg.EmailSender.Host, cfg.EmailSender.Port)
	if err != nil {
		log.Fatal(fmt.Errorf("email sender: %v", err))
	}

	tokenManager, err := auth.NewManager(cfg.Service.Users.Auth.SigningKey)
	if err != nil {
		log.Fatal(fmt.Errorf("token manager: %v", err))
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

	repo := repository.NewRepository(&repository.Deps{
		Postgres: postgres,
	})
	services := service.NewService(&service.Deps{
		Repo:         repo,
		Config:       &cfg.Service,
		Hasher:       hasher,
		OtpGenerator: otpGenerator,
		TwoFaManager: authenticatorManager,
		EmailSender:  emaildSender,
		TokenManager: tokenManager,
		PayokClient:  payokClient,
	})
	handler := delivery.NewHandler(cfg.Handler, services, tokenManager)

	srv := server.NewServer(cfg.HTTPServer, handler.Init())
	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %v\n", err)
		}
	}()
	log.Info(fmt.Sprintf("[SERVER] Started :%v", cfg.HTTPServer.Port))

	shutdown(srv, postgres)
}

func shutdown(srv *server.Server, postgres *sqlx.DB) {
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
}
