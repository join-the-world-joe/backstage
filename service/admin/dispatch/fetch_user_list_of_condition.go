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

func fetchUserListOfCondition(packet *payload.PacketInternal) {
	req := &admin.FetchUserListOfConditionReq{}
	rsp := &admin.FetchUserListOfConditionRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchUserListOfCondition.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Id = packet.GetSession().GetUserId()

	err = business.FetchUserListOfCondition(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchUserListOfCondition.business.FetchUserListOfCondition failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchUserListOfCondition.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.FetchUserListOfConditionRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchUserListOfCondition.route.Downstream failure, err: ", err.Error())
		return
	}
}
