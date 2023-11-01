package cluster

import (
	"github.com/go-redis/redis/v8"
	"go-micro-framework/common/conf"
)

func GetClient(cf *conf.CacheConf, which string) (*redis.ClusterClient, error) {
	return getRedisClusterClient(cf, which)
}
