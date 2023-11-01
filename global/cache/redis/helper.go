package redis

import (
	"backstage/common/conf"
	"github.com/go-redis/redis/v8"
)

func GetClient(cf *conf.CacheConf, which string, db int64) (*redis.Client, error) {
	return getRedisClient(cf, which, db)
}
