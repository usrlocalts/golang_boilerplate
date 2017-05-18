package logger

import (
	"os"
	"golang_boilerplate/config"
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
)

type Log interface {
	logrus.FieldLogger
}

type Logger struct {
	*logrus.Logger
}

type LoggerError struct {
	Error error
}

func panicIfError(err error) {
	if err != nil {
		panic(LoggerError{err})
	}
}

func SetupLogger(config *config.Config) *Logger {
	level, err := logrus.ParseLevel(config.LogLevel())
	panicIfError(err)

	logrusVar := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	Log := &Logger{logrusVar}

	if config.Sentry().Enabled() {
		sentryLevels := []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		}
		sentryHook, err := logrus_sentry.NewSentryHook(config.Sentry().Dsn(), sentryLevels)
		panicIfError(err)

		Log.Hooks.Add(sentryHook)
	}
	return Log
}

func BuildContext(context string) logrus.Fields {
	return logrus.Fields{
		"context": context,
	}
}
