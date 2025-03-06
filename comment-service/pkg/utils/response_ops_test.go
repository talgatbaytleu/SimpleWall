package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"commenter/pkg/apperrors"
)

var errTest = errors.New("test error")

func TestResponseErrorJson(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode int
		expectedMsg  string
	}{
		{"Not Found", apperrors.ErrNotFound, http.StatusNotFound, "NOT found"},
		{
			"No JWT Secret",
			apperrors.ErrNoJwtSecret,
			http.StatusInternalServerError,
			"JWT_SECRET not set",
		},
		{"Invalid Token", apperrors.ErrInvalidToken, http.StatusUnauthorized, "invalid token"},
		{"Expired Token", apperrors.ErrExpiredToken, http.StatusUnauthorized, "expired token"},
		{
			"Incorrect Password",
			apperrors.ErrIncorrectPswd,
			http.StatusBadRequest,
			"incorrect user or password",
		},
		{
			"Invalid Password",
			apperrors.ErrInvalidPassword,
			http.StatusBadRequest,
			"password must be at least 8 characters long, include one uppercase letter, one lowercase letter, one digit, and one special character",
		},
		{
			"Invalid Username",
			apperrors.ErrInvalidUsername,
			http.StatusBadRequest,
			"username must be 3-20 characters long and contain only letters, numbers, underscores, or dashes",
		},
		{"Unknown Error", errTest, http.StatusInternalServerError, "test error"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Создаем фейковый HTTP-ответ
			recorder := httptest.NewRecorder()
			ResponseErrorJson(tc.err, recorder)

			// Проверяем код ответа
			if recorder.Code != tc.expectedCode {
				t.Errorf("expected status %d, got %d", tc.expectedCode, recorder.Code)
			}

			// Проверяем JSON-ответ
			var response map[string]string
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("error parsing JSON response: %v", err)
			}

			if response["error"] != tc.expectedMsg {
				t.Errorf("expected error message %q, got %q", tc.expectedMsg, response["error"])
			}
		})
	}
}
