package options

// LogEntry 로그 엔트리 필드 타입
//   - ToFields(): 구조체를 map[string]interface{} 형태로 변환하는 메서드
type LogEntry interface {
	// ToFields : 구조체를 map[string]interface{} 형태로 변환하는 메서드
	ToFields() map[string]interface{}
}

// EntryOption 로깅을 위한 로그 엔트리 옵션 타입
type EntryOption func(entry *Entry)

// Entry 로깅을 위한 로그 엔트리 타입
//   - message(string): 로그 메시지
//   - fields(LogEntry): 로그 필드 (구조체)
type Entry struct {
	Message string
	Fields  LogEntry
}

// WithMessage 로그 메시지를 등록하는 옵션
//   - message(string): 로그 메시지
//
// Example:
//
//	// 로그 메시지를 등록
//	log.Info(options.WithMessage("info message"))
//	// output: {"message":"info message"}
func WithMessage(message string) func(entry *Entry) {
	return func(entry *Entry) {
		entry.Message = message
	}
}

// WithFields 로그 필드를 등록하는 옵션
//   - fields(LogEntry): 로그 필드 (구조체)
//
// Example:
//
//	 // 로그 필드를 등록
//		entry := logEntry{
//			StartTime:   "2021-01-01T00:00:00Z",
//			EndTime:     "2021-01-01T00:00:01Z",
//			Elapsed:     1000,
//			StatusCode:  200,
//		}
//		log.Info(options.WithFields(entry))
//		// output: {"start_time":"2021-01-01T00:00:00Z","end_time":"2021-01-01T00:00:01Z","elapsed":1000,"status_code":200}
func WithFields(fields LogEntry) func(entry *Entry) {
	return func(entry *Entry) {
		entry.Fields = fields
	}
}
