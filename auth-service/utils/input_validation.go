package utils

import (
	"regexp"

	"auth-service/internal/apperrors"
)

func ValidatePassword(password string) error {
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`
	matched, err := regexp.MatchString(passwordRegex, password)
	if err != nil {
		return err
	}

	if !matched {
		return apperrors.ErrInvalidPassword
	}
	return nil
}

func ValidateUsername(username string) error {
	usernameRegex := `^[a-zA-Z0-9_-]{3,20}$`
	matched, err := regexp.MatchString(usernameRegex, username)
	if err != nil {
		return err
	}

	if !matched {
		return apperrors.ErrInvalidUsername
	}
	return nil
}
