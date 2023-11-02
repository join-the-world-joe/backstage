package generic

import (
	"backstage/common/macro/service"
	"backstage/common/payload"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func Authenticate(ctx context.Context, req *AuthenticateReq, rsp *AuthenticateRsp) error {
	srv, err := global.SelectService(service.Generic)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Generic, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "Authenticate", req, rsp)
}

func Break(ctx context.Context, req *payload.FakeReq, rsp *payload.FakeRsp) error {
	srv, err := global.SelectService(service.Generic)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Generic, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "Break", req, rsp)
}
