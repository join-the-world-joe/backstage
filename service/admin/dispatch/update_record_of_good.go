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

func updateRecordOfGood(packet *payload.PacketInternal) {
	req := &admin.UpdateRecordOfGoodReq{}
	rsp := &admin.UpdateRecordOfGoodRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.updateRecordOfGood.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.UpdateRecordOfGood(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.updateRecordOfGood.business.UpdateRecordOfGood failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.updateRecordOfGood.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.UpdateRecordOfGoodRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.updateRecordOfGood.route.Downstream failure, err: ", err.Error())
		return
	}
}
