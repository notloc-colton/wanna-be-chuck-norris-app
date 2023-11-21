package logger

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelFatal
	LogLevelError
)