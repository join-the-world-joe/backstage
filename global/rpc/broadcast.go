package rpc

import (
	"backstage/abstract/registry"
	"backstage/common/code"
	service2 "backstage/common/macro/config"
	"backstage/common/macro/service"
	"backstage/common/service/gateway"
	"backstage/global"
	"backstage/global/log"
	"context"
	"github.com/spf13/cast"
)

func Broadcast(ctx context.Context, req *gateway.BroadcastReq, rsp *gateway.BroadcastRsp) error {
	srvList, err := global.Registry().ListServices(
		&registry.Service{
			Group: service2.ServiceGroup,
			Name:  service.Gateway,
		},
	)
	if err != nil {
		return err
	}
	if len(srvList) <= 0 {
		rsp.Code = code.EntryNotFound
		return nil
	}
	for _, v := range srvList {
		xClient, err := GetXClient(service.Gateway, v.Id, v.Ip, cast.ToString(v.Port))
		if err != nil {
			log.Error("Broadcast.GetXClient failure, err: ", err.Error())
			continue
		}
		err = xClient.Call(ctx, "Broadcast", req, rsp)
		if err != nil {
			log.Error("Broadcast.xClient.Call failure, err: ", err.Error())
			continue
		}
	}
	return nil
}
