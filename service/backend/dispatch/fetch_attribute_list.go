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

func fetchAttributeList(packet *payload.PacketInternal) {
	req := &backend.FetchAttributeListReq{Role: packet.Session.Role}
	rsp := &backend.FetchAttributeListRsp{}

	err := business.FetchAttributeList(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchAttributeList.business.FetchAttributeList failure, err: ", err.Error())
		return
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchAttributeList.route.Downstream failure, err: ", err.Error())
		return
	}
}
