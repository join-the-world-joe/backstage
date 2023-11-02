package runtime

import (
	"backstage/global/log"
	"backstage/service/frontend_gateway/conf"
	"encoding/json"
)

var serviceConf *conf.ServiceConf

func SetServiceConf(cf *conf.ServiceConf) {
	serviceConf = cf
}

func DumpConfig() {
	if bytes, err := json.Marshal(serviceConf); err == nil {
		log.WarnF("DumpConfig: \n %v", string(bytes))
	}
}
