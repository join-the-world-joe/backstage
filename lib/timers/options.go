package timers

import "time"

type Options struct {
	// for timers
	refreshInterval time.Duration
	offset          time.Duration
	eventBufferSize      int

	// for timer
	id       string
	duration time.Duration
	Loop     int
}

type Option func(*Options)


func WithEventBufferSize(bufferSize int) Option {
	return func(o *Options) {
		o.eventBufferSize = bufferSize
	}
}

func WithOffset(offset time.Duration) Option {
	return func(o *Options) {
		o.offset = offset
	}
}

func WithRefreshInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.refreshInterval = interval
	}
}

func WithId(id string) Option {
	return func(o *Options) {
		o.id = id
	}
}

func WithLoop(loop int) Option {
	return func(o *Options) {
		o.Loop = loop
	}
}

func WithInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.duration = interval
	}
}