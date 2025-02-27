package apperrors

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("NOT found")
	ErrNoJwtSecret     = errors.New("JWT_SECRET not set")
	ErrInvalidToken    = errors.New("invalid token")
	ErrExpiredToken    = errors.New("expired token")
	ErrIncorrectPswd   = errors.New("incorrect user or password")
	ErrInvalidPassword = errors.New(
		"password must be at least 8 characters long, include one uppercase letter, one lowercase letter, one digit, and one special character",
	)
	ErrInvalidUsername = errors.New(
		"username must be 3-20 characters long and contain only letters, numbers, underscores, or dashes",
	)
)
