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

func fetchFieldListOfCondition(packet *payload.PacketInternal) {
	req := &admin.FetchFieldListOfConditionReq{}
	rsp := &admin.FetchFieldListOfConditionRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchFieldListOfCondition.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Id = packet.GetSession().GetUserId()

	err = business.FetchFieldListOfCondition(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchFieldListOfCondition.business.FetchFieldListOfCondition failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchFieldListOfCondition.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.FetchFieldListOfConditionRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchFieldListOfCondition.route.Downstream failure, err: ", err.Error())
		return
	}
}
