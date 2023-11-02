package gateway

import "backstage/abstract/registry"

type RouteReq struct {
	ServiceName string
	Service     *registry.Service
}

type RouteRsp struct {
}
