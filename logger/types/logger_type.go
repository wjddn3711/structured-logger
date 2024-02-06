package types

type LoggerType string

const (
	Logrus  = LoggerType("logrus")
	Zap     = LoggerType("zap")
	ZeroLog = LoggerType("zerolog")
)
