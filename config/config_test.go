package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	configVars := map[string]string{
		"LOG_LEVEL": "debug",
		"APP_PORT":  "3000",
		"host":      "localhost",
		"port":      "5432",
		"database":  "database",
		"db_user":   "postgres",
		"password":  "passs",
		"pool":      "5",
		"docs_path": "./docs",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	config := Load()
	assert.Equal(t, 3000, config.Port())
	assert.Equal(t, configVars["LOG_LEVEL"], config.LogLevel())
	assert.Equal(t, configVars["host"], fatalGetString("host"))
	assert.Equal(t, configVars["database"], fatalGetString("database"))
	assert.Equal(t, configVars["db_user"], fatalGetString("db_user"))
	assert.Equal(t, configVars["password"], fatalGetString("password"))
	assert.Equal(t, configVars["pool"], fatalGetString("pool"))
	assert.Equal(t, configVars["docs_path"], config.DocsPath())
}
