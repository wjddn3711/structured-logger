package types

type LogContextKey string

const (
	ZerologKey LogContextKey = "zerolog-context-key"
	LogrusKey  LogContextKey = "logrus-context-key"
)
