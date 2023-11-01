package log

import (
	"backstage/plugin/logger/zap"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	filePath := ""
	logFileName := "DynamicConfig"
	lg, err := zap.NewLogger(
		zap.WithLevel(-1),
		zap.WithFilePath(filePath),
		zap.WithFileName(logFileName),
		zap.WithCallerSkip(2),
	)
	if err != nil {
		t.Fatal(err)
	}
	SetLogger(lg)

	for {
		Info("info: hello, world!")
		Debug("debug: hello, world!")
		Warn("warn: hello, world!")
		Error("error: hello, world!")
		time.Sleep(time.Second * 1)
	}
}
