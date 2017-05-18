package testutil

import (
	"net/http/httptest"
	"net/http"
)

func NewTestServer(router http.Handler) *httptest.Server {
	return httptest.NewServer(router)
}
