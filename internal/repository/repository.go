package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type Products interface {
	Create(product *domain.Product) (int, error)
	GetById(id int) (*domain.Product, error)
	GetBySubcategory(id int) (*[]domain.Product, error)
	// Update(product *domain.Product) (*domain.Product, error)
}

type Deps struct {
	Postgres *sqlx.DB
	Redis    *redis.Client
}

type Repository struct {
	Products Products
}

func NewRepository(deps *Deps) *Repository {
	return &Repository{
		Products: NewProductsPostgres(deps.Postgres),
	}
}
