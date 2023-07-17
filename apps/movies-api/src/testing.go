package motesting

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func ParseJSONReponse[T interface{}](t testing.TB, body io.Reader, got T) T {
	t.Helper()

	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("Unable to parse response body: %q into HealthReponse, %v", body, err)
	}

	return got
}

func NewRequest(method, url string, body interface{}) *http.Request {
	jsonBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, url, bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	return req
}
