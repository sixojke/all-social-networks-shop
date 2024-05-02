package service

import (
	"fmt"

	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/repository"
	"github.com/sixojke/pkg/otp"
)

type BindService struct {
	config config.Telegram
	repo   repository.Telegram
	otp    otp.Generator
}

func NewBindSerivce(repo repository.Telegram, config config.Telegram, otp otp.Generator) *BindService {
	return &BindService{
		config: config,
		repo:   repo,
		otp:    otp,
	}
}

func (s *BindService) CreateAuthLink(userId int) (string, error) {
	code := s.otp.RandomSecretWithLength(s.config.CodeLength)

	code, err := s.repo.CreateAuthLink(code, userId)
	if err != nil {
		return "", fmt.Errorf("error bind service CreateLinkTelegram: %v", err)
	}

	link := s.config.BaseLinkBot + code

	return link, nil
}

func (s *BindService) Bind(telegramId int, code string) (userId int, err error) {
	userId, err = s.repo.Bind(telegramId, code)
	if err != nil {
		return 0, fmt.Errorf("error bind service BindTelegram: %v", err)
	}

	return userId, nil
}

func (s *BindService) Unbind(userId int) error {
	if err := s.repo.Unbind(userId); err != nil {
		return fmt.Errorf("error bind service UnbindTelegram: %v", err)
	}

	return nil
}
