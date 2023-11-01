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

func fetchPermissionListOfRole(packet *payload.PacketInternal) {
	req := &backend.FetchPermissionListOfRoleReq{}
	rsp := &backend.FetchPermissionListOfRoleRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchAttributeListOfRole.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Role = packet.GetSession().GetRole()

	err = business.FetchPermissionListOfRole(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchPermissionListOfRole.business.FetchPermissionListOfRole failure, err: ", err.Error())
		return
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchPermissionListOfRole.route.Downstream failure, err: ", err.Error())
		return
	}
}
