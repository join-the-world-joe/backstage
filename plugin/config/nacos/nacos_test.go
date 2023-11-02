package nacos

import (
	"backstage/abstract/config"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
	"time"
)

func newConfigClient(t *testing.T) config.Config {
	scheme := "http"
	host := "192.168.130.129"
	namespaceId := "test"
	httpPort := 8848
	grpcPort := 9848
	if cc, err := NewConfig(
		WithNacosClientParam(vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				LogDir:              "./log",
				CacheDir:            "./cache",
				LogLevel:            "debug",
				TimeoutMs:           5000,
				NamespaceId:         namespaceId,
				NotLoadCacheAtStart: true,
			},
			ServerConfigs: []constant.ServerConfig{
				{
					Scheme:   scheme,
					IpAddr:   host,
					Port:     uint64(httpPort),
					GrpcPort: uint64(grpcPort),
				},
			},
		}),
	); err == nil {
		return cc
	} else {
		t.Fatal(err)
		return nil
	}
}

func TestPublish(t *testing.T) {
	cc := newConfigClient(t)
	err := cc.Publish(config.Parameter{
		Group:   "group",
		DataId:  "data-id-1",
		Content: "content-data-id-1",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	cc := newConfigClient(t)
	bytes, err := cc.Load(config.Parameter{
		Group:  "group",
		DataId: "data-id-1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestSubscribe(t *testing.T) {
	cc := newConfigClient(t)
	err := cc.Subscribe(config.Parameter{
		Group:  "group",
		DataId: "data-id-1",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("Namespace: ", namespace)
			fmt.Println("Group: ", group)
			fmt.Println("Data-ID: ", dataId)
			fmt.Println("Content: ", data)
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	cc.UnSubscribe(config.Parameter{
		Group:  "group",
		DataId: "data-id-1",
	})
	select {}
}

func TestUnSubscribe(t *testing.T) {
	cc := newConfigClient(t)
	err := cc.Subscribe(config.Parameter{
		Group:  "group",
		DataId: "data-id-1",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("Namespace: ", namespace)
			fmt.Println("Group: ", group)
			fmt.Println("Data-ID: ", dataId)
			fmt.Println("Content: ", data)
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("UnSubscribe now")
	if err = cc.UnSubscribe(config.Parameter{
		Group:  "group",
		DataId: "data-id-1",
	}); err != nil {
		t.Fatal(err)
	}
	select {}
}
