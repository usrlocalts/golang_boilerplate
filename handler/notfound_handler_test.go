package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotFoundHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/not_found", nil)
	require.NoError(t, err, "failed to create a request")

	NotFoundHandler(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
