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

func updateRecordOfAdvertisement(packet *payload.PacketInternal) {
	req := &admin.UpdateRecordOfAdvertisementReq{}
	rsp := &admin.UpdateRecordOfAdvertisementRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.updateRecordOfAdvertisement.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.UpdateRecordOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.updateRecordOfAdvertisement.business.UpdateRecordOfAdvertisement failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.updateRecordOfAdvertisement.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.UpdateRecordOfAdvertisementRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.updateRecordOfAdvertisement.route.Downstream failure, err: ", err.Error())
		return
	}
}
