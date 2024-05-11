package domain

import "errors"

var (
	ErrInvalidPassword         = errors.New("invalid password")
	ErrInvalidPin              = errors.New("invalid pin")
	ErrUserNotFound            = errors.New("user doesn't exists")
	ErrVerificationCodeInvalid = errors.New("verification code is invalid")
	ErrCodeIsEmpty             = errors.New("code is empty")
	ErrDuplicateKey            = errors.New("dublicate key")
	ErrUserNotVerified         = errors.New("user is not verified")
	ErrProductsNotFound        = errors.New("products not found")
	ErrDataNotFound            = errors.New("data not found")
	ErrCategoriesNotFound      = errors.New("categories not found")
	ErrSubcategoriesNotFound   = errors.New("subcategories not found")
)
