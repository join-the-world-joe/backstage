package dispatch

import (
	"backstage/common/payload"
	"backstage/common/route"
	"backstage/common/service/backend"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/backend/business"
	"context"
)

func fetchPermissionList(packet *payload.PacketInternal) {
	req := &backend.FetchPermissionListReq{Role: packet.Session.Role}
	rsp := &backend.FetchPermissionListRsp{}

	err := business.FetchPermissionList(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchPermissionList.business.FetchMenuList failure, err: ", err.Error())
		return
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchPermissionList.route.Downstream failure, err: ", err.Error())
		return
	}
}
