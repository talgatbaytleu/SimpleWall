package service

import (
	"net/http"
	"os"
)

func AuthValidateRequest(headers http.Header) (*http.Response, error) {
	authURL := os.Getenv("AUTH_SERVICE_ADDR") + "/validate"

	req, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header = headers.Clone()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
