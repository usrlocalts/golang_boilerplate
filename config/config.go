package config

import (
	"github.com/gojek-engineering/goconfig"
)

type Config struct {
	goconfig.BaseConfig
	sentry *sentryConfig
}

func Load() *Config {
	config := &Config{}
	config.LoadWithOptions(map[string]interface{}{"newrelic": true, "db": true})
	config.sentry = NewSentryConfig()
	return config
}

func (config *Config) Port() int {
	return config.GetIntValue("port")
}

func (config *Config) LogLevel() string {
	return config.GetValue("log_level")
}

func (config *Config) DatabaseConfig() *goconfig.DBConfig {
	return config.DBConfig()
}

func (config *Config) Sentry() *sentryConfig {
	return config.sentry
}

func (config *Config) DocsPath() string {
	return fatalGetString("docs_path")
}
