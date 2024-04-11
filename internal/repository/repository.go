package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/config"
)

type Deps struct {
	Postgres *sqlx.DB
	Redis    *redis.Client
	Config   *config.Config
}

type Repository struct {
}

func NewRepository(deps *Deps) *Repository {
	return &Repository{}
}
