package service

import (
	"fmt"

	"github.com/sixojke/pkg/email"
)

type emailSerive struct {
	sender email.Sender
}

func NewEmailService(sender email.Sender) *emailSerive {
	return &emailSerive{sender: sender}
}

type VerificationEmailInp struct {
	Email            string
	VerificationCode string
}

type PasswordRecoveryInp struct {
	Email        string
	RecoveryLink string
}

func (s *emailSerive) SendUserVerificationEmail(inp *VerificationEmailInp) error {
	if err := s.sender.Send(email.SendEmailInput{
		To:      inp.Email,
		Subject: "Confirm authorization",
		Body:    fmt.Sprintf("Verification code: %v", inp.VerificationCode),
	}); err != nil {
		return err
	}

	return nil
}

func (s *emailSerive) SendUserPasswordRecoveryEmail(inp *PasswordRecoveryInp) error {
	if err := s.sender.Send(email.SendEmailInput{
		To:      inp.Email,
		Subject: "Password recovery",
		Body:    fmt.Sprintf("Link: %s", inp.RecoveryLink),
	}); err != nil {
		return err
	}

	return nil
}
