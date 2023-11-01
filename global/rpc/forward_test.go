package rpc

import (
	"backstage/common/macro/service"
	"context"
	"testing"
)

func TestRawForward(t *testing.T) {
	host := "127.0.0.1"
	port := "11005"
	method := "Forward"
	servicePath := service.SMS
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

func TestForward(t *testing.T) {
	err := Forward(nil)
	if err != nil {
		t.Fatal(err)
	}
}
