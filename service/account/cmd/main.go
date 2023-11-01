package main

import (
	"backstage/abstract/registry"
	"backstage/common/http/scheme"
	"backstage/common/macro/config"
	service2 "backstage/common/macro/service"
	"backstage/common/tls"
	"backstage/service/account"
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

// .\main.exe -id="1" -version="1.0" -host="172.20.10.6" -rpc_port=11008 -nacos_host="192.168.130.129" -nacos_namespace_id="Test" -nacos_group="Service"
// .\main.exe -id="2" -version="1.0" -host="172.20.10.6" -rpc_port=11009 -nacos_host="192.168.130.129" -nacos_namespace_id="Test" -nacos_group="Service"
// .\main.exe -id="3" -version="1.0" -host="172.20.10.6" -rpc_port=11010 -nacos_host="192.168.130.129" -nacos_namespace_id="Test" -nacos_group="Service"

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
			return fmt.Sprintf("report-account-%v", *_id)
		}
		return *_logFilePath
	}()
}

func main() {
	service, err := account.NewService(
		account.WithId(*_id),
		account.WithHost(*_host),
		account.WithHTTPPort(*_httpPort),
		account.WithRPCPort(*_rpcPort),
		account.WithVersion(*_version),
		account.WithNacosScheme(_scheme),
		account.WithNacosHost(*_nacosHost),
		account.WithNacosGroup(*_nacosGroup),
		account.WithNacosLogLevel(*_nacosLogLevel),
		account.WithLogFilePath(*_logFilePath),
		account.WithLogFileName(*_logFileName),
		account.WithNacosLogDir(*_nacosLogPath),
		account.WithNacosHttpPort(*_nacosHttpPort),
		account.WithNacosGrpcPort(*_nacosGrpcPort),
		account.WithNacosCacheDir(*_nacosCachePath),
		account.WithRegisterTTL(registry.DefaultTTL),
		account.WithNacosNamespaceId(*_nacosNamespaceId),
		account.WithConfig(config.RPCGroup, config.ServerDataId),
		account.WithConfig(config.RouteGroup, config.ProtocolDataId),
		account.WithConfig(config.ServiceGroup, service2.Account),
		account.WithRegisterInterval(registry.DefaultInterval),
		account.WithConfig(config.ComponentGroup, config.BrokerDataId),
		account.WithConfig(config.ComponentGroup, config.MySQLDataId),
		account.WithConfig(config.ComponentGroup, config.RedisDataId),
		account.WithConfig(config.ComponentGroup, config.MongoDBDataId),
		account.WithConfig(config.NotifierGroup, service2.Account),
		account.WithConfig(config.SecurityGroup, config.GracefulShutdownDataId),
		account.WithNotifyBufferSize(service2.DefaultNotifyBufferSize),
		account.WithForwardChannelBufferSize(service2.DefaultForwardChannelBufferSize),
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
