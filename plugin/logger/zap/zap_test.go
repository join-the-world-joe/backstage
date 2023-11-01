package zap

import (
	"backstage/abstract/logger"
	"github.com/jonboulle/clockwork"
	"testing"
	"time"
)

func TestLogMessage(t *testing.T) {
	lg, err := NewLogger(
		WithFileName("DynamicConfig"),
		WithLevel(-1), // -1, Debug; 0, Info; 1, Warn; 2, Error
		WithRotationTime(time.Second*65),
		WithCallerSkip(1),
		WithMaxAge(time.Hour*24*365),
	)
	if err != nil {
		t.Error(err)
		return
	}

	for {
		lg.Log(logger.Info, "info: hello, world!")
		lg.Log(logger.Debug, "debug: hello, world!")
		lg.Log(logger.Warn, "warn: hello, world!")
		lg.Log(logger.Error, "error: hello, world!")
		//lg.Log(logger.Fatal, "fatal: hello, world!")
		time.Sleep(time.Second * 5)
	}
}

func TestRotation(t *testing.T) {
	dummyTime := time.Now().Add(-7 * 24 * time.Hour)
	dummyTime = dummyTime.Add(time.Duration(-1 * dummyTime.Nanosecond()))
	clock := clockwork.NewFakeClockAt(dummyTime)

	go func() {
		for {
			clock.Advance(time.Duration(24 * time.Hour))
			t.Log("clock: ", clock.Now())
			time.Sleep(time.Second * 5)
		}
	}()

	lg, err := NewLogger(
		WithFileName("DynamicConfig"),
		WithLevel(-1),
		WithRotationTime(time.Second*65),
		WithCallerSkip(1),
		WithMaxAge(time.Hour*24*365),
		WithClock(clock),
	)
	if err != nil {
		t.Error(err)
		return
	}

	for {
		lg.Log(logger.Info, "info: hello, world!")
		lg.Log(logger.Debug, "debug: hello, world!")
		lg.Log(logger.Warn, "warn: hello, world!")
		lg.Log(logger.Error, "error: hello, world!")
		//lg.Log(logger.Fatal, "fatal: hello, world!")
		time.Sleep(time.Second * 5)
	}
}
