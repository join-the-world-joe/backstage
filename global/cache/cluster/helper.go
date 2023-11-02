package cluster

import (
	"backstage/common/conf"
	"github.com/go-redis/redis/v8"
)

func GetClient(cf *conf.CacheConf, which string) (*redis.ClusterClient, error) {
	return getRedisClusterClient(cf, which)
}
