package logger

import (
	"context"

	"github.com/wjddn3711/structured-logger/logger/types"
)

type LoggerWrapper struct {
	logger Logger
}

func NewLoggerWrapper(loggerType types.LoggerType, ctx context.Context) Logger {
	var logger Logger

	switch loggerType {
	case types.Logrus:
		// logger = NewLogrusLogger()
	case types.ZeroLog:
		logger = NewZerologLogger()
	default:
		// logger = NewLogrusLogger(loggerType, ctx)
	}

	return logger
}
