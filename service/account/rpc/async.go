package rpc

import (
	"backstage/common/payload"
	"backstage/common/protocol/account"
	"backstage/global"
	"backstage/service/account/business"
	"context"
)

type Async struct {
}

func (p *Async) Forward(ctx context.Context, packet *payload.PacketInternal, rsp *interface{}) error {
	if err := global.Forward().Push(packet); err != nil {
		return err
	}
	return nil
}

func (p *Async) Register(ctx context.Context, req *account.RegisterReq, rsp *account.RegisterRsp) error {
	return business.Register(ctx, req, rsp)
}

func (p *Async) Login(ctx context.Context, req *account.LoginReq, rsp *account.LoginRsp) error {
	return business.Login(ctx, req, rsp)
}
