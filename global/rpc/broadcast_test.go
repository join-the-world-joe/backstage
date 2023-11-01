package rpc

import (
	"backstage/common/payload"
	"backstage/common/service/gateway"
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestBroadcast(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	packet := &payload.PacketClient{
		Header: &payload.Header{
			Major: "major",
			Minor: "minor",
		},
		Body: []byte("{\"message\":\"停服公告\"}"),
	}
	req := &gateway.BroadcastReq{Packet: packet}
	rsp := &gateway.BroadcastRsp{}

	err := Broadcast(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Code: ", rsp.Code)
}
