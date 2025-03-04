package utils

import (
	"errors"
	"strings"
)

func GetURLVar(nth int, url string) (string, error) {
	parts := strings.Split(url, "/")
	if len(parts) < nth+1 {
		return "", errors.New("URL can't be parsed")
	}
	urlVar := parts[nth]
	return urlVar, nil
}
