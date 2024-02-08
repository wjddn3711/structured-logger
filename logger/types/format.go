package types

// LogFormat : 로그 포맷 타입
//
// Example:
//
//	// JSON
//	logFormat := types.JSON
//	// 텍스트
//	logFormat := types.Text
type LogFormat string

const (
	// JSON : JSON
	JSON LogFormat = "json"
	// Text : 텍스트
	Text LogFormat = "text"
)
