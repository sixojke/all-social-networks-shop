package service

import (
	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
)

type Products interface {
	Create(product *domain.Product) (int, error)
	GetById(id int) (*domain.Product, error)
	GetBySubcategoryId(id int) (*[]domain.Product, error)
	// Update(product *domain.Product) (*domain.Product, error)
}

type Deps struct {
	Repo *repository.Repository
}

type Service struct {
	Products Products
}

func NewService(deps *Deps) *Service {
	return &Service{
		Products: NewProductsService(deps.Repo.Products),
	}
}
