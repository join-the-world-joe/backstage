package product

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func FetchIdListOfProduct(ctx context.Context, req *FetchIdListOfProductReq, rsp *FetchIdListOfProductRsp) error {
	srv, err := global.SelectService(service.Product)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Product, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchIdListOfProduct", req, rsp)
}

func FetchRecordsOfProduct(ctx context.Context, req *FetchRecordsOfProductReq, rsp *FetchRecordsOfProductRsp) error {
	srv, err := global.SelectService(service.Product)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.Product, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchRecordsOfProduct", req, rsp)
}
