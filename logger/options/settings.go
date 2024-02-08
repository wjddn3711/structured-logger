package options

import (
	"io"

	"github.com/wjddn3711/structured-logger/logger/types"
)

// LogSetting : 로그 설정
type LogSetting struct {
	// Level : 로그 레벨
	//   - types.Debug: 디버그 레벨
	//   - types.Info: 정보 레벨
	//   - types.Warn: 경고 레벨
	//   - types.Error: 에러 레벨
	Level types.LogLevel
	// Output: 로그 출력
	//   - os.Stdout: 표준 출력
	//   - os.Stderr: 표준 에러
	//   - *lumberjack.Logger: 파일
	Output io.Writer
	// timeFormat: 시간 포맷
	TimeFormat string
}

// LogSettingOption 로그 설정을 위한 옵션 타입
//   - WithLevel: 로그 레벨을 설정하는 옵션 (default: info)
//   - WithOutput: 로그 출력 위치를 설정하는 옵션 (default: os.Stdout)
//   - WithTimeFormat: 로그의 시간 포맷을 설정하는 옵션 (default: "2006-01-02 15:04:05")
type LogSettingOption func(*LogSetting)

// WithLevel 로그 레벨을 설정하는 옵션
//
// Example:
//
//	// 디버그 레벨
//	level := "debug"
//	// 정보 레벨
//	level := "info"
//	// 경고 레벨
//	level := "warn"
//	// 에러 레벨
//	level := "error"
//	log := logger.NewLoggerWrapper(types.ZeroLog, logger.WithLevel(level))
func WithLevel(level types.LogLevel) LogSettingOption {
	return func(setting *LogSetting) {
		setting.Level = level
	}
}

// WithOutput 로그 출력 위치를 설정하는 옵션
//
// Example:
//
//	// 표준 출력
//	output := os.Stdout
//	// 표준 에러
//	output := os.Stderr
//	// 파일
//	output := &lumberjack.Logger{
//		Filename:   "log/test.log",
//		MaxSize:    1, // megabytes
//		MaxBackups: 3,
//		MaxAge:     28, // days
//	}
//	log := logger.NewLoggerWrapper(types.ZeroLog, logger.WithOutput(output))
func WithOutput(output io.Writer) LogSettingOption {
	return func(setting *LogSetting) {
		setting.Output = output
	}
}

// WithTimeFormat 로그의 시간 포맷을 설정하는 옵션
//   - timeFormat(string): 시간 포맷, 지정 하지 않을 경우 "2006-01-02 15:04:05"
//
// Example:
//
//	// 시간 포맷
//	timeFormat := "2006-01-02 15:04:05"
//	log := logger.NewLoggerWrapper(types.ZeroLog, logger.WithTimeFormat(timeFormat))
func WithTimeFormat(timeFormat string) LogSettingOption {
	return func(setting *LogSetting) {
		setting.TimeFormat = timeFormat
	}
}
