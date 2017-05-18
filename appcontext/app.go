package appcontext

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent"
	"golang_boilerplate/config"
	"golang_boilerplate/logger"
)

type AppContext struct {
	newrelicApp newrelic.Application
	logger      logger.Log
	config      *config.Config
}

type appContextError struct {
	Error error
}

func panicIfError(err error, werr error) {
	if err != nil {
		panic(appContextError{werr})
	}
}

func NewAppContext(logger logger.Log, config *config.Config) *AppContext {
	newrelicApp, err := newrelic.NewApplication(config.Newrelic())
	panicIfError(err, fmt.Errorf("Unable initiate NewRelic: %v", err))
	return &AppContext{
		newrelicApp: newrelicApp,
		config:      config,
		logger:      logger,
	}
}

func (s *AppContext) NewrelicApp() newrelic.Application {
	return s.newrelicApp
}

func (s *AppContext) GetLogger() logger.Log {
	return s.logger
}

func (s *AppContext) GetConfig() *config.Config {
	return s.config
}
