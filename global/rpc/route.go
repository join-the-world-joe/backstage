package rpc

import (
	"backstage/common/service/gateway"
	"context"
)

func Route(ctx context.Context, srvName, srvId, host, port string, req *gateway.RouteReq, rsp *gateway.RouteRsp) error {
	xClient, err := GetXClient(srvName, srvId, host, port)
	if err != nil {
		return err
	}
	return xClient.Call(ctx, "Route", req, rsp)
}
