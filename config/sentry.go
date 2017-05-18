package config

type sentryConfig struct {
	enabled bool
	dsn     string
}

func (self sentryConfig) Enabled() bool {
	return self.enabled
}

func (self sentryConfig) Dsn() string {
	return self.dsn
}

func NewSentryConfig() *sentryConfig {
	return &sentryConfig{
		enabled: getFeature("sentry_enabled"),
		dsn:     getString("sentry_dsn"),
	}
}
