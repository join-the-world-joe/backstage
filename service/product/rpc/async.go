package rpc

import (
	"backstage/common/payload"
	"backstage/common/protocol/product"
	"backstage/global"
	"backstage/service/product/business"
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

func (p *Async) FetchIdListOfProduct(ctx context.Context, req *product.FetchIdListOfProductReq, rsp *product.FetchIdListOfProductRsp) error {
	return business.FetchIdListOfProduct(ctx, req, rsp)
}

func (p *Async) FetchRecordsOfProduct(ctx context.Context, req *product.FetchRecordsOfProductReq, rsp *product.FetchRecordsOfProductRsp) error {
	return business.FetchRecordsOfProduct(ctx, req, rsp)
}
