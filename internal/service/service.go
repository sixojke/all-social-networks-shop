package service

import (
	"github.com/sixojke/internal/repository"
)

type Deps struct {
	Repo *repository.Repository
}

type Service struct{}

func NewService(deps *Deps) *Service {
	return &Service{}
}
