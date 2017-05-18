package config

import (
	"testing"
	"reflect"
	"github.com/stretchr/testify/assert"
)

func TestNewSentryConfig(t *testing.T) {
	Load()
	sentryConfig := NewSentryConfig()
	assert.Equal(t, reflect.TypeOf(sentryConfig).String(), "*config.sentryConfig")
	assert.Equal(t, sentryConfig.Enabled(), false)
	assert.NotNil(t, sentryConfig.Dsn())
}
