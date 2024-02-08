package logger

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/wjddn3711/structured-logger/logger/options"
	"github.com/wjddn3711/structured-logger/logger/types"
)

type zerologLogger struct {
	logger zerolog.Logger
	ctx    context.Context
	event  zerolog.Context
	entry  map[string]interface{}
}

func newZerologLogger(settings options.LogSetting) Logger {
	logger := zerolog.New(settings.Output).With().Timestamp().Logger()

	switch settings.Level {
	case types.Debug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case types.Info:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case types.Warn:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case types.Error:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		// default level: info
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	zerolog.TimeFieldFormat = settings.TimeFormat

	return &zerologLogger{logger: logger}
}

// AddHook : 로거에 후크를 추가하는 메서드
func (l *zerologLogger) AddHook(hook interface{}) {
	// zerolog.Hook 만 지원
	if h, ok := hook.(zerolog.Hook); ok {
		l.logger = l.logger.Hook(h)
	}
}

// ApplyOption : 로그 엔트리 옵션을 적용하는 메서드
func (l *zerologLogger) ApplyOption(opts []options.EntryOption) {
	entryOtp := &options.Entry{}
	for _, opt := range opts {
		opt(entryOtp)
	}

	if entryOtp.Message != "" {
		if l.entry != nil {
			l.entry[types.MessageField] = entryOtp.Message
		} else {
			l.entry = map[string]interface{}{types.MessageField: entryOtp.Message}
		}
	}
	if entryOtp.Fields != nil {
		if l.entry != nil {
			for k, v := range entryOtp.Fields.ToFields() {
				l.entry[k] = v
			}
		} else {
			l.entry = entryOtp.Fields.ToFields()
		}
	}
}

// WithContext : 컨텍스트에 로거를 등록하는 메서드
func (l *zerologLogger) RegisterCommonField(key string, value interface{}) {
	l.logger = l.logger.With().Interface(key, value).Logger()
}

// WithContext : 컨텍스트에 로거를 등록하는 메서드
func (l *zerologLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, types.ZerologKey, l)
}

// RegisterCommonFields : 로거에 공통 필드들을 등록하는 메서드
func (l *zerologLogger) RegisterCommonFields(entry options.LogEntry) {
	l.logger = l.logger.With().Fields(entry.ToFields()).Logger()
}

// Debug : 디버그 로그를 출력하는 메서드
func (l *zerologLogger) Debug(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	temp := l.logger.With().Fields(l.entry).Logger()
	temp.Debug().Send()
}

// Info : 정보 로그를 출력하는 메서드
func (l *zerologLogger) Info(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	temp := l.logger.With().Fields(l.entry).Logger()
	temp.Info().Send()
}

// Warn : 경고 로그를 출력하는 메서드
func (l *zerologLogger) Warn(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	temp := l.logger.With().Fields(l.entry).Logger()
	temp.Warn().Send()
}

// Error : 에러 로그를 출력하는 메서드
func (l *zerologLogger) Error(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	temp := l.logger.With().Fields(l.entry).Logger()
	temp.Error().Send()
}

// Fatal : 치명적인 에러 로그를 출력하는 메서드
func (l *zerologLogger) Fatal(opts ...options.EntryOption) {
	l.ApplyOption(opts)
	temp := l.logger.With().Fields(l.entry).Logger()
	temp.Fatal().Send()
}
