package rpc

import (
	"backstage/common/payload"
	"backstage/common/service/backend"
	"backstage/global"
	"backstage/service/backend/business"
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

func (p *Async) SignIn(ctx context.Context, req *backend.SignInReq, rsp *backend.SignInRsp) error {
	return business.SignIn(ctx, req, rsp)
}

func (p *Async) FetchMenuList(ctx context.Context, req *backend.FetchMenuListReq, rsp *backend.FetchMenuListRsp) error {
	return business.FetchMenuList(ctx, req, rsp)
}
