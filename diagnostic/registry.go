package diagnostic

import (
	"backstage/global"
	nacos2 "backstage/plugin/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func SetupRegistry() {
	scheme := "http"
	host := "119.23.224.221"
	httpPort := uint64(8848)
	grpcPort := uint64(9848)
	nameSpaceId := "test"

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
