package logger

import (
	"context"
	"os"

	"github.com/wjddn3711/structured-logger/logger/options"
	"github.com/wjddn3711/structured-logger/logger/types"
)

// Wrapper : 로거 래퍼
//   - logger(Logger): 로거 인터페이스
//
// Example:
//
//	// 로거 래퍼 생성
//	log := logger.NewLoggerWrapper(types.ZeroLog)
type Wrapper struct {
	logger Logger
}

// NewWrapper : 로거 래퍼 생성자
//   - loggerType(types.LoggerType): 로거 타입
//   - settingOpts(...LogSettingOption): 로거 설정 옵션
//
// Example:
//
//	// zerolog 로거 생성
//	log := logger.NewLoggerWrapper(types.ZeroLog)
//	// zerolog with option 로거 생성
//	log := logger.NewLoggerWrapper(types.ZeroLog, logger.WithLevel(types.Debug))
//	// logrus 로거 생성
//	log := logger.NewLoggerWrapper(types.Logrus)
func NewWrapper(
	loggerType types.LoggerType,
	settingOpts ...options.LogSettingOption,
) (logger Logger) {
	settings := &options.LogSetting{
		Level:      types.Info,            // default log level
		TimeFormat: "2006-01-02 15:04:05", // default time format
		Output:     os.Stdout,             // default output
	}

	for _, opt := range settingOpts {
		opt(settings)
	}

	switch loggerType {
	case types.Logrus:
		logger = newLogrusLogger(*settings)
	case types.ZeroLog:
		logger = newZerologLogger(*settings)
	default:
		// no op
	}
	return
}

// FromContext : 컨텍스트에서 지정된 로거를 가져오는 메서드
//
// 만약 존재하지 않는 경우, 지정된 타입의 새로운 로거를 생성하여 반환
//
// Example:
//
//	// 컨텍스트에서 로거를 가져오는 예제
//	// ---------request handler 레벨에서 context에 로거를 등록---------
//	log := logger.NewLoggerWrapper(types.ZeroLog) // 로거 생성
//	ctx = log.WithContext(context.Background()) // 컨텍스트에 로거 등록
//	doSomething(ctx)
//	// ---------doSomething 함수 내부---------
//	log := logger.FromContext(types.ZeroLog, ctx)
func FromContext(ctx context.Context, loggerType types.LoggerType) (logger Logger) {
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
		return NewWrapper(loggerType)
	}

	return logger
}
