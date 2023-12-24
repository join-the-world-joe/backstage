package oss

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func FetchHeaderListOfObjectFileListOfAdvertisement(ctx context.Context, req *FetchHeaderListOfObjectFileListOfAdvertisementReq, rsp *FetchHeaderListOfObjectFileListOfAdvertisementRsp) error {
	srv, err := global.SelectService(service.OSS)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.OSS, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchHeaderListOfObjectFileListOfAdvertisement", req, rsp)
}

func VerifyObjectFileListOfAdvertisement(ctx context.Context, req *VerifyObjectFileListOfAdvertisementReq, rsp *VerifyObjectFileListOfAdvertisementRsp) error {
	srv, err := global.SelectService(service.OSS)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.OSS, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "VerifyObjectFileListOfAdvertisement", req, rsp)
}

func RemoveListOfObjectFile(ctx context.Context, req *RemoveListOfObjectFileReq, rsp *RemoveListOfObjectFileRsp) error {
	srv, err := global.SelectService(service.OSS)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.OSS, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "RemoveListOfObjectFile", req, rsp)
}
