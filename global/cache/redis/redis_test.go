package redis

import (
	"backstage/common/conf"
	"github.com/BurntSushi/toml"
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

func TestConnection1(t *testing.T) {
	db := int64(0)
	server := "192.168.130.128:16381"
	user := ""
	password := "123456"
	c, err := connectToRedis(server, user, password, db)
	if err != nil {
		t.Fatal(err)
	}
	if err = c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestConnection2(t *testing.T) {
	db := int64(0)
	which := "test1"
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(redis_conf), &cf); err != nil {
		t.Fatal(err)
	}

	client, err := GetClient(cf, which, db)
	if err != nil {
		t.Fatal(err)
	}

	if err = client.Close(); err != nil {
		t.Fatal(err)
	}
}
