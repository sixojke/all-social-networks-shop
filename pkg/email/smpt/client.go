package email

import (
	"errors"
	"fmt"

	"github.com/sixojke/pkg/email"
	"gopkg.in/gomail.v2"
)

type SMTPSender struct {
	from string
	pass string
	host string
	port int
}

func NewSMTPSender(from, pass, host string, port int) (*SMTPSender, error) {
	if !email.IsEmailValid(from) {
		return nil, errors.New("invalid from email")
	}

	return &SMTPSender{from: from, pass: pass, host: host, port: port}, nil
}

func (s *SMTPSender) Send(input email.SendEmailInput) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetBody("text/html", input.Body)

	d := gomail.NewDialer(s.host, s.port, s.from, s.pass)

	if err := d.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email via smtp: %v", err)
	}

	return nil
}
