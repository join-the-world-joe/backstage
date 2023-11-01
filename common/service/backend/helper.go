package backend

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func SignIn(ctx context.Context, req *SignInReq, rsp *SignInRsp) error {
	srv, err := global.SelectService(service.Backend)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Backend, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "SignIn", req, rsp)
}
