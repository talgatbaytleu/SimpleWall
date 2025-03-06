package service

import (
	"net/http"
)

func AuthValidateRequest(headers http.Header) (*http.Response, error) {
	authURL := "http://localhost:8081/validate"

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
