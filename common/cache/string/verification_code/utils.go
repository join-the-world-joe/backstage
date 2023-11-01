package verification_code

import (
	"backstage/global/cache/redis"
	"backstage/global/config"
	"context"
	"errors"
	"fmt"
	"strings"
)

func Create(behavior, countryCode, phoneNumber, value string) error {
	key := GetKey(behavior, countryCode, phoneNumber)
	client, err := redis.GetClient(config.CacheConf(), GetWhich(), GetDB())
	if err != nil {
		return err
	}

	_, err = client.Set(context.Background(), key, value, GetExpire()).Result()
	if err != nil {
		return err
	}

	if err = client.Expire(context.Background(), key, GetExpire()).Err(); err != nil {
		return err
	}

	return nil
}

func Check(behavior, countryCode, phoneNumber, code string) error {
	key := GetKey(behavior, countryCode, phoneNumber)
	client, err := redis.GetClient(config.CacheConf(), GetWhich(), GetDB())
	if err != nil {
		return err
	}

	cache, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if strings.Compare(cache, code) == 0 {
		return nil
	}

	return errors.New(fmt.Sprintf("strings.Compare(cache, code) ！= 0"))
}
