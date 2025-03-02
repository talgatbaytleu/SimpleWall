package utils

import (
	"regexp"
	"unicode"

	"gateway/pkg/apperrors"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return apperrors.ErrInvalidPassword
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
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
