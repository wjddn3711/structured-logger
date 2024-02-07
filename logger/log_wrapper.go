package logger

import (
	"context"

	"github.com/wjddn3711/structured-logger/logger/types"
)

type LoggerWrapper struct {
	logger Logger
}

func NewLoggerWrapper(loggerType types.LoggerType) (logger Logger) {
	switch loggerType {
	case types.Logrus:
		logger = NewLogrusLogger()
	case types.ZeroLog:
		logger = NewZerologLogger()
	default:
		// no op
	}
	return
}

func FromContext(loggerType types.LoggerType, ctx context.Context) (logger Logger) {
	var logKey types.LogContextKey
	switch loggerType {
	case types.Logrus:
		logKey = types.LogrusKey
	case types.ZeroLog:
		logKey = types.ZerologKey
	default:
		// no op
		return
	}

	logger, ok := ctx.Value(logKey).(Logger)
	if !ok {
		return NewLoggerWrapper(loggerType)
	}

	return logger
}
