package service

import (
	"fmt"

	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *domain.Category) (id int, err error) {
	id, err = s.repo.CreateCategory(category)
	if err != nil {
		return 0, fmt.Errorf("category service CreateCategory: %v", err)
	}

	return
}

func (s *CategoryService) GetCategories() (*[]domain.Category, error) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, fmt.Errorf("errors products service GetCategories: %v", err)
	}

	return categories, nil
}

func (s *CategoryService) GetSubcategories(categoryId int) (*[]domain.Subcategory, error) {
	subcategories, err := s.repo.GetSubcategories(categoryId)
	if err != nil {
		return nil, fmt.Errorf("errors products service GetCategories: %v", err)
	}

	return subcategories, nil
}
