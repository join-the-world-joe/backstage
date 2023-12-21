package config

import (
	"backstage/common/conf"
	"backstage/global/log"
	"github.com/BurntSushi/toml"
)

var _OSSConf *conf.OSSConf

func SetOSSConf(cf *conf.OSSConf) {
	_OSSConf = cf
}

func OSSConf() *conf.OSSConf {
	return _OSSConf
}

func OSS(namespace, group, dataId, data string) {
	cf := &conf.OSSConf{}
	err := toml.Unmarshal([]byte(data), cf)
	if err != nil {
		log.Error("OSS failure, err = ", err.Error())
		return
	}
	SetOSSConf(cf)
	log.Debug("OSS updated: ", data)
}
