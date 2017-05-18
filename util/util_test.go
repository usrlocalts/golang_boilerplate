package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathForNewRelic(t *testing.T) {
	path := "/posts"
	expectedPath := "/posts"
	newRelicCustomPath := GetPathForNewRelic(path)
	assert.Equal(t, expectedPath, newRelicCustomPath)
}

func TestGetPathForNewRelicIfPathIsEmpty(t *testing.T) {
	path := ""
	expectedPath := ""
	newRelicCustomPath := GetPathForNewRelic(path)
	assert.Equal(t, expectedPath, newRelicCustomPath)
}

func TestGetPathForNewRelicIfPathDoesNotContainMultipleDigit(t *testing.T) {
	path := "/posts"
	expectedPath := "/posts"
	newRelicCustomPath := GetPathForNewRelic(path)
	assert.Equal(t, expectedPath, newRelicCustomPath)
}
