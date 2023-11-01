package rate

import "golang.org/x/time/rate"

type Options struct {
	capacity  int
	frequency rate.Limit
}

type Option func(*Options)

func WithCapacity(cap int) Option {
	return func(o *Options) {
		o.capacity = cap
	}
}

func WithFrequency(freq float64) Option {
	return func(o *Options) {
		o.frequency = rate.Limit(freq)
	}
}

