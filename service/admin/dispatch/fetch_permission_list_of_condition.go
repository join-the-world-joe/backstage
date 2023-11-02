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

func fetchPermissionListOfCondition(packet *payload.PacketInternal) {
	req := &admin.FetchPermissionListOfConditionReq{}
	rsp := &admin.FetchPermissionListOfConditionRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchPermissionListOfCondition.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Id = packet.GetSession().GetUserId()

	err = business.FetchPermissionListOfCondition(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchPermissionListOfCondition.business.FetchPermissionListOfCondition failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchPermissionListOfCondition.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.FetchPermissionListOfConditionRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchPermissionListOfCondition.route.Downstream failure, err: ", err.Error())
		return
	}
}
