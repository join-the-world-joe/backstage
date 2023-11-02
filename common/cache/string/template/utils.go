package template

import (
	"backstage/common/conf"
	"backstage/global/cache/redis"
	"context"
)

func Create(cf *conf.CacheConf, value string) error {
	client, err := redis.GetClient(cf, GetWhich(), GetDB())
	if err != nil {
		return err
	}

	_, err = client.Set(context.Background(), GetKey(), value, GetExpire()).Result()
	if err != nil {
		return err
	}

	if err = client.Expire(context.Background(), GetKey(), GetExpire()).Err(); err != nil {
		return err
	}

	return nil
}
