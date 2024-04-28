package otp

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/xlzd/gotp"
)

type Generator interface {
	RandomSecret() (string, error)
	RandomSecretWithLength(length int) string
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

func (g *GOTPGenerator) RandomSecretWithLength(length int) string {
	return gotp.RandomSecret(length)
}
