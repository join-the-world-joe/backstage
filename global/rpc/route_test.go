package rpc

import (
	"backstage/abstract/registry"
	"backstage/common/macro/service"
	"backstage/common/service/gateway"
	"context"
	"testing"
)

func TestRoute(t *testing.T) {
	var req *gateway.RouteReq
	choose := false
	srvName := "Gateway"
	srvId := "1"
	host := "172.20.10.6"
	port := "11001"

	if choose {
		req = &gateway.RouteReq{
			ServiceName: service.Generic,
			Service: &registry.Service{
				Id:      "1",
				Ip:      "172.20.10.6",
				Port:    11002,
				Name:    "Generic",
				Group:   "Service",
				Version: "???",
			},
		}
	} else {
		req = &gateway.RouteReq{
			ServiceName: service.Generic,
			Service:     nil,
		}
	}

	rsp := &gateway.RouteRsp{}
	err := Route(context.Background(), srvName, srvId, host, port, req, rsp)
	if err != nil {
		t.Fatal(err)
	}
}
