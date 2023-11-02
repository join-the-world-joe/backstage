package template

import (
	"backstage/common/conf"
	"backstage/global/cache/redis"
	"context"
)

func Create(cf *conf.CacheConf, id int, value string) error {
	client, err := redis.GetClient(cf, GetWhich(id), GetDB())
	if err != nil {
		return err
	}

	_, err = client.Set(context.Background(), GetKey(id), value, GetExpire()).Result()
	if err != nil {
		return err
	}

	if err = client.Expire(context.Background(), GetKey(id), GetExpire()).Err(); err != nil {
		return err
	}

	return nil
}
