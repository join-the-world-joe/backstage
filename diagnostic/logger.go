package diagnostic

import (
	"backstage/global/log"
	"backstage/plugin/logger/zap"
)

func SetupLogger() {
	logFilePath := "D:\\Projects\\github\\backstage\\logs"
	logFileName := "test.log"
	// for logger
	lg, err := zap.NewLogger(
		zap.WithLevel(-1),
		zap.WithFilePath(logFilePath),
		zap.WithFileName(logFileName),
		zap.WithCallerSkip(2),
	)
	if err != nil {
		panic(err)
	}
	log.SetLogger(lg)
}
