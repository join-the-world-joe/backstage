package cluster

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-micro-framework/common/conf"
	"sync"
)

var g_lock sync.Mutex

func getRedisClusterClient(cf *conf.CacheConf, which string) (*redis.ClusterClient, error) {
	_client, err := load(which)
	if err == nil { // created
		return _client, nil
	}

	g_lock.Lock()
	defer g_lock.Unlock()

	if _, fExist := cf.Redis[which]; fExist { // try to find redis info by full_name
		_client, err = load(which)
		if err == nil { // created
			return _client, nil
		}
		_client, err = connectToClusteredRedis(
			cf.Redis[which].Servers,
			cf.Redis[which].User,
			cf.Redis[which].Password,
		)
		if err != nil {
			return nil, err
		}
		store(which, _client)
		return _client, nil
	}

	return nil, errors.New(fmt.Sprintf("cann't find server info of %s", which))
}

func connectToClusteredRedis(servers []string, user, password string) (*redis.ClusterClient, error) {
	_client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    servers,
		Password: password,
		Username: user,
	})

	return _client, _client.Ping(context.Background()).Err()
}
