package utils

import (
	"testing"
)

func TestHashPasswordAndCheckPassword(t *testing.T) {
	password := "SecureP@ssw0rd"

	// Хешируем пароль
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword() error: %v", err)
	}

	// Проверяем, что хеш не пустой
	if hashedPassword == "" {
		t.Fatal("HashPassword() returned empty string")
	}

	// Проверяем правильный пароль
	if !CheckPassword(password, hashedPassword) {
		t.Error("CheckPassword() returned false for correct password")
	}

	// Проверяем неправильный пароль
	if CheckPassword("WrongPassword", hashedPassword) {
		t.Error("CheckPassword() returned true for incorrect password")
	}
}
