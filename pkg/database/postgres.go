package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sixojke/internal/config"
)

func NewPostgresDB(cfg config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping: %v", err)
	}

	return db, nil
}
