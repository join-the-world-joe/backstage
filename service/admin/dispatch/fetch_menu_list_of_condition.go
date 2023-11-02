package dispatch

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/admin"
	"backstage/common/route"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/admin/business"
	"context"
	"encoding/json"
)

func fetchMenuListOfCondition(packet *payload.PacketInternal) {
	req := &admin.FetchMenuListOfConditionReq{}
	rsp := &admin.FetchMenuListOfConditionRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchMenuListOfCondition.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Id = packet.GetSession().GetUserId()

	err = business.FetchMenuListOfCondition(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchMenuListOfCondition.business.FetchMenuListOfRole failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchMenuListOfCondition.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.FetchMenuListOfConditionRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchMenuListOfRoleList.route.Downstream failure, err: ", err.Error())
		return
	}
}
