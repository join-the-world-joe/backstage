package rpc

import (
	"backstage/common/macro/service"
	"backstage/common/payload"
	"context"
	"testing"
)

func TestGetXClient(t *testing.T) {
	var packet *payload.PacketInternal

	client, err := GetXClient(service.Generic, "1", "192.168.0.11", "3335")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client)
	err = client.Call(context.Background(), "Forward", packet, payload.GetFakeRsp())
	if err != nil {
		t.Fatal(err)
	}
}

func TestConnectionRaw(t *testing.T) {
	host := "172.20.1.3"
	port := "11005"
	servicePath := "SMS"
	client, err := connectToRPCServer(host, port, servicePath)
	if err != nil {
		t.Fatal(err)
	}
	client.Close()
}
