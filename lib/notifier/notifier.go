package notifier

import (
	"backstage/abstract/notifier"
	"fmt"
	"time"
)

const (
	Name               = "Channel Notifier"
	DefaultPushTimeout = time.Duration(50) * time.Millisecond
	DefaultBufferSize  = 5
)

type _notifier struct {
	channel     chan string
	opts        *Options
	emitTimeout time.Duration
}

func NewNotifier(opts ...Option) (notifier.Notifier, error) {
	var bufferSize int
	var emitTimeout time.Duration
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if options.emitTimeout.Milliseconds() > 0 {
		emitTimeout = options.emitTimeout
	} else {
		emitTimeout = DefaultPushTimeout
	}

	if options.bufferSize > 0 {
		bufferSize = options.bufferSize
	} else {
		bufferSize = DefaultBufferSize
	}

	return &_notifier{
		opts:        &options,
		emitTimeout: emitTimeout,
		channel:     make(chan string, bufferSize),
	}, nil
}

func (p *_notifier) Name() string {
	return Name
}

func (p *_notifier) Wait() <-chan string {
	return p.channel
}

func (p *_notifier) Emit(signal string) error {
	select {
	case p.channel <- signal:
		return nil
	case <-time.After(p.emitTimeout):
		return fmt.Errorf("timeout")
	}
}

func (p *_notifier) Destroy() {
	close(p.channel)
}
