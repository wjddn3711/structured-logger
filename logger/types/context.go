package types

// LogContextKey : 로거 컨텍스트 키 타입
//   - 컨텍스트 키를 통해 context로 부터 등록된 로거를 가져올 수 있음
type LogContextKey string

const (
	// ZerologKey : zerolog 로거 컨텍스트 키
	ZerologKey LogContextKey = "zerolog-context-key"
	// LogrusKey : logrus 로거 컨텍스트 키
	LogrusKey LogContextKey = "logrus-context-key"
)
