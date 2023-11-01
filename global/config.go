package global

import "backstage/abstract/config"

var _nacos_group string
var _config_handler config.Config

func SetNacosGroup(group string) {
	_nacos_group = group
}

func NacosGroup() string {
	return _nacos_group
}

func SetConfigHandler(handle config.Config) {
	_config_handler = handle
}

func ConfigHandler() config.Config {
	return _config_handler
}
