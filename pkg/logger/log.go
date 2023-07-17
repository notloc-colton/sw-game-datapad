// This is a stub package that takes the place of a custom made logging solution
package logger

import "go.uber.org/zap"

type noMsgLogger func(args ...any)
type msgLogger func(msg string, args ...any)

var logger zap.SugaredLogger

func Log(level LogLevel, msg string, additionalInfo ...any) {
	switch level {
	case LogLevelError:
		logger.Errorw(msg, additionalInfo...)
	case LogLevelFatal:
		logger.Fatalw(msg, additionalInfo...)
	default:
		logger.Infow(msg, additionalInfo...)
	}
}
