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

type ProductFilters struct {
	Limit         int
	Offset        int
	CategoryId    int
	SubcategoryId int
	IsAvailable   int
	SortPrice     string
	SortDefect    string
}

func (s *ProductsService) GetAll(filters *domain.ProductFilters) (*domain.Pagination, error) {
	products, err := s.repo.GetAll(filters)
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

func (s *ProductsService) GetCategories() (*[]domain.Category, error) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, fmt.Errorf("errors products service GetCategories: %v", err)
	}

	return categories, nil
}

func (s *ProductsService) GetSubcategories(categoryId int) (*[]domain.Subcategory, error) {
	subcategories, err := s.repo.GetSubcategories(categoryId)
	if err != nil {
		return nil, fmt.Errorf("errors products service GetCategories: %v", err)
	}

	return subcategories, nil
}
