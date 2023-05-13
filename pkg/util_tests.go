package pkg

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

// CreateRequestTest returns a new request and a new recorder (response).And sets the necessary headers
func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}