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

func fetchRecordsOfGood(packet *payload.PacketInternal) {
	req := &admin.FetchRecordsOfGoodReq{}
	rsp := &admin.FetchRecordsOfGoodRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchRecordsOfGood.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.FetchRecordsOfGood(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchRecordsOfGood.business.FetchRecordsOfGood failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchRecordsOfGood.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.FetchRecordsOfGoodRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchRecordsOfGood.route.Downstream failure, err: ", err.Error())
		return
	}
}
