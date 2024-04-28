package service

import (
	"fmt"

	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
)

type LogService struct {
	repo repository.Log
}

func NewLogService(repo repository.Log) *LogService {
	return &LogService{
		repo: repo,
	}
}

func (s *LogService) WriteAdminLog(log *domain.Log) error {
	if err := s.repo.WriteAdminLog(log); err != nil {
		return fmt.Errorf("error log service WriteAdminLog: %v", err)
	}

	return nil
}

func (s *LogService) GetAdminLogs(limit, offset int) (*domain.Pagination, error) {
	logs, err := s.repo.GetAdminLogs(limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error log service GetAdminLogs: %v", err)
	}

	return logs, nil
}
