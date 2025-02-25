package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrNotFound      = errors.New("NOT found")
	ErrNoJwtSecret   = errors.New("JWT_SECRET in not set")
	ErrInvalidToken  = errors.New("invalid token")
	ErrExpiredToken  = errors.New("expired token")
	ErrIncorrectPswd = errors.New("incorrect user or password")
)

type myErrorType struct {
	ErrorMessage string `json:"error"`
}

func ResponseErrorJson(err error, w http.ResponseWriter) {
	var errStruct myErrorType
	var statusCode int = 400
	errStruct.ErrorMessage = err.Error()

	// specify appropriate statusCode

	jsonData, _ := json.MarshalIndent(errStruct, "", " ")
	http.Error(w, string(jsonData), statusCode)
}
