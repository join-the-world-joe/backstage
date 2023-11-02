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

func softDeleteUserRecord(packet *payload.PacketInternal) {
	req := &admin.SoftDeleteUserRecordReq{}
	rsp := &admin.SoftDeleteUserRecordRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.softDeleteUserRecord.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.Id = packet.GetSession().GetUserId()

	err = business.SoftDeleteUserRecord(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.softDeleteUserRecord.business.SoftDeleteUserRecord failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.softDeleteUserRecord.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.SoftDeleteUserRecordRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.softDeleteUserRecord.route.Downstream failure, err: ", err.Error())
		return
	}
}
