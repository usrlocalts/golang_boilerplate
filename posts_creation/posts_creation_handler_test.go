package posts_creation

import (
	"net/http/httptest"
	"github.com/stretchr/testify/require"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
	"golang_boilerplate/testutil"
	e "golang_boilerplate/errors"
	"errors"
	"github.com/newrelic/go-agent"
)

type postsServiceMock struct {
	mock.Mock
}

func (t *postsServiceMock) Create(post *Post, txn newrelic.Transaction) (string, error) {
	args := t.Called(post)
	return args.String(0), args.Error(1)
}

func TestCreatePostsHandler(t *testing.T) {
	mockPostsService := &postsServiceMock{}
	logger := testutil.TestLogger()

	post := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",

	}

	data := strings.NewReader(`{"topic":"Why am I awesome?","body":"Just because Im awesome"}`)
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/posts", data)
	require.NoError(t, err, "failed to create a request")

	mockPostsService.On("Create", post).Return("ID1", nil)

	PostsCreationHandler(mockPostsService, logger, false)(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"post":{"id":"ID1","body":"Just because Im awesome","topic":"Why am I awesome?"},"errors":null}`, w.Body.String())

}

func TestCreatePostsHandlerMalformedError(t *testing.T) {
	mockPostsService := &postsServiceMock{}
	logger := testutil.TestLogger()

	data := strings.NewReader(`{topic":"Why am I awesome?","body":"Just because Im awesome"}`)
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/posts", data)
	require.NoError(t, err, "failed to create a request")

	PostsCreationHandler(mockPostsService, logger, false)(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"post":null,"errors":[{"code":"902","entity":"Malformed JSON"}]}`, w.Body.String())

}

func TestCreatePostsHandlerError(t *testing.T) {
	mockPostsService := &postsServiceMock{}
	logger := testutil.TestLogger()

	post := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",

	}

	data := strings.NewReader(`{"topic":"Why am I awesome?","body":"Just because Im awesome"}`)
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/posts", data)
	require.NoError(t, err, "failed to create a request")

	mockPostsService.On("Create", post).Return("", e.NewError(e.GenericServiceError, errors.New("Database error")))

	PostsCreationHandler(mockPostsService, logger, false)(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"post":null,"errors":[{"code":"900","entity":"Failed to create Post"}]}`, w.Body.String())

}
