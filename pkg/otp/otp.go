package otp

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

type Generator interface {
	RandomSecret() (string, error)
}

type GOTPGenerator struct{}

func NewGOTPGenerator() *GOTPGenerator {
	return &GOTPGenerator{}
}

func (g *GOTPGenerator) RandomSecret() (string, error) {
	key := "a@#$)(OJ)"
	secretKeyBase32 := base32.StdEncoding.EncodeToString([]byte(key))
	code, err := totp.GenerateCode(secretKeyBase32, time.Now())
	if err != nil {
		return "", fmt.Errorf("error generate code: %v", err)
	}

	return code, nil
}
