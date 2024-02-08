package logger

import (
	"context"

	"github.com/wjddn3711/structured-logger/logger/options"
)

// LogEntry 로그 엔트리 필드 타입
//   - ToFields(): 구조체를 map[string]interface{} 형태로 변환하는 메서드
type LogEntry interface {
	// ToFields : 구조체를 map[string]interface{} 형태로 변환하는 메서드
	ToFields() map[string]interface{}
}

// Logger 공통 로거 인터페이스
type Logger interface {
	// AddHook : 로거에 후크를 추가하는 메서드
	//   - hook(interface{}): 로거 후크
	// Reference:
	//   - https://github.com/rs/zerolog/blob/master/hook.go
	//   - https://pkg.go.dev/github.com/sirupsen/logrus/hooks/writer
	//
	// Example:
	//   // 로거에 후크를 추가
	//   log.AddHook(hook)
	AddHook(hook interface{})
	// WithContext : 컨텍스트에 로거를 등록하는 메서드
	//
	// Example:
	//   // request handler 레벨에서 context에 로거를 등록
	//   log := logger.NewLoggerWrapper(types.ZeroLog) // 로거 생성
	//   ctx = log.WithContext(context.Background()) // 컨텍스트에 로거 등록
	WithContext(ctx context.Context) context.Context
	// RegisterCommonField : 로거에 공통 필드를 등록하는 메서드
	//   - key(string): 필드 키
	//   - value(interface{}): 필드 값
	//
	// Example:
	//   // request handler 레벨에서 공통 필드(requestID)를 등록, 이후 같은 context를 사용하는 모든 로그에 requestID 필드가 등록됨
	//   reqID := ctx.Value("rid") // 공통 필드를 등록하기 위한 값
	//   log := logger.NewLoggerWrapper(types.ZeroLog) // 로거 생성
	//   log.RegisterCommonField("rid", reqID) // 공통 필드 등록
	//   ctx = log.WithContext(context.Background()) // 컨텍스트에 로거 등록
	RegisterCommonField(key string, value interface{})
	// RegisterCommonFields : 로거에 공통 필드들을 등록하는 메서드
	//   - fields(LogEntry): 로그 필드 (구조체)
	//
	// Example:
	//   // request handler 레벨에서 공통 필드들을 등록, 이후 같은 context를 사용하는 모든 로그에 공통 필드들이 등록됨
	//   entry := logEntry{ // 공통 필드를 등록하기 위한 값
	//     StartTime: "2021-01-01T00:00:00Z",
	//     reqID: "uuid",
	//   }
	//   log := logger.NewLoggerWrapper(types.ZeroLog) // 로거 생성
	//   log.RegisterCommonFields(entry) // 공통 필드들 등록
	//   ctx = log.WithContext(context.Background()) // 컨텍스트에 로거 등록
	RegisterCommonFields(fields options.LogEntry)
	// ApplyOption : 로그 엔트리 옵션을 적용하는 메서드
	//   - opts([]EntryOption): 로그 엔트리 옵션
	ApplyOption([]options.EntryOption)
	// Debug : 디버그 로그를 출력하는 메서드
	//   - opts(...EntryOption): 로그 엔트리 옵션
	//
	// Example:
	//   // 로그 메시지와 로그 필드를 함께 출력
	//   log.Debug(options.WithMessage("debug message"), options.WithFields(entry))
	//   // 로그 메시지만 출력
	//   log.Debug(options.WithMessage("debug message"))
	//   // 로그 필드만 출력
	//   log.Debug(options.WithFields(entry))
	//   // 기존 등록된 엔트리를 그대로 출력
	//   log.Debug()
	Debug(opts ...options.EntryOption)
	// Info : 정보 로그를 출력하는 메서드
	//   - opts(...EntryOption): 로그 엔트리 옵션
	//
	// Example:
	//   // 로그 메시지와 로그 필드를 함께 출력
	//   log.Info(options.WithMessage("info message"), options.WithFields(entry))
	//   // 로그 메시지만 출력
	//   log.Info(options.WithMessage("info message"))
	//   // 로그 필드만 출력
	//   log.Info(options.WithFields(entry))
	//   // 기존 등록된 엔트리를 그대로 출력
	//   log.Info()
	Info(opts ...options.EntryOption)
	// Warn : 경고 로그를 출력하는 메서드
	//   - opts(...EntryOption): 로그 엔트리 옵션
	//
	// Example:
	//   // 로그 메시지와 로그 필드를 함께 출력
	//   log.Warn(options.WithMessage("warn message"), options.WithFields(entry))
	//   // 로그 메시지만 출력
	//   log.Warn(options.WithMessage("warn message"))
	//   // 로그 필드만 출력
	//   log.Warn(options.WithFields(entry))
	//   // 기존 등록된 엔트리를 그대로 출력
	//   log.Warn()
	Warn(opts ...options.EntryOption)
	// Error : 에러 로그를 출력하는 메서드
	//   - opts(...EntryOption): 로그 엔트리 옵션
	//
	// Example:
	//   // 로그 메시지와 로그 필드를 함께 출력
	//   log.Error(options.WithMessage("error message"), options.WithFields(entry))
	//   // 로그 메시지만 출력
	//   log.Error(options.WithMessage("error message"))
	//   // 로그 필드만 출력
	//   log.Error(options.WithFields(entry))
	//   // 기존 등록된 엔트리를 그대로 출력
	//   log.Error()
	Error(opts ...options.EntryOption)
	// Fatal : 치명적인 에러 로그를 출력하는 메서드
	//   - opts(...EntryOption): 로그 엔트리 옵션
	//
	// Example:
	//   // 로그 메시지와 로그 필드를 함께 출력
	//   log.Fatal(options.WithMessage("fatal message"), options.WithFields(entry))
	//   // 로그 메시지만 출력
	//   log.Fatal(options.WithMessage("fatal message"))
	//   // 로그 필드만 출력
	//   log.Fatal(options.WithFields(entry))
	//   // 기존 등록된 엔트리를 그대로 출력
	//   log.Fatal()
	Fatal(opts ...options.EntryOption)
}
