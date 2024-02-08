package logger

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wjddn3711/structured-logger/logger/options"
	"github.com/wjddn3711/structured-logger/logger/types"
)

type logrusLogger struct {
	logger *logrus.Logger
	entry  *logrus.Entry
	ctx    context.Context
}

func newLogrusLogger(settings options.LogSetting) Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: settings.TimeFormat,
	})

	// 로그 레벨 설정
	switch settings.Level {
	case types.Debug:
		logger.SetLevel(logrus.DebugLevel)
	case types.Info:
		logger.SetLevel(logrus.InfoLevel)
	case types.Warn:
		logger.SetLevel(logrus.WarnLevel)
	case types.Error:
		logger.SetLevel(logrus.ErrorLevel)
	default:
		// default level: info
		logger.SetLevel(logrus.InfoLevel)
	}

	// 로그 출력 설정
	logger.SetOutput(settings.Output)

	return &logrusLogger{logger: logger, entry: logger.WithFields(logrus.Fields{})}
}

// AddHook : 로거에 후크를 추가하는 메서드
func (l *logrusLogger) AddHook(hook interface{}) {
	// logrus.Hook 만 지원
	if h, ok := hook.(logrus.Hook); ok {
		l.logger.AddHook(h)
	}
}

// ApplyOption : 로그 엔트리 옵션을 적용하는 메서드
func (l *logrusLogger) ApplyOption(opts []options.EntryOption) {
	entryOpt := &options.Entry{}

	for _, opt := range opts {
		opt(entryOpt)
	}

	if entryOpt.Message != "" {
		l.entry = l.entry.WithField("message", entryOpt.Message)
	}
	if entryOpt.Fields != nil {
		l.entry = l.entry.WithFields(logrus.Fields(entryOpt.Fields.ToFields()))
	}
}

// WithContext : 컨텍스트에 로거를 등록하는 메서드
func (l *logrusLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, types.LogrusKey, l)
}

// RegisterCommonField : 로거에 공통 필드를 등록하는 메서드
func (l *logrusLogger) RegisterCommonField(key string, value interface{}) {
	l.entry = l.entry.WithField(key, value)
}

// RegisterCommonFields : 로거에 공통 필드들을 등록하는 메서드
func (l *logrusLogger) RegisterCommonFields(entry options.LogEntry) {
	l.entry = l.entry.WithFields(logrus.Fields(entry.ToFields()))
}

// Debug : 디버그 로그를 출력하는 메서드
func (l *logrusLogger) Debug(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	l.entry.Debug()
}

// Info : 정보 로그를 출력하는 메서드
func (l *logrusLogger) Info(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	l.entry.Info()
}

// Warn : 경고 로그를 출력하는 메서드
func (l *logrusLogger) Warn(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	l.entry.Warn()
}

// Error : 에러 로그를 출력하는 메서드
func (l *logrusLogger) Error(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	l.entry.Error()
}

// Fatal : 치명적인 에러 로그를 출력하는 메서드
func (l *logrusLogger) Fatal(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	l.entry.Fatal()
}
