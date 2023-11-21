package logger

import "go.uber.org/zap"

func init() {
	//TODO: Will need to configure for prod environment
	//TODO: Mock logger for when running unit tests
	if devLog, err := zap.NewDevelopment(zap.AddCallerSkip(2)); err != nil || devLog == nil {
		panic("could not initialize logger")
	} else {
		logger = *devLog.Sugar()
	}
}
