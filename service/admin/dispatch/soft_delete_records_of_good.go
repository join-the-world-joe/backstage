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

func softDeleteRecordsOfGood(packet *payload.PacketInternal) {
	req := &admin.SoftDeleteRecordsOfGoodReq{}
	rsp := &admin.SoftDeleteRecordsOfGoodRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.softDeleteRecordsOfGood.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.SoftDeleteRecordsOfGood(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.softDeleteRecordsOfGood.business.SoftDeleteRecordsOfGood failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.softDeleteRecordsOfGood.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.SoftDeleteRecordsOfGoodRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.softDeleteRecordsOfGood.route.Downstream failure, err: ", err.Error())
		return
	}
}
