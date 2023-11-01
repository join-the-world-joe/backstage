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

func fetchMenuList(packet *payload.PacketInternal) {
	req := &backend.FetchMenuListReq{Role: packet.Session.Role}
	rsp := &backend.FetchMenuListRsp{}

	err := business.FetchMenuList(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchMenuList.business.FetchMenuList failure, err: ", err.Error())
		return
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchMenuList.route.Downstream failure, err: ", err.Error())
		return
	}
}
