package config

import "backstage/common/conf"

var _CacheConf *conf.CacheConf

func SetCacheConf(cf *conf.CacheConf) {
	_CacheConf = cf
}

func CacheConf() *conf.CacheConf {
	return _CacheConf
}
