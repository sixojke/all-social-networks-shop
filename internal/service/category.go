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
		return 0, fmt.Errorf("error category service CreateCategory: %v", err)
	}

	return
}

func (s *CategoryService) UpdateCategory(category *domain.Category) error {
	if err := s.repo.CategoryEdit(category); err != nil {
		return fmt.Errorf("error category service UpdateCategory: %v", err)
	}

	return nil
}

func (s *CategoryService) GetCategories() (*[]domain.Category, error) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, fmt.Errorf("error category service GetCategories: %v", err)
	}

	return categories, nil
}

func (s *CategoryService) GetSubcategories(categoryId int) (*[]domain.Subcategory, error) {
	subcategories, err := s.repo.GetSubcategories(categoryId)
	if err != nil {
		return nil, fmt.Errorf("error category service GetCategories: %v", err)
	}

	return subcategories, nil
}
