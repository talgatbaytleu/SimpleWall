package utils

import (
	"io"
	"net/http"
)

func ParseFormData(r *http.Request) (string, io.ReadCloser, error) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		return "", nil, err
	}

	description := r.FormValue("description")

	// Получаем файл image
	file, _, err := r.FormFile("image")
	if err != nil {
		return "", nil, err
	}

	return description, file, nil
}
