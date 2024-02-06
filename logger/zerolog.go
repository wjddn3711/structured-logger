package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	logger zerolog.Logger
	ctx    context.Context
}

func NewZerologLogger() Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &ZerologLogger{logger: logger}
}

func (l *ZerologLogger) RegisterCommonField(key string, value interface{}) {
	l.logger = l.logger.With().Interface(key, value).Logger()
}

func (l *ZerologLogger) WithContext(ctx context.Context) Logger {
	return &ZerologLogger{logger: l.logger, ctx: ctx}
}

func (l *ZerologLogger) RegisterCommonFields(entry LogEntry) {
	// 중복된 키가 있는 경우 덮어 씌워짐
	for k, v := range entry.ToFields() {
		l.logger = l.logger.With().Interface(k, v).Logger()
	}
}

func (l *ZerologLogger) Debug(entry LogEntry, opts ...LogOptions) {
	log := l.logger.Debug().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	log.Send()
}

func (l *ZerologLogger) Info(entry LogEntry, opts ...LogOptions) {
	log := l.logger.Info().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	log.Send()
}

func (l *ZerologLogger) Warn(entry LogEntry, opts ...LogOptions) {
	log := l.logger.Warn().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	log.Send()
}

func (l *ZerologLogger) Error(entry LogEntry, opts ...LogOptions) {
	log := l.logger.Error().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	log.Send()
}

func (l *ZerologLogger) Fatal(entry LogEntry, opts ...LogOptions) {
	log := l.logger.Fatal().Fields(entry.ToFields())
	for _, opt := range opts {
		if opt.message != "" {
			log = log.Str("message", opt.message)
		}
	}
	log.Send()
}
