package fifo

import "time"

type Options struct {
	bufferSize  int
	pushTimeout time.Duration
	popTimeout  time.Duration
}

type Option func(*Options)

func WithBufferSize(bufferSize int) Option {
	return func(o *Options) {
		o.bufferSize = bufferSize
	}
}

func WithPushTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.pushTimeout = timeout
	}
}

func WithPopTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.popTimeout = timeout
	}
}
