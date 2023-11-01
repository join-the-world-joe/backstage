package cluster

import (
	"context"
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
`

func TestConnection1(t *testing.T) {
	servers := []string{"192.168.130.128:7001", "192.168.130.128:7002", "192.168.130.128:7003"}
	user := ""
	password := "123456"

	cluster, err := connectToClusteredRedis(servers, user, password)
	if err != nil {
		t.Fatal(err)
	}

	name, err := cluster.Get(context.Background(), "name").Result()
	if err != nil {
		t.Log("err： ", err.Error())
	} else {
		t.Log("name: ", name)
	}

	name, err = cluster.Set(context.Background(), "name", uuid.New().String(), 0).Result()
	if err != nil {
		t.Log("err： ", err.Error())
	} else {
		t.Log("name: ", name)
	}

	if err = cluster.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestConnection2(t *testing.T) {
	which := "test"
	cacheConf := &conf.CacheConf{}

	if err := toml.Unmarshal([]byte(cluster_conf), &cacheConf); err != nil {
		t.Fatal(err)
	}

	cluster, err := GetClient(cacheConf, which)
	if err != nil {
		t.Fatal(err)
	}

	name, err := cluster.Get(context.Background(), "name").Result()
	if err != nil {
		t.Log("err： ", err.Error())
	} else {
		t.Log("name: ", name)
	}

	name, err = cluster.Set(context.Background(), "name", uuid.New().String(), 0).Result()
	if err != nil {
		t.Log("err： ", err.Error())
	} else {
		t.Log("name: ", name)
	}

	if err = cluster.Close(); err != nil {
		t.Fatal(err)
	}
}
