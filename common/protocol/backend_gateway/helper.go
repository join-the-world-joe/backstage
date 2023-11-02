package backend_gateway

import (
	"backstage/common/macro/service"
	"backstage/common/protocol/gateway"
	"backstage/global"
	"backstage/global/rpc"
	"context"
	"github.com/spf13/cast"
)

func ForceOffline(ctx context.Context, req *gateway.ForceOfflineReq, rsp *gateway.ForceOfflineRsp) error {
	srv, err := global.SelectService(service.BackendGateway)
	if err != nil {
		return err
	}
	xClient, err := rpc.GetXClient(service.BackendGateway, srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "ForceOffline", req, rsp)
}
