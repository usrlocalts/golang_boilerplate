package posts_creation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	e "golang_boilerplate/errors"
	"errors"
)

func TestPostsCreationResponse(t *testing.T) {

	postsDescription := &PostsDescription{Id: "ID1", Topic: "Why am I awesome?", Body: "Just because Im awesome"}

	postCreationResponse := PostsCreationResponse{Post: postsDescription}
	jsonResponse, err := json.Marshal(postCreationResponse)
	assert.Equal(t, string(jsonResponse), `{"post":{"id":"ID1","body":"Just because Im awesome","topic":"Why am I awesome?"},"errors":null}`)
	assert.Nil(t, err)
}

func TestNewPostsCreationResponse(t *testing.T) {
	post := &Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",

	}

	postCreationResponse := NewPostsCreationResponse(post, "ID1")
	jsonResponse, err := json.Marshal(postCreationResponse)
	assert.Equal(t, string(jsonResponse), `{"post":{"id":"ID1","body":"Just because Im awesome","topic":"Why am I awesome?"},"errors":null}`)
	assert.Nil(t, err)
}

func TestPostsCreationResponseError(t *testing.T) {

	postCreationResponse := NewErrorPostCreationResponse([]*e.Error{e.NewError("900", errors.New("database error"))})
	jsonResponse, err := json.Marshal(postCreationResponse)
	assert.Equal(t, string(jsonResponse), `{"post":null,"errors":[{"code":"900","entity":"database error"}]}`)
	assert.Nil(t, err)
}
