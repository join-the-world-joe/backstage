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

func removeOutdatedRecordsOfADOfSnacks(packet *payload.PacketInternal) {
	req := &admin.RemoveOutdatedRecordsOfADOfSnacksReq{}
	rsp := &admin.RemoveOutdatedRecordsOfADOfSnacksRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.RemoveOutdatedRecordsOfADOfSnacks(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.RemoveOutdatedRecordsOfADOfSnacks failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.RemoveOutdatedRecordsOfADOfSnacksRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("route.Downstream failure, err: ", err.Error())
		return
	}
}
