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

func insertRecordOfAdvertisement(packet *payload.PacketInternal) {
	req := &admin.InsertRecordOfAdvertisementReq{}
	rsp := &admin.InsertRecordOfAdvertisementRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.insertRecordOfAdvertisement.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.InsertRecordOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.insertRecordOfAdvertisement.business.InsertRecordOfAdvertisement failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.insertRecordOfAdvertisement.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.InsertRecordOfAdvertisementRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.insertRecordOfAdvertisement.route.Downstream failure, err: ", err.Error())
		return
	}
}
