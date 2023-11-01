package rpc

import (
	"backstage/common/macro/service"
	"context"
	"testing"
)

func TestRawP2PTest(t *testing.T) {
	host := "127.0.0.1"
	port := "10010"
	method := "P2P"
	servicePath := service.Gateway
	xClient, err := connectToRPCServer(host, port, servicePath)
	if err != nil {
		t.Fatal(err)
	}

	req := &[]string{""}
	rsp := &[]string{""}
	err = xClient.Call(context.Background(), method, req, rsp)
	if err != nil {
		t.Fatal(err)
	}
}
