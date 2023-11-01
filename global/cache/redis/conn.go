package redis

import (
	"backstage/common/conf"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

var g_lock sync.Mutex

func getRedisClient(cf *conf.CacheConf, which string, db int64) (*redis.Client, error) {
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
		_client, err = connectToRedis(
			cf.Redis[which].Servers[0],
			cf.Redis[which].User,
			cf.Redis[which].Password,
			db,
		)
		if err != nil {
			return nil, err
		}
		store(which, _client)
		return _client, nil
	}

	return nil, errors.New(fmt.Sprintf("cann't find server info of %s", which))
}

func connectToRedis(server string, user, password string, db int64) (*redis.Client, error) {
	_client := redis.NewClient(
		&redis.Options{
			Addr:     server,
			Password: password,
			Username: user,
			DB:       int(db),
		},
	)

	return _client, _client.Ping(context.Background()).Err()
}
