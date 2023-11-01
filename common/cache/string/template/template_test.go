package template

import (
	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	"go-micro-framework/common/conf"
	"testing"
)

var cluster_conf = `
	[Redis.test]
	Name = "Clustered Redis Server"
	Servers = ["192.168.130.128:7001", "192.168.130.128:7002", "192.168.130.128:7003"]
	Password = "123456"
	[Redis.test.Param]
	db = 0
`

func TestCreate(t *testing.T) {
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(cluster_conf), &cf); err != nil {
		t.Fatal(err)
	}

	if err := Create(cf, uuid.New().String()); err != nil {
		t.Fatal(err)
	}
}
