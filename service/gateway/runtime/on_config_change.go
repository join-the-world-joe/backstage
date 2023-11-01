package runtime

import (
	"backstage/abstract/config"
	config3 "backstage/common/macro/config"
	"backstage/common/macro/service"
	config2 "backstage/global/config"
	"backstage/global/log"
	"backstage/service/gateway/conf"
	"fmt"
	"github.com/BurntSushi/toml"
)

func Listener() func(cf config.Config) {
	return listen
}

func listen(cf config.Config) {
	cf.Subscribe(config.Parameter{
		Group:    config3.ServiceGroup,
		DataId:   service.Gateway,
		OnChange: servant,
	})
	cf.Subscribe(config.Parameter{
		Group:    config3.BackendGroup,
		DataId:   config3.RBACDataId,
		OnChange: config2.RBAC,
	})
	cf.Subscribe(config.Parameter{
		Group:    config3.NotifierGroup,
		DataId:   service.Gateway,
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

	serviceConf.Encryption.Enable = temp.Encryption.Enable
	serviceConf.Encryption.Algorithm = temp.Encryption.Algorithm

	serviceConf.Servant.QPS = temp.Servant.QPS
	serviceConf.Servant.IPLimit = temp.Servant.IPLimit
	serviceConf.Servant.WebsocketReadLimit = temp.Servant.WebsocketReadLimit
	serviceConf.Servant.WebsocketIdleTimeout = temp.Servant.WebsocketIdleTimeout
	serviceConf.Servant.WebsocketReadDeadline = temp.Servant.WebsocketReadDeadline
	serviceConf.Servant.WebsocketReadBufferSize = temp.Servant.WebsocketReadBufferSize
	serviceConf.Servant.WebsocketWriteBufferSize = temp.Servant.WebsocketWriteBufferSize
	serviceConf.Servant.WebsocketAuthReadDeadline = temp.Servant.WebsocketAuthReadDeadline

	serviceConf.Feedback.Enable = temp.Feedback.Enable
	serviceConf.Feedback.Major = temp.Feedback.Major
	serviceConf.Feedback.Minor = temp.Feedback.Minor
	serviceConf.Feedback.Message = temp.Feedback.Message
	serviceConf.Feedback.WaitForCloseInterval = temp.Feedback.WaitForCloseInterval
}
