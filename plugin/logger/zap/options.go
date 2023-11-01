package zap

import (
	"github.com/jonboulle/clockwork"
	"go.uber.org/zap/zapcore"
	"time"
)

type Options struct {
	callerSkip int
	filePath   string
	fileName   string
	maxAge     time.Duration

	level        zapcore.Level
	rotationTime time.Duration

	clock clockwork.FakeClock
}

type Option func(*Options)

func WithClock(clock clockwork.FakeClock) Option {
	return func(o *Options) {
		o.clock = clock
	}
}

func WithCallerSkip(callerSkip int) Option {
	return func(o *Options) {
		o.callerSkip = callerSkip
	}
}

func WithMaxAge(maxAge time.Duration) Option {
	return func(o *Options) {
		o.maxAge = maxAge
	}
}

func WithFilePath(path string) Option {
	return func(o *Options) {
		o.filePath = path
	}
}

func WithFileName(fileName string) Option {
	return func(o *Options) {
		o.fileName = fileName
	}
}

func WithLevel(lvl int8) Option {
	return func(o *Options) {
		o.level = zapcore.Level(lvl)
	}
}

func WithRotationTime(rotationTime time.Duration) Option {
	return func(o *Options) {
		o.rotationTime = rotationTime
	}
}