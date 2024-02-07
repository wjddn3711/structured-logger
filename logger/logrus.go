package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wjddn3711/structured-logger/logger/types"
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

func (l *LogrusLogger) beforeSend(entry *logrus.Entry) {
	if l.ctx != nil {
		l.ctx = context.WithValue(l.ctx, types.LogrusKey, entry.Logger)
	}
	return
}

func (l *LogrusLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, types.LogrusKey, l)
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
	l.beforeSend(logEntry)
	logEntry.Debug()
}

func (l *LogrusLogger) Info(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	l.beforeSend(logEntry)
	logEntry.Info()
}

func (l *LogrusLogger) Warn(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	l.beforeSend(logEntry)
	logEntry.Warn()
}

func (l *LogrusLogger) Error(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	l.beforeSend(logEntry)
	logEntry.Error()
}

func (l *LogrusLogger) Fatal(entry LogEntry, opts ...LogOptions) {
	logEntry := l.logger.WithFields(logrus.Fields(entry.ToFields()))
	for _, opt := range opts {
		if opt.message != "" {
			logEntry = logEntry.WithField("message", opt.message)
		}
	}
	l.beforeSend(logEntry)
	logEntry.Fatal()
}
