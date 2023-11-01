package timers

import "time"

type Timer struct {
	Id       string
	Duration time.Duration // in milliseconds
	Loop     int
	Done     int
	LastTime time.Time
}

type Timers interface {
	Name() string
	AddTimer(*Timer) error
	RemoveTimer(string)
	Wait() <-chan *Timer
	Start() Timers
	Destroy()
}
