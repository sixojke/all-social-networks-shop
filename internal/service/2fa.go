package service

import (
	"errors"
	"fmt"

	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
	"github.com/sixojke/pkg/2fa/authenticator"
	"github.com/sixojke/pkg/otp"
)

type TwoFaService struct {
	config config.TwoFa
	repo   repository.TwoFa
	otp    otp.Generator
	twoFa  authenticator.TwoFaManager
}

func NewTwoFaService(config config.TwoFa, repo repository.TwoFa, otp otp.Generator, twoFa authenticator.TwoFaManager) *TwoFaService {
	return &TwoFaService{
		config: config,
		repo:   repo,
		otp:    otp,
		twoFa:  twoFa,
	}
}

func (s *TwoFaService) CreatePairingLink(userId int) (string, error) {
	secretCode := s.otp.RandomSecretWithLength(s.config.SecretCodeLength)

	if err := s.repo.CreatePairingLink(userId, secretCode); err != nil {
		return "", fmt.Errorf("error 2fa service CreatePairingLink: %v", err)
	}

	link := s.twoFa.GeneratePairingLink(fmt.Sprintf("%v", userId), secretCode)

	return link, nil
}

func (s *TwoFaService) СheckTwoFactorPin(userId int, pin int) (bool, error) {
	secretCode, err := s.repo.GetSecretCode(userId)
	if err != nil {
		return false, fmt.Errorf("error 2fa service СheckTwoFactorCode: %v", err)
	}

	check, err := s.twoFa.СheckTwoFactorPin(pin, secretCode)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidPin) {
			return false, domain.ErrInvalidPin
		}

		return false, fmt.Errorf("error 2fa service СheckTwoFactorCode: %v", err)
	}

	return check, nil
}
