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

func insertRecordOfProduct(packet *payload.PacketInternal) {
	req := &admin.InsertRecordOfProductReq{}
	rsp := &admin.InsertRecordOfProductRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.insertRecordOfGood.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.InsertRecordOfProduct(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.InsertRecordOfGood failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.insertRecordOfGood.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Admin,
			Minor: admin.InsertRecordOfProductRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.insertRecordOfGood.route.Downstream failure, err: ", err.Error())
		return
	}
}
