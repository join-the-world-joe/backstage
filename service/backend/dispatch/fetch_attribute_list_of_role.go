package dispatch

import (
	"backstage/common/payload"
	"backstage/common/route"
	"backstage/common/service/backend"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/backend/business"
	"context"
	"encoding/json"
)

func fetchAttributeListOfRole(packet *payload.PacketInternal) {
	req := &backend.FetchAttributeListOfRoleReq{}
	rsp := &backend.FetchAttributeListOfRoleRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchAttributeListOfRole.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Role = packet.GetSession().GetRole()

	err = business.FetchAttributeListOfRole(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchAttributeListOfRole.business.FetchAttributeListOfRole failure, err: ", err.Error())
		return
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchAttributeListOfRole.route.Downstream failure, err: ", err.Error())
		return
	}
}
