package cluster

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

var g_clustered_redis_map sync.Map // it holds all connected redis

func store(unique string, clusterClient *redis.ClusterClient) {
	g_clustered_redis_map.Store(unique, clusterClient)
}

func load(unique string) (*redis.ClusterClient, error) {
	value, ok := g_clustered_redis_map.Load(unique)
	if ok {
		return value.(*redis.ClusterClient), nil
	}
	return nil, errors.New(fmt.Sprintf("%s doesn't exist", unique))
}

func Dump() {
	a := g_clustered_redis_map
	fmt.Println("a = ", a)
}
