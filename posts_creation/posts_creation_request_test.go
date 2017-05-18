package posts_creation

import (
	"strings"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostsCreationRequestJsonParsing(t *testing.T) {
	var postsCreationRequest PostsCreationRequest
	data := strings.NewReader(`{"topic":"Why am I awesome?","body":"Just because Im awesome"}`)
	err := json.NewDecoder(data).Decode(&postsCreationRequest)
	assert.Nil(t, err)
}

func TestPostsCreationRequestMalformedJsonParsing(t *testing.T) {
	var postsCreationRequest PostsCreationRequest
	data := strings.NewReader(`{malformedJson":"20183292837389373"}`)
	err := json.NewDecoder(data).Decode(&postsCreationRequest)
	assert.Error(t, err, "malformed Posts Creation Request JSON ")
}
