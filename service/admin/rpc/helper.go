package rpc

import (
	server2 "backstage/common/server"
	"backstage/global"
	"github.com/spf13/cast"
)

func Setup() error {
	return server2.NewServer(
		server2.WithHost(global.Host()),
		server2.WithPort(cast.ToString(global.RPCPort())),
		server2.WithServicePath(global.ServicePath()),
		server2.WithServer(&_server{}),
	)
}
