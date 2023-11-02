package nacos

import (
	"backstage/abstract/registry"
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
	"time"
)

func newConfigClient(t *testing.T) registry.Registry {
	scheme := "http"
	host := "192.168.130.129"
	namespaceId := "Test"
	httpPort := 8848
	grpcPort := 9848
	r, err := NewRegistry(
		WithNacosClientParam(
			vo.NacosClientParam{
				ClientConfig: &constant.ClientConfig{
					LogDir:               "./nacos_log",
					CacheDir:             "./nacos_cache",
					LogLevel:             "debug",
					TimeoutMs:            30000, // Healthy timeout
					NamespaceId:          namespaceId,
					NotLoadCacheAtStart:  true,
					UpdateCacheWhenEmpty: true,
				},
				ServerConfigs: []constant.ServerConfig{
					{
						Scheme:   scheme,
						IpAddr:   host,
						Port:     uint64(httpPort),
						GrpcPort: uint64(grpcPort),
					},
				},
			},
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func TestRegister(t *testing.T) {
	r := newConfigClient(t)
	service := &registry.Service{
		Id:      "1",
		Ip:      "192.168.0.70",
		Port:    10001,
		Name:    "Generic",
		Group:   "Service",
		Version: "v2.0",
	}
	if err := r.Register(service); err != nil {
		t.Fatal(err)
	}
}

func TestDeRegister(t *testing.T) {
	r := newConfigClient(t)
	service := &registry.Service{
		Id:      "1",
		Ip:      "172.20.10.6",
		Port:    11001,
		Name:    "Gateway",
		Group:   "Service",
		Version: "1.0",
	}
	if err := r.DeRegister(service); err != nil {
		t.Fatal(err)
	}
	r.Destroy()
}

func TestGetServiceList(t *testing.T) {
	r := newConfigClient(t)

	for {
		list, err := r.ListServices(&registry.Service{
			Group: "Service",
			Name:  "Gateway",
		})
		if err != nil {
			t.Fatal(err)
		}

		//r.Destroy()

		bytes, err := json.Marshal(&list)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("list: ", string(bytes))
		time.Sleep(time.Second * 3)
	}
}

func TestDestroy(t *testing.T) {
	r := newConfigClient(t)
	r.Destroy()
}
