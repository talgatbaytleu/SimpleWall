package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"auth-service/internal/apperrors"
)

func ResponseErrorJson(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError // 500 по умолчанию

	// Определяем статус-код в зависимости от типа ошибки
	switch err {
	case apperrors.ErrNotFound:
		statusCode = http.StatusNotFound // 404
	case apperrors.ErrNoJwtSecret:
		statusCode = http.StatusInternalServerError // 500
	case apperrors.ErrInvalidToken, apperrors.ErrExpiredToken:
		statusCode = http.StatusUnauthorized // 401
	case apperrors.ErrIncorrectPswd, apperrors.ErrInvalidPassword, apperrors.ErrInvalidUsername:
		statusCode = http.StatusBadRequest // 400
	}

	// Логируем ошибку в stderr
	log.SetOutput(os.Stderr)
	log.Println(err)

	// Отправляем JSON-ответ
	jsonData, _ := json.MarshalIndent(map[string]string{"error": err.Error()}, "", " ")
	http.Error(w, string(jsonData), statusCode)
}
