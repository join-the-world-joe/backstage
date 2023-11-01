package rpc

import (
	"backstage/common/payload"
	"backstage/common/service/sms"
	"backstage/global"
	"backstage/service/sms/business"
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

func (p *Async) SendVerificationCode(ctx context.Context, req *sms.SendVerificationCodeReq, rsp *sms.SendVerificationCodeRsp) error {
	return business.SendVerificationCode(ctx, req, rsp)
}
