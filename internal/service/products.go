package service

import (
	"fmt"

	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
)

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{repo: repo}
}

func (r *ProductsService) Create(product *domain.Product) (int, error) {
	id, err := r.repo.Create(product)
	if err != nil {
		return 0, fmt.Errorf("error product service: %v", err)
	}

	return id, nil
}
