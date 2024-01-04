package oss

import (
	"backstage/common/macro/service"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func FetchHeaderListOfObjectFileList(ctx context.Context, req *FetchHeaderListOfObjectFileListReq, rsp *FetchHeaderListOfObjectFileListRsp) error {
	srv, err := global.SelectService(service.OSS)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.OSS, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "FetchHeaderListOfObjectFileList", req, rsp)
}

func VerifyObjectFileList(ctx context.Context, req *VerifyObjectFileListReq, rsp *VerifyObjectFileListRsp) error {
	srv, err := global.SelectService(service.OSS)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.OSS, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "VerifyObjectFileList", req, rsp)
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
