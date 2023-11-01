package config

import (
	"backstage/common/conf"
	"backstage/global/log"
	"github.com/BurntSushi/toml"
)

var _RBACConf *conf.RBACConf

func SetRBACConf(cf *conf.RBACConf) {
	_RBACConf = cf
}

func RBACConf() *conf.RBACConf {
	return _RBACConf
}

func RBAC(namespace, group, dataId, data string) {
	cf := &conf.RBACConf{}
	err := toml.Unmarshal([]byte(data), cf)
	if err != nil {
		log.Error("RBAC failure, err = ", err.Error())
		return
	}
	SetRBACConf(cf)
	log.Debug("RBAC updated: ", data)
}
