package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"auth-service/pkg/apperrors"
)

func ResponseErrorJson(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError // 500

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

	log.SetOutput(os.Stderr)
	log.Println(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Write(jsonData)
}
