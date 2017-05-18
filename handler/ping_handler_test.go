package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPingHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/ping", nil)
	require.NoError(t, err, "failed to create a request")

	PingHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"success\": \"pong\"}", w.Body.String())
}
