package template

import (
	"backstage/common/conf"
	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	"testing"
)

var redis_conf = `
[Redis.test1]
	Name = "Redis Server"
	Servers = ["192.168.130.128:16381"]
	Password = "123456"

[Redis.test2]
	Name = "Redis Server"
	Servers = ["192.168.130.128:16382"]
	Password = "123456"

[Redis.test3]
	Name = "Redis Server"
	Servers = ["192.168.130.128:16383"]
	Password = "123456"
`

func TestGetWhich(t *testing.T) {
	for i := 1; i <= Mod; i++ {
		t.Log(GetWhich(i))
	}
}

func TestCreate(t *testing.T) {
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(redis_conf), &cf); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= Mod; i++ {
		if err := Create(cf, i, uuid.New().String()); err != nil {
			t.Fatal(err)
		}
	}
}
