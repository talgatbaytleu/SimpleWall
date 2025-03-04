package service

import (
	"io"
	"net/http"
)

var s3url = "http://localhost:8086/" // SHOULD BE IN .env!!!

func SendRequest(method string, reqURL *string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, *reqURL, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
