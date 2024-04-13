package service

import (
	"fmt"

	"github.com/sixojke/pkg/email"
)

type EmailSerive struct {
	sender email.Sender
}

func NewEmailService(sender email.Sender) *EmailSerive {
	return &EmailSerive{sender: sender}
}

type VerificationEmailInp struct {
	Email string
	Code  string
}

func (s *EmailSerive) SendUserVerificationEmail(inp *VerificationEmailInp) error {
	if err := s.sender.Send(email.SendEmailInput{
		To:      inp.Email,
		Subject: "Confirm authorization",
		Body:    fmt.Sprintf("Verification code: %v", inp.Code),
	}); err != nil {
		return err
	}

	return nil
}
