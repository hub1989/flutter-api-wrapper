package utils

import (
	"encoding/json"
	"net/http"
)

func MapToJson(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func EnrichRequestWithAuth(req *http.Request, token string, headers map[string]string) *http.Request {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	return req
}

func EnrichRequest(req *http.Request, headers map[string]string) *http.Request {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	return req
}
