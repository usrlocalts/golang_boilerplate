package posts_creation

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPostCreation(t *testing.T) {
	actualPost := Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",

	}

	expectedPost := Post{
		Topic: "Why am I awesome?",
		Body:  "Just because Im awesome",
	}

	assert.Equal(t, actualPost, expectedPost)

}

func TestShouldCreatePost(t *testing.T) {
	topic := "TID1"
	body := "MID1"

	till := NewPost(topic, body)
	assert.Equal(t, till.Topic, topic, "expected topic to be equal")
	assert.Equal(t, till.Body, body, "expected bopdy to be equal")
}
