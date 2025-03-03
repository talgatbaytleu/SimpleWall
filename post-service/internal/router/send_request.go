package router

import (
	"io"
	"net/http"
)

func SendRequest(method string, reqURL *string, body *io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequest(method, *reqURL, *body)
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
