package service

import (
	"net/http"
)

func AuthValidateRequest(headers http.Header) (*http.Response, error) {
	authURL := "http://localhost:8081/validate"

	// Создаем новый HTTP-запрос, передавая заголовки авторизации от клиента
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

	// fmt.Println(resp)
	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("unauthorized")
	// }
	//
	// if resp.Header.Get("X-User-ID") == "" {
	// 	return nil, fmt.Errorf("missing X-User-ID")
	// }

	return resp, nil
}
