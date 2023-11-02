package server

import (
	"backstage/common/macro/service"
	"context"
	"fmt"
	"testing"
)

type _server struct {
	*Async
}

type Async struct {
}

func (p *Async) P2P(ctx context.Context, req *interface{}, rsp *interface{}) error {
	fmt.Println("P2P has benn called")
	return nil
}

func (p *Async) Forward(ctx context.Context, req *interface{}, rsp *interface{}) error {
	fmt.Println("Forward has benn called")
	return nil
}

func TestNewServer(t *testing.T) {
	host := "172.20.10.6"
	port := "10010"
	if err := NewServer(
		WithHost(host),
		WithPort(port),
		WithServer(&_server{}),
		WithServicePath(service.Gateway),
	); err != nil {
		t.Fatal(err)
	}
	select {}
}
