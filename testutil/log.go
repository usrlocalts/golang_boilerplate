package testutil

import (
	"golang_boilerplate/logger"
	"github.com/Sirupsen/logrus"
	"bytes"
)

func TestLogger() logger.Log {
	logBuffer := []byte{}
	logWriter := bytes.NewBuffer(logBuffer)
	log := &logrus.Logger{
		Out:       logWriter,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.DebugLevel,
	}
	return log
}
