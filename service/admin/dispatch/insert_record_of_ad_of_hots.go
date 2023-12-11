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

func insertRecordOfADOfHots(packet *payload.PacketInternal) {
	req := &admin.InsertRecordOfADOfHotsReq{}
	rsp := &admin.InsertRecordOfADOfHotsRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("json.Unmarshal fail, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.InsertRecordOfADOfHots(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.InsertRecordOfADOfHots fail, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("json.Marshal fail, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.InsertRecordOfADOfHotsRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("route.Downstream fail, err: ", err.Error())
		return
	}
}
