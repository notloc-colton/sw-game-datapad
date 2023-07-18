package logger

import "go.uber.org/zap"

func init() {
	if devLog, err := zap.NewDevelopment(zap.AddCallerSkip(2)); err != nil || devLog == nil {
		panic("could not initialize logger")
	} else {
		logger = *devLog.Sugar()
	}
}
