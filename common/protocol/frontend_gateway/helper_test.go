package frontend_gateway

import (
	"backstage/common/protocol/gateway"
	"backstage/common/protocol/inform"
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestForceOffline(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &gateway.ForceOfflineReq{
		UserId: 1,
		Notification: &inform.Notification{
			Event:   inform.EventForceOffline,
			Message: inform.MessageForceOffline,
		},
	}
	rsp := &gateway.ForceOfflineRsp{}
	err := ForceOffline(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	//time.Sleep(time.Second)
}
