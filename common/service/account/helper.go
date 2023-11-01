package account

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func Register(ctx context.Context, req *RegisterReq, rsp *RegisterRsp) error {
	srv, err := global.SelectService(service.Account)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Account, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "Register", req, rsp)
}

func Login(ctx context.Context, req *LoginReq, rsp *LoginRsp) error {
	srv, err := global.SelectService(service.Account)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Account, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "Login", req, rsp)
}

func Logout(ctx context.Context, req *LogoutReq, rsp *LogoutRsp) error {
	srv, err := global.SelectService(service.Account)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Account, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "Logout", req, rsp)
}
