package global

import (
	"backstage/abstract/logger"
	"backstage/global/log"
	"backstage/plugin/logger/zap"
	nacos2 "backstage/plugin/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
	"time"
)

var lg = func() logger.Logger {
	logFilePath := "D:\\Projects\\app\\v2\\log"
	logFileName := "test.log"
	// for logger
	lg, err := zap.NewLogger(
		zap.WithLevel(-1),
		zap.WithFilePath(logFilePath),
		zap.WithFileName(logFileName),
		zap.WithCallerSkip(2),
	)
	if err != nil {
		panic(err)
	}
	return lg
}()

func TestSubscribe(t *testing.T) {
	scheme := "http"
	host := "192.168.130.129"
	httpPort := uint64(8848)
	grpcPort := uint64(9848)
	srvName := "Generic"
	nameSpaceId := "test"

	log.SetLogger(lg)

	r, err := nacos2.NewRegistry(
		nacos2.WithNacosClientParam(
			vo.NacosClientParam{
				ClientConfig: &constant.ClientConfig{
					LogDir:               "nacos_log",
					CacheDir:             "nacos_cache",
					LogLevel:             "debug",
					TimeoutMs:            60 * 1000, // Healthy timeout
					NamespaceId:          nameSpaceId,
					NotLoadCacheAtStart:  true,
					UpdateCacheWhenEmpty: true,
				},
				ServerConfigs: []constant.ServerConfig{
					{
						Scheme:   scheme,
						IpAddr:   host,
						Port:     httpPort,
						GrpcPort: grpcPort,
					},
				},
			},
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	SetRegistry(r)

	subscribe(srvName)

	for {
		time.Sleep(time.Second)
		srv, err := SelectService(srvName)
		if err != nil {
			t.Log("err: ", err)
			continue
		}
		t.Log("srv: ", srv)
	}
}
