package fifo

import (
	"backstage/abstract/fifo"
	"fmt"
	"time"
)

const (
	Name               = "Channel FIFO"
	DefaultPushTimeout = time.Duration(50) * time.Millisecond
	DefaultPopTimeout  = time.Duration(50) * time.Millisecond
	DefaultBufferSize  = 10
)

type _FIFO struct {
	opts        *Options
	pushTimeout time.Duration
	popTimeout  time.Duration
	bufferSize  int
	channel     chan interface{}
}

func NewFIFO(opts ...Option) fifo.FIFO {
	bufferSize := DefaultBufferSize
	pushTimeout := DefaultPushTimeout
	popTimeout := DefaultPopTimeout

	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if options.pushTimeout.Milliseconds() > 0 {
		pushTimeout = options.pushTimeout
	}

	if options.popTimeout.Milliseconds() > 0 {
		popTimeout = options.popTimeout
	}

	if options.bufferSize > 0 {
		bufferSize = options.bufferSize
	}

	return &_FIFO{
		opts:        &options,
		bufferSize:  bufferSize,
		pushTimeout: pushTimeout,
		popTimeout:  popTimeout,
		channel:     make(chan interface{}, bufferSize),
	}
}

func (p *_FIFO) Name() string {
	return Name
}

func (p *_FIFO) Push(any interface{}) error {
	select {
	case p.channel <- any:
		return nil
	case <-time.After(p.pushTimeout):
		return fmt.Errorf("push timeout")
	}
}

func (p *_FIFO) Pop() (interface{}, error) {
	for {
		select {
		case temp := <-p.channel:
			return temp, nil
		case <-time.After(p.popTimeout):
			return nil, fmt.Errorf("pop timeout")
		}
	}
}

func (p *_FIFO) Channel() <-chan interface{} {
	return p.channel
}

func (p *_FIFO) Destroy() {
	close(p.channel)
}
