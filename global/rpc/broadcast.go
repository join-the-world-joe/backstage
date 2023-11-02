package rpc

import (
	"backstage/abstract/registry"
	"backstage/common/code"
	service2 "backstage/common/macro/config"
	"backstage/common/macro/service"
	"backstage/common/protocol/gateway"
	"backstage/global"
	"backstage/global/log"
	"context"
	"github.com/spf13/cast"
)

func BroadcastBackend(ctx context.Context, req *gateway.BroadcastReq, rsp *gateway.BroadcastRsp) error {
	srvList, err := global.Registry().ListServices(
		&registry.Service{
			Group: service2.ServiceGroup,
			Name:  service.BackendGateway,
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
		xClient, err := GetXClient(service.BackendGateway, v.Id, v.Ip, cast.ToString(v.Port))
		if err != nil {
			log.Error("BroadcastBackend.GetXClient failure, err: ", err.Error())
			continue
		}
		err = xClient.Call(ctx, "Broadcast", req, rsp)
		if err != nil {
			log.Error("BroadcastBackend.xClient.Call failure, err: ", err.Error())
			continue
		}
	}
	return nil
}

func BroadcastFrontend(ctx context.Context, req *gateway.BroadcastReq, rsp *gateway.BroadcastRsp) error {
	srvList, err := global.Registry().ListServices(
		&registry.Service{
			Group: service2.ServiceGroup,
			Name:  service.FrontendGateway,
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
		xClient, err := GetXClient(service.FrontendGateway, v.Id, v.Ip, cast.ToString(v.Port))
		if err != nil {
			log.Error("FrontendGateway.GetXClient failure, err: ", err.Error())
			continue
		}
		err = xClient.Call(ctx, "Broadcast", req, rsp)
		if err != nil {
			log.Error("FrontendGateway.xClient.Call failure, err: ", err.Error())
			continue
		}
	}
	return nil
}
