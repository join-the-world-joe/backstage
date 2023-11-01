package main

import (
	"backstage/abstract/registry"
	"backstage/common/http/scheme"
	"backstage/common/macro/config"
	service2 "backstage/common/macro/service"
	"backstage/common/tls"
	"backstage/service/backend"
	"flag"
	"fmt"
)

var _scheme string
var _id *string
var _host *string
var _httpPort *uint64
var _rpcPort *uint64
var _logFileName *string
var _logFilePath *string
var _version *string

var _nacosHost *string
var _nacosHttpPort *uint64
var _nacosGrpcPort *uint64
var _nacosNamespaceId *string
var _nacosLogLevel *string
var _nacosCachePath *string
var _nacosLogPath *string
var _nacosGroup *string

// .\main.exe -id="1" -version="1.0" -host="172.20.10.6" -rpc_port=11011 -nacos_host="192.168.130.129" -nacos_namespace_id="Test" -nacos_group="Service"
// .\main.exe -id="2" -version="1.0" -host="172.20.10.6" -rpc_port=11012 -nacos_host="192.168.130.129" -nacos_namespace_id="Test" -nacos_group="Service"
// .\main.exe -id="3" -version="1.0" -host="172.20.10.6" -rpc_port=11013 -nacos_host="192.168.130.129" -nacos_namespace_id="Test" -nacos_group="Service"

func init() {
	_id = flag.String("id", "", "the id of server")
	_host = flag.String("host", "172.20.10.4", "the ip address of service")
	_httpPort = flag.Uint64("http_port", uint64(0), "the http port of service")
	_rpcPort = flag.Uint64("rpc_port", uint64(0), "the rpc port of service")
	_version = flag.String("version", "", "the version of service")

	_nacosGroup = flag.String("nacos_group", "", "the group of registry")
	_nacosHost = flag.String("nacos_host", "192.168.130.129", "the ip address of nacos server")
	_nacosNamespaceId = flag.String("nacos_namespace_id", "", "the id of namespace of nacos for both registry and config")
	_logFileName = flag.String("log_file_name", "", "the file name of log")
	_nacosLogLevel = flag.String("nacos_log_level", "debug", "log level")
	_nacosHttpPort = flag.Uint64("nacos_http_port", uint64(8848), "the http port of nacos server")
	_nacosGrpcPort = flag.Uint64("nacos_grpc_port", uint64(9848), "the grpc port of nacos server")
	_logFilePath = flag.String("log_file_path", "D:\\Projects\\github\\backstage\\logs", "the file path of log")
	_nacosCachePath = flag.String("nacos_cache_path", "nacos_cache", "the cache path of nacos")
	_nacosLogPath = flag.String("nacos_log_path", "nacos_log", "the log path of nacos")

	flag.Parse()

	_scheme = func() string {
		if *_httpPort == uint64(tls.Port) {
			return scheme.HTTPS
		}
		return scheme.HTTP
	}()

	*_logFileName = func() string {
		if *_logFileName == "" {
			return fmt.Sprintf("report-backend-%v", *_id)
		}
		return *_logFilePath
	}()
}

func main() {
	service, err := backend.NewService(
		backend.WithId(*_id),
		backend.WithHost(*_host),
		backend.WithHTTPPort(*_httpPort),
		backend.WithRPCPort(*_rpcPort),
		backend.WithVersion(*_version),
		backend.WithNacosScheme(_scheme),
		backend.WithNacosHost(*_nacosHost),
		backend.WithNacosGroup(*_nacosGroup),
		backend.WithNacosLogLevel(*_nacosLogLevel),
		backend.WithLogFilePath(*_logFilePath),
		backend.WithLogFileName(*_logFileName),
		backend.WithNacosLogDir(*_nacosLogPath),
		backend.WithNacosHttpPort(*_nacosHttpPort),
		backend.WithNacosGrpcPort(*_nacosGrpcPort),
		backend.WithNacosCacheDir(*_nacosCachePath),
		backend.WithRegisterTTL(registry.DefaultTTL),
		backend.WithNacosNamespaceId(*_nacosNamespaceId),
		backend.WithConfig(config.RPCGroup, config.ServerDataId),
		backend.WithConfig(config.RouteGroup, config.ProtocolDataId),
		backend.WithConfig(config.ServiceGroup, service2.Backend),
		backend.WithRegisterInterval(registry.DefaultInterval),
		backend.WithConfig(config.ComponentGroup, config.BrokerDataId),
		backend.WithConfig(config.ComponentGroup, config.MySQLDataId),
		backend.WithConfig(config.ComponentGroup, config.RedisDataId),
		backend.WithConfig(config.ComponentGroup, config.MongoDBDataId),
		backend.WithConfig(config.NotifierGroup, service2.Backend),
		backend.WithConfig(config.BackendGroup, config.RBACDataId),
		backend.WithConfig(config.SecurityGroup, config.GracefulShutdownDataId),
		backend.WithNotifyBufferSize(service2.DefaultNotifyBufferSize),
		backend.WithForwardChannelBufferSize(service2.DefaultForwardChannelBufferSize),
	)
	if err != nil {
		panic(err)
	}

	if err = service.Init(); err != nil {
		panic(err)
	}

	if err = service.Start(); err != nil {
		panic(err)
	}

	service.Run()
}
