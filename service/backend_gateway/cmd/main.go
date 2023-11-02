package main

import (
	"backstage/abstract/registry"
	"backstage/common/http/scheme"
	"backstage/common/macro/config"
	service2 "backstage/common/macro/service"
	"backstage/common/tls"
	"backstage/service/backend_gateway"
	"flag"
	"fmt"
)

var _scheme string
var _id *string
var _host *string
var _rpcPort *uint64
var _httpPort *uint64
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

// .\main.exe -id="1" -version="1.0" -host="" -http_port=10001 -rpc_port=11001 -nacos_host="172.20.10.6" -nacos_namespace_id="test" -nacos_group="Service"

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
			return fmt.Sprintf("report-backend_gateway-%v", *_id)
		}
		return *_logFilePath
	}()
}

func main() {
	service, err := backend_gateway.NewService(
		backend_gateway.WithId(*_id),
		backend_gateway.WithHost(*_host),
		backend_gateway.WithHTTPPort(*_httpPort),
		backend_gateway.WithRPCPort(*_rpcPort),
		backend_gateway.WithVersion(*_version),
		backend_gateway.WithNacosScheme(_scheme),
		backend_gateway.WithNacosHost(*_nacosHost),
		backend_gateway.WithNacosGroup(*_nacosGroup),
		backend_gateway.WithNacosLogLevel(*_nacosLogLevel),
		backend_gateway.WithLogFilePath(*_logFilePath),
		backend_gateway.WithLogFileName(*_logFileName),
		backend_gateway.WithNacosLogDir(*_nacosLogPath),
		backend_gateway.WithNacosHttpPort(*_nacosHttpPort),
		backend_gateway.WithNacosGrpcPort(*_nacosGrpcPort),
		backend_gateway.WithNacosCacheDir(*_nacosCachePath),
		backend_gateway.WithRegisterTTL(registry.DefaultTTL),
		backend_gateway.WithNacosNamespaceId(*_nacosNamespaceId),
		backend_gateway.WithConfig(config.RPCGroup, config.ServerDataId),
		backend_gateway.WithConfig(config.RouteGroup, config.ProtocolDataId),
		backend_gateway.WithConfig(config.ServiceGroup, service2.BackendGateway),
		backend_gateway.WithRegisterInterval(registry.DefaultInterval),
		backend_gateway.WithConfig(config.ComponentGroup, config.BrokerDataId),
		backend_gateway.WithConfig(config.NotifierGroup, service2.BackendGateway),
		backend_gateway.WithConfig(config.SecurityGroup, config.GracefulShutdownDataId),
		backend_gateway.WithConfig(config.SecurityGroup, config.RateLimitingDataId),
		backend_gateway.WithConfig(config.ComponentGroup, config.MongoDBDataId),
		backend_gateway.WithNotifyBufferSize(service2.DefaultNotifyBufferSize),
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
