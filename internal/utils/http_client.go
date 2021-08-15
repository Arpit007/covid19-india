package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

func GetClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout}
}

func DecodeResponseBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(out)
}
