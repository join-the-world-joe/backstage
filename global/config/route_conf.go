package config

import "backstage/common/conf"

var _RouteConf *conf.RouteConf

func SetRouteConf(cf *conf.RouteConf) {
	_RouteConf = cf
}

func RouteConf() *conf.RouteConf {
	return _RouteConf
}

func DownstreamProtocol() string {
	return _RouteConf.Downstream.Protocol
}

func UpstreamProtocol() string {
	return _RouteConf.Upstream.Protocol
}
