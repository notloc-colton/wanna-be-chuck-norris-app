// This is a stub package that takes the place of a custom made logging solution
package logger

import "go.uber.org/zap"

type levelLogFunc func(args ...any)

var logger zap.SugaredLogger

func Log(level LogLevel, msg string, additionalInfo ...any) {
	switch level {
	case LogLevelError:
		logMsg(logger.Error, msg, additionalInfo...)
	case LogLevelFatal:
		logMsg(logger.Fatal, msg, additionalInfo...)
	default:
		logMsg(logger.Info, msg, additionalInfo...)
	}
}
func logMsg[LevelLog levelLogFunc](levelLog levelLogFunc, msg string, additionalInfo ...any) {
	levelLog(msg)
	if len(additionalInfo) > 0 {
		levelLog(additionalInfo)
	}
}
