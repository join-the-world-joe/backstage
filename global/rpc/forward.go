package rpc

import (
	"backstage/abstract/registry"
	"backstage/common/payload"
	"backstage/global"
	"backstage/global/routing"
	"context"
	"errors"
	"github.com/spf13/cast"
)

func Forward(packet *payload.PacketInternal) error {
	var err error
	var srv *registry.Service
	if srv, err = routing.Load(payload.GetUpstreamServiceName(packet)); err == nil {
		if srv == nil {
			srv, err = global.SelectService(payload.GetUpstreamServiceName(packet))
			if err != nil {
				return err
			}
		}
	} else {
		srv, err = global.SelectService(payload.GetUpstreamServiceName(packet))
		if err != nil {
			return err
		}
		if srv == nil {
			return errors.New("Forward.global.SelectService, srv == nil")
		}
	}

	xClient, err := GetXClient(payload.GetUpstreamServiceName(packet), srv.Id, srv.Ip, cast.ToString(srv.Port))
	if err != nil {
		return err
	}
	return xClient.Call(context.Background(), "Forward", packet, payload.GetFakeRsp())
}
