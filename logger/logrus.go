package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

func NewLogrusLogger() Logger {
	logger := logrus.New()
	logger.Out = os.Stdout

	return &LogrusLogger{logger: logger}
}

func (l *LogrusLogger) WithContext(ctx context.Context) Logger {
	return &LogrusLogger{logger: l.logger, ctx: ctx}
}

func (l *LogrusLogger) RegisterCommonField(key string, value interface{}) {
	l.logger = l.logger.WithField(key, value).Logger
}

func (l *LogrusLogger) RegisterCommonFields(entry LogEntry) {
	l.logger = l.logger.WithFields(logrus.Fields(entry.ToFields())).Logger
}

func (l *LogrusLogger) Debug(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	logEntry.Debug()
}

func (l *LogrusLogger) Info(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	logEntry.Info()
}

func (l *LogrusLogger) Warn(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	logEntry.Warn()
}

func (l *LogrusLogger) Error(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	logEntry.Error()
}

func (l *LogrusLogger) Fatal(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	logEntry.Fatal()
}
