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

func softDeleteRecordOfUser(packet *payload.PacketInternal) {
	req := &admin.SoftDeleteRecordOfUserReq{}
	rsp := &admin.SoftDeleteRecordOfUserRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.SoftDeleteRecordOfUser(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.SoftDeleteUserRecord failure, err: ", err.Error())
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
			Minor: admin.SoftDeleteRecordOfUserRsp_,
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
