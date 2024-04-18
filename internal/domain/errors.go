package domain

import "errors"

var (
	ErrUserNotFound            = errors.New("user doesn't exists")
	ErrVerificationCodeInvalid = errors.New("verification code is invalid")
	ErrCodeIsEmpty             = errors.New("code is empty")
	ErrDuplicateKey            = errors.New("dublicate key")
	ErrUserNotVerified         = errors.New("user is not verified")
)
