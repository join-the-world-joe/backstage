package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

var g_redis_map sync.Map // it holds all connected redis

func store(unique string, client *redis.Client) {
	g_redis_map.Store(unique, client)
}

func load(unique string) (*redis.Client, error) {
	value, ok := g_redis_map.Load(unique)
	if ok {
		return value.(*redis.Client), nil
	}
	return nil, errors.New(fmt.Sprintf("%s doesn't exist", unique))
}

func Dump() {
	a := g_redis_map
	fmt.Println("a = ", a)
}