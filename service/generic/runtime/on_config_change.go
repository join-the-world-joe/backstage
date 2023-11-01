package runtime

import (
	"backstage/abstract/config"
	config3 "backstage/common/macro/config"
	"backstage/common/macro/service"
	config2 "backstage/global/config"
	"backstage/global/log"
	"backstage/service/generic/conf"
	"fmt"
	"github.com/BurntSushi/toml"
)

func Listener() func(cf config.Config) {
	return listen
}

func listen(cf config.Config) {
	cf.Subscribe(config.Parameter{
		Group:    config3.ServiceGroup,
		DataId:   service.Generic,
		OnChange: servant,
	})
	cf.Subscribe(config.Parameter{
		Group:    config3.BackendGroup,
		DataId:   config3.RBACDataId,
		OnChange: config2.RBAC,
	})
	cf.Subscribe(config.Parameter{
		Group:    config3.NotifierGroup,
		DataId:   service.Generic,
		OnChange: config2.Notifier,
	})
}

func servant(namespace, group, dataId, data string) {
	temp := new(conf.ServantConf)
	err := toml.Unmarshal([]byte(data), temp)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Warn(fmt.Sprintf("on_config_change.servant: %s", data))

	if temp.Servant.LogLevel != serviceConf.Servant.LogLevel {
		log.SetLevel(temp.Servant.LogLevel)
		log.Warn(fmt.Sprintf("Update log level from %s to %s", serviceConf.Servant.LogLevel, temp.Servant.LogLevel))
		serviceConf.Servant.LogLevel = temp.Servant.LogLevel
	}
}
