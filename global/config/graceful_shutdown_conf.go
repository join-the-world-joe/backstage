package config

import "backstage/common/conf"

const (
	DefaultGracefulShutdownTimeout       = 60 * 60
	DefaultGracefulShutdownCheckInterval = 60
)

var _GracefulShutdownConf *conf.GracefulShutdownConf

func SetGracefulShutdownConf(cf *conf.GracefulShutdownConf) {
	_GracefulShutdownConf = cf
}

func GracefulShutdownConf() *conf.GracefulShutdownConf {
	return _GracefulShutdownConf
}

func GracefulShutdownTimeout() int {
	if _GracefulShutdownConf.GracefulShutdown.Timeout == 0 {
		return DefaultGracefulShutdownTimeout
	}
	return _GracefulShutdownConf.GracefulShutdown.Timeout
}

func GracefulShutdownCheckInterval() int {
	if _GracefulShutdownConf.GracefulShutdown.CheckInterval == 0 {
		return DefaultGracefulShutdownCheckInterval
	}
	return _GracefulShutdownConf.GracefulShutdown.CheckInterval
}
