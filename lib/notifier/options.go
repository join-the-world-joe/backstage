package notifier

import "time"

type Options struct {
	bufferSize  int
	emitTimeout time.Duration
}

type Option func(*Options)

func WithEmitTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.emitTimeout = timeout
	}
}

func WithBufferSize(bufferSize int) Option {
	return func(o *Options) {
		o.bufferSize = bufferSize
	}
}