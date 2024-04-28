package service

import (
	"fmt"

	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
	"github.com/sixojke/pkg/otp"
)

type ReferralSystemService struct {
	repo         repository.ReferralSystem
	config       config.ReferralSystemService
	otpGenerator otp.Generator
}

func NewReferralSystemService(repo repository.ReferralSystem, config config.ReferralSystemService, otpGenerator otp.Generator) *ReferralSystemService {
	return &ReferralSystemService{
		repo:         repo,
		config:       config,
		otpGenerator: otpGenerator,
	}
}

func (s *ReferralSystemService) CreateCode(description string) (link string, err error) {
	referralCode := s.otpGenerator.RandomSecretWithLength(s.config.CodeLength)

	if err := s.repo.CreateCode(domain.ReferralSystem{
		ReferralCode: referralCode,
		Description:  description,
	}); err != nil {
		return "", fmt.Errorf("error referral system service Create: %v", err)
	}

	referralLink := referralCode

	return referralLink, nil
}

func (s *ReferralSystemService) AddVisitor(referralCode string) error {
	if err := s.repo.AddVisitor(referralCode); err != nil {
		return fmt.Errorf("error referral system service AddVisitor: %v", err)
	}

	return nil
}

func (s *ReferralSystemService) GetStats(limit, offset int) (*domain.Pagination, error) {
	stats, err := s.repo.GetStats(limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error referral system service GetStats: %v", err)
	}

	return stats, nil
}

func (s ReferralSystemService) DeleteCode(referralCode string) error {
	if err := s.repo.DeleteCode(referralCode); err != nil {
		return fmt.Errorf("error referral system service DeleteCode: %v", err)
	}

	return nil
}
