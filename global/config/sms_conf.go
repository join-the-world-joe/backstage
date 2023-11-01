package config

import (
	"backstage/common/conf"
	"backstage/global/log"
	"github.com/BurntSushi/toml"
)

var _SMSConf *conf.SMSConf

func SetSMSConf(cf *conf.SMSConf) {
	_SMSConf = cf
}

func SMSConf() *conf.SMSConf {
	return _SMSConf
}

func SMS(namespace, group, dataId, data string) {
	cf := &conf.SMSConf{}
	err := toml.Unmarshal([]byte(data), cf)
	if err != nil {
		log.Error("SMS failure, err = ", err.Error())
		return
	}
	SetSMSConf(cf)
	log.Debug("SMS updated: ", data)
}
