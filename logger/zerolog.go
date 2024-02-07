package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/wjddn3711/structured-logger/logger/types"
)

type ZerologLogger struct {
	logger zerolog.Logger
	ctx    context.Context
}

func NewZerologLogger() Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &ZerologLogger{logger: logger}
}

func (l *ZerologLogger) beforeSend(entry zerolog.Context) {
	// if l.ctx != nil {
	// 	l.ctx = context.WithValue(l.ctx, types.ZerologKey, entry.Logger())
	// }
	return
}

func (l *ZerologLogger) RegisterCommonField(key string, value interface{}) {
	l.logger = l.logger.With().Interface(key, value).Logger()
}

func (l *ZerologLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, types.ZerologKey, l)
}

func (l *ZerologLogger) RegisterCommonFields(entry LogEntry) {
	// 중복된 키가 있는 경우 덮어 씌워짐
	l.logger = l.logger.With().Fields(entry.ToFields()).Logger()
}

func (l *ZerologLogger) Debug(entry LogEntry, opts ...LogOptions) {
	log := l.logger.With().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	l.beforeSend(log)
	logger := log.Logger()
	logger.Debug().Send()
}

func (l *ZerologLogger) Info(entry LogEntry, opts ...LogOptions) {
	log := l.logger.With().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	l.beforeSend(log)
	logger := log.Logger()
	logger.Info().Send()
}

func (l *ZerologLogger) Warn(entry LogEntry, opts ...LogOptions) {
	log := l.logger.With().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	l.beforeSend(log)
	logger := log.Logger()
	logger.Warn().Send()
}

func (l *ZerologLogger) Error(entry LogEntry, opts ...LogOptions) {
	log := l.logger.With().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	l.beforeSend(log)
	logger := log.Logger()
	logger.Error().Send()
}

func (l *ZerologLogger) Fatal(entry LogEntry, opts ...LogOptions) {
	log := l.logger.With().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	l.beforeSend(log)
	logger := log.Logger()
	logger.Fatal().Send()
}
