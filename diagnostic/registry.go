package diagnostic

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go-micro-framework/global"
	nacos2 "go-micro-framework/plugin/registry/nacos"
)

func SetupRegistry() {
	scheme := "http"
	host := "192.168.130.129"
	httpPort := uint64(8848)
	grpcPort := uint64(9848)
	nameSpaceId := "Test"

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
