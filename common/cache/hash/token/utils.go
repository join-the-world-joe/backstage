package token

import (
	"backstage/global/cache/redis"
	"backstage/global/config"
	"backstage/global/log"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func Create(countryCode, phoneNumber, userId, unique string) error {
	key := GetKey(unique)
	client, err := redis.GetClient(config.CacheConf(), GetWhich(), GetDB())
	if err != nil {
		return err
	}

	fields := map[string]interface{}{FCountryCode: countryCode, FPhoneNumber: phoneNumber, FUserId: userId}

	_, err = client.HSet(context.Background(), key, fields).Result()
	if err != nil {
		return err
	}

	if err = client.Expire(context.Background(), key, GetExpire()).Err(); err != nil {
		return err
	}

	return nil
}

func Get(countryCode, phoneNumber, unique string) (*Token, error) {
	key := GetKey(unique)
	client, err := redis.GetClient(config.CacheConf(), GetWhich(), GetDB())
	if err != nil {
		return nil, err
	}

	m, err := client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	if len(m) <= 0 {
		warning := fmt.Sprintf("CountryCode[%v], PhoneNumber[%v] token[%v], len(m) <= 0",
			countryCode, phoneNumber, unique,
		)
		log.Warn(warning)
		return nil, errors.New(fmt.Sprintf("Get failure, len(m) <= 0"))
	}

	bytes, err := json.Marshal(&m)
	if err != nil {
		return nil, err
	}

	token := &Token{}
	err = json.Unmarshal(bytes, token)
	if err != nil {
		return nil, err
	}

	if strings.Compare(countryCode, token.CountryCode) == 0 &&
		strings.Compare(phoneNumber, token.PhoneNumber) == 0 {
		return token, nil
	}

	return nil, errors.New("input mobile doesn't match the cached mobile")
}
