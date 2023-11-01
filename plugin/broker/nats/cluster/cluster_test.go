package cluster

import (
	"fmt"
	"testing"
)

var servers = []string{"192.168.130.128:14444", "192.168.130.128:14333", "192.168.130.128:14222"}
var user, password = "root", "123456"
var handler = func(topic string, msg []byte) {
	fmt.Println("topic: ", topic, ", msg: ", msg)
}

func TestConnection1(t *testing.T) {
	brk, err := NewBroker(
		WithServers(servers),
		WithAuth(user, password),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err = brk.Connect(); err != nil {
		t.Fatal(err)
	}
	if err = brk.Disconnect(); err != nil {
		t.Fatal()
	}
}

func TestSubscribe(t *testing.T) {
	topic := "any"
	brk, err := NewBroker(
		WithServers(servers),
		WithAuth(user, password),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err = brk.Connect(); err != nil {
		t.Fatal(err)
	}
	if err = brk.Subscribe(topic, handler); err != nil {
		t.Fatal(err)
	}

	select {}
}

func TestPublish(t *testing.T) {
	topic := "any"
	brk, err := NewBroker(
		WithServers(servers),
		WithAuth(user, password),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err = brk.Connect(); err != nil {
		t.Fatal(err)
	}
	if err = brk.Publish(topic, []byte("hello")); err != nil {
		t.Fatal(err)
	}
	if err = brk.Disconnect(); err != nil {
		t.Fatal(err)
	}
}

func TestUnSubscribe(t *testing.T) {
	topic1 := "topic1"
	topic2 := "topic2"
	brk, err := NewBroker(
		WithServers(servers),
		WithAuth(user, password),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err = brk.Connect(); err != nil {
		t.Fatal(err)
	}
	if err = brk.Subscribe(topic1, handler); err != nil {
		t.Fatal(err)
	}
	if err = brk.Subscribe(topic2, handler); err != nil {
		t.Fatal(err)
	}
	brk.Debug()
	if err = brk.UnSubscribe(topic2); err != nil {
		t.Fatal(err)
	}
	brk.Debug()
}
