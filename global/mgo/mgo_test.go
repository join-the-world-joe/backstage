package mgo

import (
	"context"
	"github.com/BurntSushi/toml"
	"go-micro-framework/common/conf"
	"testing"
)

var mgo_server = `
[Mongo.Singapore1]
	Servers = ["mongodb://192.168.130.128:37017"]
	User = "root"
	Password = "123456"
[Mongo.Singapore2]
	Servers = ["mongodb://192.168.130.128:37018"]
	User = "root"
	Password = "123456"
[Mongo.Singapore3]
	Servers = ["mongodb://192.168.130.128:37019"]
	User = "root"
	Password = "123456"
`

var mgo_replica = `
[Mongo.Replica]
	Servers = ["mongodb://192.168.130.128:27021", "mongodb://192.168.130.128:27022", "mongodb://192.168.130.128:27023"]
	User = "root"
	Password = "123456"
`

func TestRawConnection1(t *testing.T) {
	servers := []string{"mongodb://192.168.130.128:37017"}
	user := "root"
	password := "123456"
	client, err := connectToMongo(servers, user, password)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Disconnect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestConnection1(t *testing.T) {
	which := "Replica"
	cf := &conf.MongoConf{}
	if err := toml.Unmarshal([]byte(mgo_replica), &cf); err != nil {
		t.Fatal(err)
	}

	client, err := GetClient(cf, which)
	if err != nil {
		t.Fatal(err)
	}

	err = client.Disconnect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
