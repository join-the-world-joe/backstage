package rate

import (
	"backstage/abstract/limiter"
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

const (
	Name = "Rate Limiter"
)

type _limiter struct {
	opts        *Options
	rateLimiter *rate.Limiter
}

func NewLimiter(opts ...Option) (limiter.Limiter, error) {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	if options.frequency <= 0 || options.capacity <= 0 {
		return nil, fmt.Errorf("zero value, frequency[%v] capacity[%v]",
			options.frequency, options.capacity)
	}

	return &_limiter{
		opts:        &options,
		rateLimiter: rate.NewLimiter(options.frequency, options.capacity),
	}, nil
}

func (p *_limiter) Name() string {
	return Name
}

func (p *_limiter) Take() time.Time {
	p.rateLimiter.Wait(context.Background())
	return time.Now()
}

func (p *_limiter) Allow() bool {
	return p.rateLimiter.Allow()
}

func (p *_limiter) Destroy() {
}
