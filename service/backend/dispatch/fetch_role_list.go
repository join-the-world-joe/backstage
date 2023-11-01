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

func fetchRoleList(packet *payload.PacketInternal) {
	req := &backend.FetchRoleListReq{Role: packet.Session.Role}
	rsp := &backend.FetchRoleListRsp{}

	err := business.FetchRoleList(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchRoleList.business.FetchMenuList failure, err: ", err.Error())
		return
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchRoleList.route.Downstream failure, err: ", err.Error())
		return
	}
}
