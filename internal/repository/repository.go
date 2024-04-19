package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type Users interface {
	Create(user *domain.User, code string) (int, error)
	GetByCredentials(username, password string) (*domain.User, error)
	GetByRefreshToken(refreshToken string) (*domain.Session, error)
	Verify(userId int, code string) error
	SetSession(session *domain.Session) error
}

type Products interface {
	// Create(product *domain.Product) (int, error)
	GetAll(filters *domain.ProductFilters) (*domain.Pagination, error)
	// GetById(id int) (*domain.Product, error)
	// GetBySubcategory(id int) (*[]domain.Product, error)
	// Update(product *domain.Product) (*domain.Product, error)
	GetCategories() (*[]domain.Category, error)
	GetSubcategories(categoryId int) (*[]domain.Subcategory, error)
}

type Deps struct {
	Postgres *sqlx.DB
	Redis    *redis.Client
}

type Repository struct {
	Users    Users
	Products Products
}

func NewRepository(deps *Deps) *Repository {
	return &Repository{
		Users:    NewUsersPostgres(deps.Postgres),
		Products: NewProductsPostgres(deps.Postgres),
	}
}
