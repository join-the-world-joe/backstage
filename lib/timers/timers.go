package timers

import (
	"context"
	"errors"
	"fmt"
	"go-micro-framework/abstract/timers"
	"sync"
	"time"
)

const (
	Name                   = "Timer Manager"
	DefaultRefreshInterval = time.Duration(100) * time.Millisecond
	DefaultEventBufferSize = 10
	DefaultOffset          = time.Duration(0)
)

type _timers struct {
	opts   *Options
	event  chan *timers.Timer
	timers sync.Map

	offset          time.Duration
	refreshInterval time.Duration
	eventBufferSize int

	ctx    context.Context
	cancel context.CancelFunc
}

func NewTimers(opts ...Option) timers.Timers {

	offset := DefaultOffset
	refreshInterval := DefaultRefreshInterval
	eventBufferSize := DefaultEventBufferSize

	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if options.refreshInterval.Milliseconds() > 0 {
		refreshInterval = options.refreshInterval
	}

	if options.eventBufferSize > 0 {
		eventBufferSize = options.eventBufferSize
	}

	if options.offset.Milliseconds() > 0 {
		offset = options.offset
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &_timers{
		ctx:             ctx,
		cancel:          cancel,
		opts:            &options,
		offset:          offset,
		eventBufferSize: eventBufferSize,
		refreshInterval: refreshInterval,
		event:           make(chan *timers.Timer, eventBufferSize),
	}
}

func (p *_timers) Name() string {
	return Name
}

func (p *_timers) AddTimer(timer *timers.Timer) error {
	if timer == nil {
		return errors.New("timer == nil")
	}
	if _, err := p.load(timer.Id); err == nil {
		return err
	}
	p.store(timer.Id, timer)
	return nil
}

func (p *_timers) RemoveTimer(id string) {
	p.delete(id)
}

func (p *_timers) Wait() <-chan *timers.Timer {
	return p.event
}

func (p *_timers) Start() timers.Timers {
	go func() {
		defer func() {
			recover()
		}()
		for {
			select {
			case <-p.ctx.Done():
				return
			case <-time.After(p.refreshInterval):
				now := time.Now().Add(p.offset)
				p.timers.Range(func(key, value interface{}) bool {
					timer, ok := value.(*timers.Timer)
					if !ok { // TODO: something wrong here
						fmt.Println("value.(*timers.Timer) is not ok")
						return true
					}
					if now.Sub(timer.LastTime) < timer.Duration {
						return true
					}
					timer.Done++
					timer.LastTime = time.Now()
					p.event <- timer
					if timer.Loop > 0 && timer.Done >= timer.Loop {
						p.RemoveTimer(timer.Id)
					}
					return true
				})
			}
		}
	}()

	return p
}

func (p *_timers) Destroy() {
	p.cancel()
	p.timers.Range(
		func(k, v interface{}) bool {
			p.timers.Delete(k)
			return true
		},
	)
	close(p.event)
}

func (p *_timers) store(id string, timer *timers.Timer) {
	p.timers.Store(id, timer)
}

func (p *_timers) load(id string) (*timers.Timer, error) {
	value, ok := p.timers.Load(id)
	if ok {
		return value.(*timers.Timer), nil
	}
	return nil, errors.New(fmt.Sprintf("no such timer %s", id))
}

func (p *_timers) delete(id string) {
	p.timers.Delete(id)
}
