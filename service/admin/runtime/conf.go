package runtime

import (
	"backstage/service/admin/conf"
	"encoding/json"
)

var serviceConf *conf.ServiceConf

func SetServiceConf(cf *conf.ServiceConf) {
	serviceConf = cf
}

func DumpConfig() []byte {
	if bytes, err := json.Marshal(serviceConf); err == nil {
		return bytes
	}
	return []byte(``)
}
