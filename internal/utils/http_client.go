package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

// GetClient Get http.Client with Timeout
func GetClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout}
}

// DecodeResponseBody Decode Response.Body to specified type
func DecodeResponseBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(out)
}
