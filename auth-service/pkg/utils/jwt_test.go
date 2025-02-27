package utils

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestValidateJWT_ExpiredToken(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "supersecret")

	claims := jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(-time.Minute).Unix(), // Истёкший токен
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("supersecret"))

	_, err := ValidateJWT(tokenString)
	if !errors.Is(err, jwt.ErrTokenExpired) {
		t.Fatalf("Expected ErrTokenExpired, got %v", err)
	}
}

func TestValidateJWT_ValidToken(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "supersecret")

	claims := jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(time.Hour).Unix(), // Валидный токен
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("supersecret"))

	userID, err := ValidateJWT(tokenString)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if userID != "123" {
		t.Fatalf("Expected user_id 123, got %v", userID)
	}
}

func TestValidateJWT_InvalidSignature(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "supersecret")

	claims := jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("wrongsecret")) // Подписали другим ключом

	_, err := ValidateJWT(tokenString)
	if !errors.Is(err, jwt.ErrSignatureInvalid) {
		t.Fatalf("Expected ErrSignatureInvalid, got %v", err)
	}
}

func TestValidateJWT_MalformedToken(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "supersecret")

	_, err := ValidateJWT("invalid.token.string")
	if !errors.Is(err, jwt.ErrTokenMalformed) {
		t.Fatalf("Expected ErrTokenMalformed, got %v", err)
	}
}

func TestValidateJWT_MissingSecret(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Unsetenv("JWT_SECRET") // Убираем JWT_SECRET

	claims := jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("supersecret"))

	_, err := ValidateJWT(tokenString)
	if err == nil {
		t.Fatal("Expected error due to missing JWT_SECRET, got nil")
	}
}
