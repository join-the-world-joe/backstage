package diagnostic

import (
	"go-micro-framework/global/log"
	"go-micro-framework/plugin/logger/zap"
)

func SetupLogger() {
	logFilePath := "D:\\Projects\\github\\go-micro-framework\\logs"
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
