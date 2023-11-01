package rpc

import (
	"backstage/common/payload"
	"backstage/common/service/generic"
	"backstage/global"
	"backstage/service/generic/business"
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

func (p *Async) Authenticate(ctx context.Context, req *generic.AuthenticateReq, rsp *generic.AuthenticateRsp) error {
	return business.Authenticate(ctx, req, rsp)
}
