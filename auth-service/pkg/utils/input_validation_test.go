package utils

import (
	"testing"

	"auth-service/pkg/apperrors"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  error
	}{
		{"Valid password", "Str0ng@Pass!", nil},
		{"Too short", "Ab1!", apperrors.ErrInvalidPassword},
		{"No special char", "Abcdef12", apperrors.ErrInvalidPassword},
		{"No uppercase", "abcdefg1!", apperrors.ErrInvalidPassword},
		{"No lowercase", "ABCDEFG1!", apperrors.ErrInvalidPassword},
		{"No digit", "Abcdefgh!", apperrors.ErrInvalidPassword},
		{"Only special chars", "@$%^&*()_", apperrors.ErrInvalidPassword},
		{"Only numbers", "12345678", apperrors.ErrInvalidPassword},
		{"Only uppercase", "ABCDEFGH", apperrors.ErrInvalidPassword},
		{"Only lowercase", "abcdefgh", apperrors.ErrInvalidPassword},
		{
			"Mix of required chars",
			"Aa1@",
			apperrors.ErrInvalidPassword,
		}, // Минимальный валидный пароль (8 символов)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if err != tt.wantErr {
				t.Errorf("ValidatePassword(%q) = %v, want %v", tt.password, err, tt.wantErr)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  error
	}{
		{"Valid username", "valid_user123", nil},
		{"Too short", "ab", apperrors.ErrInvalidUsername},
		{"Too long", "thisisaverylongusername123", apperrors.ErrInvalidUsername},
		{"Invalid characters", "user@name!", apperrors.ErrInvalidUsername},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUsername(tt.username)
			if err != tt.wantErr {
				t.Errorf("ValidateUsername(%q) = %v, want %v", tt.username, err, tt.wantErr)
			}
		})
	}
}
