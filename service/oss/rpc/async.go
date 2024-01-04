package rpc

import (
	"backstage/common/payload"
	"backstage/common/protocol/oss"
	"backstage/global"
	"backstage/service/oss/business"
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

func (p *Async) FetchHeaderListOfObjectFileListOfAdvertisement(ctx context.Context, req *oss.FetchHeaderListOfObjectFileListReq, rsp *oss.FetchHeaderListOfObjectFileListRsp) error {
	return business.FetchHeaderListOfObjectFileList(ctx, req, rsp)
}

func (p *Async) VerifyObjectFileList(ctx context.Context, req *oss.VerifyObjectFileListReq, rsp *oss.VerifyObjectFileListRsp) error {
	return business.VerifyObjectFileList(ctx, req, rsp)
}

func (p *Async) RemoveListOfObjectFile(ctx context.Context, req *oss.RemoveListOfObjectFileReq, rsp *oss.RemoveListOfObjectFileRsp) error {
	return business.RemoveListOfObjectFile(ctx, req, rsp)
}
