package sms

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func SendVerificationCode(ctx context.Context, req *SendVerificationCodeReq, rsp *SendVerificationCodeRsp) error {
	srv, err := global.SelectService(service.SMS)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.SMS, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "SendVerificationCode", req, rsp)
}
