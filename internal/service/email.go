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
