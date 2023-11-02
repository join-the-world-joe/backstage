package config

import (
	"backstage/common/conf"
	"backstage/global/log"
	"github.com/BurntSushi/toml"
)

var _RateLimiting *conf.RateLimitingConf
var _callback = func() {}

func SetRateLimiting(cf *conf.RateLimitingConf) {
	_RateLimiting = cf
	_callback()
}

func RateLimitingConf() *conf.RateLimitingConf {
	return _RateLimiting
}

func SetRateLimitingCallback(callback func()) {
	_callback = callback
}

func RateLimiting(namespace, group, dataId, data string) {
	cf := &conf.RateLimitingConf{}
	err := toml.Unmarshal([]byte(data), cf)
	if err != nil {
		log.Error("RateLimiting failure, err = ", err.Error())
		return
	}
	SetRateLimiting(cf)
	log.Debug("RateLimiting updated: ", data)
}
