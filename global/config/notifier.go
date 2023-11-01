package config

import (
	"backstage/common/conf"
	"backstage/global"
	"backstage/global/log"
	"github.com/BurntSushi/toml"
)

func Notifier(namespace, group, dataId, data string) {
	cf := &conf.NotifierConf{}
	err := toml.Unmarshal([]byte(data), cf)
	if err != nil {
		log.Error("Notifier failure, err = ", err.Error())
		return
	}
	if notify, exist := cf.Notify[global.ServiceName()]; exist {
		if id, ok := notify.Id[global.ServiceId()]; ok {
			if err := global.Notifier().Emit(id.CMD); err != nil {
				log.Error("Notifier failure, err = ", err.Error())
				return
			} else {
				log.Debug("Notifier.notify successfully, cmd = ", id.CMD)
				return
			}
		}
	}
}
