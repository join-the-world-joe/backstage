package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type Options struct {
	timeout time.Duration // for each call
	client  *redis.Client
}

type Option func(*Options)

func WithClient(client *redis.Client) Option {
	return func(o *Options) {
		o.client = client
	}
}
