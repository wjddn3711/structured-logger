package types

// LoggerType 로거 타입
//
// Example:
//
//	// logrus 로거 타입
//	logType := types.Logrus
//	// zap 로거 타입 (현재 미지원)
//	logType := types.Zap
//	// zerolog 로거 타입
//	logType := types.ZeroLog
//	// 로거 래퍼 생성
//	log := logger.NewWrapper(logType)
type LoggerType string

const (
	// Logrus : logrus 로거 타입
	Logrus = LoggerType("logrus")
	// Zap : zap 로거 타입 (현재 미지원)
	Zap = LoggerType("zap")
	// ZeroLog : zerolog 로거 타입
	ZeroLog = LoggerType("zerolog")
)
