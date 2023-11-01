package dispatch

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/route"
	"backstage/common/service/backend"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/backend/business"
	"context"
	"encoding/json"
)

func fetchMenuListOfRole(packet *payload.PacketInternal) {
	req := &backend.FetchMenuListOfRoleReq{}
	rsp := &backend.FetchMenuListOfRoleRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchAttributeListOfRole.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Role = packet.GetSession().GetRole()

	err = business.FetchMenuListOfRole(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchMenuListOfRole.business.FetchMenuListOfRole failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchMenuListOfRole.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Backend,
			Minor: backend.FetchMenuListOfRoleRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchMenuListOfRole.route.Downstream failure, err: ", err.Error())
		return
	}
}
