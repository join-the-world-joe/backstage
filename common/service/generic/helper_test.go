package generic

import (
	"backstage/abstract/logger"
	"backstage/common/payload"
	"backstage/global"
	"backstage/global/log"
	"backstage/plugin/logger/zap"
	nacos2 "backstage/plugin/registry/nacos"
	"context"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
	"time"
)

var lg = func() logger.Logger {
	logFilePath := "D:\\Projects\\github\\backstage\\logs"
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

func registry() {
	scheme := "http"
	host := "192.168.130.129"
	httpPort := uint64(8848)
	grpcPort := uint64(9848)
	nameSpaceId := "Test"

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
		panic(err)
	}
	global.SetRegistry(r)
}

func TestBreak(t *testing.T) {
	registry()
	req := &payload.FakeReq{}
	rsp := &payload.FakeRsp{}
	err := Break(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 3)
}
