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

// func (s *ProductsService) Create(product *domain.Product) (int, error) {
// 	id, err := s.repo.Create(product)
// 	if err != nil {
// 		return 0, fmt.Errorf("error product service Create: %v", err)
// 	}

// 	return id, nil
// }

func (s *ProductsService) GetAll(limit, offset int) (*[]domain.Product, error) {
	products, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error product service GetAll: %v", err)
	}

	return products, nil
}

// func (s *ProductsService) GetById(id int) (*domain.Product, error) {
// 	product, err := s.repo.GetById(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("error product service GetById: %v", err)
// 	}

// 	return product, nil
// }

// func (s *ProductsService) GetBySubcategoryId(id int) (*[]domain.Product, error) {
// 	products, err := s.repo.GetBySubcategory(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("error product service GetBySubcategoryId")
// 	}

// 	return products, nil
// }

// func (s *ProductsService) Update(product *domain.Product) (*domain.Product, error) {
// 	product, err := s.repo.Update(product)
// 	if err != nil {
// 		return nil, fmt.Errorf("error product service Update: %v", err)
// 	}

// 	return product, nil
// }
