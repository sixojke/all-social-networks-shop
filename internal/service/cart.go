package service

import (
	"fmt"

	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) GetById(userId int) (*[]domain.CartGetByIdOut, error) {
	cart, err := s.repo.GetById(userId)
	if err != nil {
		return nil, fmt.Errorf("error cart service GetById: %v", err)
	}

	return cart, nil
}

func (s *CartService) SetQuantity(inp *domain.CartSetQuantityInp) error {
	if err := s.repo.SetQuantity(inp); err != nil {
		return fmt.Errorf("error cart service SetQuantity: %v", err)
	}

	return nil
}
