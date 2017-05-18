package posts_creation

import (
	"github.com/stretchr/testify/mock"
	"testing"
	"golang_boilerplate/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
	"net/http/httptest"
	"golang_boilerplate/instrumentation"
	"github.com/newrelic/go-agent"
)

type postRepositoryMock struct {
	mock.Mock
}

func (t *postRepositoryMock) Create(post *Post, txn newrelic.Transaction) (string, error) {
	args := t.Called(post)
	return args.String(0), args.Error(1)
}

func TestCreatePost(t *testing.T) {
	logger := testutil.TestLogger()
	w := httptest.NewRecorder()
	txn := instrumentation.GetNewRelicTransaction(w, false)

	mockPostsRepository := &postRepositoryMock{}

	requestedPost := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	}

	mockPostsRepository.On("Create", requestedPost).Return("ID1", nil)

	postsService := NewPostsService(mockPostsRepository, logger)

	id, err := postsService.Create(requestedPost, txn)
	assert.NoError(t, err)
	assert.Equal(t, id, "ID1")
	mockPostsRepository.AssertExpectations(t)
}

func TestErrorWhileCreatingPost(t *testing.T) {
	logger := testutil.TestLogger()
	w := httptest.NewRecorder()
	txn := instrumentation.GetNewRelicTransaction(w, false)

	mockPostsRepository := &postRepositoryMock{}

	requestedPost := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	}

	mockPostsRepository.On("Create", requestedPost).Return("", errors.New("Repository Error"))

	postsService := NewPostsService(mockPostsRepository, logger)

	id, err := postsService.Create(requestedPost, txn)

	assert.Error(t, err)
	assert.Empty(t, id)
	mockPostsRepository.AssertExpectations(t)
}
