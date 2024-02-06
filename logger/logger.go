package logger

import "context"

type LogOptions struct {
	message string
}

type LogEntry interface {
	ToFields() map[string]interface{}
}

type Logger interface {
	WithContext(ctx context.Context) Logger
	RegisterCommonField(key string, value interface{})
	RegisterCommonFields(fields LogEntry)
	Debug(fields LogEntry, opts ...LogOptions)
	Info(fields LogEntry, opts ...LogOptions)
	Warn(fields LogEntry, opts ...LogOptions)
	Error(fields LogEntry, opts ...LogOptions)
	Fatal(fields LogEntry, opts ...LogOptions)
}

func WithMessage(message string) LogOptions {
	return LogOptions{message: message}
}
