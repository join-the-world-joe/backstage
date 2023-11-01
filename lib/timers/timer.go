package timers

import (
	"go-micro-framework/abstract/timers"
	"time"
)

func NewTimer(opts ...Option) *timers.Timer {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return &timers.Timer{
		Id:       options.id,
		Loop:     options.Loop,
		LastTime: time.Now(),
		Duration: options.duration,
	}
}
