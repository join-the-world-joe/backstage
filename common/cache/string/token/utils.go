package token

import (
	"backstage/global/cache/redis"
	"backstage/global/config"
	"context"
)

func Create(memberId, value string) error {
	key := GetKey(memberId)
	client, err := redis.GetClient(config.CacheConf(), GetWhich(), GetDB())
	if err != nil {
		return err
	}

	_, err = client.Set(context.Background(), key, value, GetExpire()).Result()
	if err != nil {
		return err
	}

	return nil
}

func Get(memberId string) (string, error) {
	key := GetKey(memberId)
	client, err := redis.GetClient(config.CacheConf(), GetWhich(), GetDB())
	if err != nil {
		return "", err
	}

	temp, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return temp, nil
}

func Expire(memberId string) error {
	key := GetKey(memberId)
	client, err := redis.GetClient(config.CacheConf(), GetWhich(), GetDB())
	if err != nil {
		return err
	}

	if err = client.Expire(context.Background(), key, GetExpire()).Err(); err != nil {
		return err
	}

	return nil
}
