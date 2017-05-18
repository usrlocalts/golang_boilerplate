package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestNewError(t *testing.T) {
	bgErr := errors.New("golang_boilerplate error")
	err := NewError("123", bgErr)

	expectedError := &Error{Code: "123", Entity: bgErr.Error()}
	assert.Equal(t, err, expectedError)
	assert.Equal(t, err.ErrorCode(), "123")
	assert.Equal(t, err.Error(), bgErr.Error())
}

func TestNewErrorSerializesTheErrorObjectToCustomError(t *testing.T) {
	customErr := NewError("123", errors.New("golang_boilerplate error"))

	expectedError, err := json.Marshal(&Error{
		Code:   "123",
		Entity: "golang_boilerplate error",
	})

	assert.NoError(t, err)
	marshalledErrorResponse, err := json.Marshal(customErr)
	assert.NoError(t, err)
	assert.Equal(t, string(marshalledErrorResponse), string(expectedError))
}
