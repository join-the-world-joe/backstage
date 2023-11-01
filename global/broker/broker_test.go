package broker

import (
	"backstage/abstract/selector"
	"backstage/common/broker"
	"backstage/common/conf"
	"backstage/global/log"
	"backstage/plugin/logger/zap"
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
	"time"
)

var server_conf = `
[Broker.Notification1]
	User = "root"
	Password = "123456"
	Category = "NATS Server"
	Servers = ["192.168.130.128:14111"]

[Broker.Notification2]
	User = "root"
	Password = "123456"
	Category = "NATS Server"
	Servers = ["192.168.130.128:14222"]

[Broker.Notification3]
	User = "root"
	Password = "123456"
	Category = "NATS Server"
	Servers = ["192.168.130.128:14333"]
`

var cluster_conf = `
[Broker.Cluster]
	Servers = ["192.168.130.128:14555", "192.168.130.128:14666", "192.168.130.128:14777"]
	User = "root"
	Password = "123456"
	Category = "Clustered NATS Server"
`

var handler = func(topic string, msg []byte) {
	fmt.Println("topic: ", topic)
	fmt.Println("msg: ", string(msg))
}

func logger() {
	filePath := ""
	logFileName := "broker"
	lg, err := zap.NewLogger(
		zap.WithLevel(-1),
		zap.WithFilePath(filePath),
		zap.WithFileName(logFileName),
		zap.WithCallerSkip(2),
	)
	if err != nil {
		panic(err)
	}
	log.SetLogger(lg)
	log.Debug()
}

func TestConnection1(t *testing.T) {
	category := broker.NATS_SERVER
	servers := []string{
		"192.168.130.128:14111",
		"192.168.130.128:14222",
		"192.168.130.128:14333",
	}
	user, password := "root", "123456"
	param := map[string]interface{}{}
	for _, v := range servers {
		brk, err := connectToBroker(
			category,
			[]string{v},
			user,
			password,
			param,
		)
		if err != nil {
			t.Fatal(err)
		}
		if err = brk.Connect(); err != nil {
			t.Fatal(err)
		}
		if err = brk.Disconnect(); err != nil {
			t.Fatal(err)
		}
	}

}

func TestConnection2(t *testing.T) {
	category := broker.NATS_CLUSTER
	servers := []string{
		"192.168.130.128:14555",
		"192.168.130.128:14666",
		"192.168.130.128:14777",
	}
	user, password := "root", "123456"
	param := map[string]interface{}{}
	brk, err := connectToBroker(
		category,
		servers,
		user,
		password,
		param,
	)
	if err != nil {
		t.Fatal(err)
	}
	if err = brk.Connect(); err != nil {
		t.Fatal(err)
	}
	if err = brk.Disconnect(); err != nil {
		t.Fatal(err)
	}
}

func TestBroker1(t *testing.T) {
	nbOfBroker := 3
	topic := "any"
	brokerConfig := &conf.BrokerConf{}
	if err := toml.Unmarshal([]byte(server_conf), &brokerConfig); err != nil {
		t.Fatal(err)
	}

	logger()

	if err := Broker(brokerConfig, 0, topic, handler); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= nbOfBroker; i++ {
		which, err := Select()
		if err != nil {
			t.Fatal(err)
		}
		Debug(which)
	}
}

func TestBroker2(t *testing.T) {
	nbOfBroker := 1
	topic := "any"
	cf := &conf.BrokerConf{}
	if err := toml.Unmarshal([]byte(cluster_conf), &cf); err != nil {
		t.Fatal(err)
	}

	logger()

	if err := Broker(cf, 0, topic, handler); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= nbOfBroker; i++ {
		which, err := Select()
		if err != nil {
			t.Fatal(err)
		}
		Debug(which)
	}
}

func TestSubscribe1(t *testing.T) {
	topic1 := "any.1"
	topic2 := "any.2"
	nbOfBroker := 3
	brokerConfig := &conf.BrokerConf{}
	if err := toml.Unmarshal([]byte(server_conf), &brokerConfig); err != nil {
		t.Fatal(err)
	}

	logger()

	if err := Broker(brokerConfig, selector.RoundRobin, topic1, handler); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= nbOfBroker; i++ {
		which, err := Select()
		if err != nil {
			t.Fatal(err)
		}
		if err = Subscribe(which, topic2, handler); err != nil {
			t.Fatal(err)
		}
	}

	select {}
}

func TestPublish1(t *testing.T) {
	n := 3
	topic1 := "any.1"
	topic2 := "any.2"
	brokerConfig := &conf.BrokerConf{}
	if err := toml.Unmarshal([]byte(server_conf), &brokerConfig); err != nil {
		t.Fatal(err)
	}

	logger()

	if err := Broker(brokerConfig, selector.RoundRobin, "", nil); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= n; i++ {
		which, err := Select()
		if err != nil {
			t.Fatal(err)
		}
		if err = Publish(which, topic1, []byte("hello, world!")); err != nil {
			t.Fatal(err)
		}
		if err = Publish(which, topic2, []byte("hello, world!")); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)
	}

	select {}
}

func TestSubscribe2(t *testing.T) {
	topic1 := "any.1"
	topic2 := "any.2"
	nbOfBroker := 1
	brokerConfig := &conf.BrokerConf{}
	if err := toml.Unmarshal([]byte(cluster_conf), &brokerConfig); err != nil {
		t.Fatal(err)
	}

	logger()

	if err := Broker(brokerConfig, selector.RoundRobin, topic1, handler); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= nbOfBroker; i++ {
		which, err := Select()
		if err != nil {
			t.Fatal(err)
		}
		if err = Subscribe(which, topic2, handler); err != nil {
			t.Fatal(err)
		}
	}

	select {}
}

func TestPublish2(t *testing.T) {
	n := 1
	topic1 := "any.1"
	topic2 := "any.2"
	brokerConfig := &conf.BrokerConf{}
	if err := toml.Unmarshal([]byte(cluster_conf), &brokerConfig); err != nil {
		t.Fatal(err)
	}

	logger()

	if err := Broker(brokerConfig, selector.RoundRobin, "", nil); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= n; i++ {
		which, err := Select()
		if err != nil {
			t.Fatal(err)
		}
		if err = Publish(which, topic1, []byte("hello, world!")); err != nil {
			t.Fatal(err)
		}
		if err = Publish(which, topic2, []byte("hello, world!")); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)
	}

	select {}
}
