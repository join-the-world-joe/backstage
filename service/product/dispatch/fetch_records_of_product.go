package dispatch

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/product"
	"backstage/common/route"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/product/business"
	"context"
	"encoding/json"
)

func fetchRecordsOfProduct(packet *payload.PacketInternal) {
	req := &product.FetchRecordsOfProductReq{}
	rsp := &product.FetchRecordsOfProductRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.FetchRecordsOfProduct(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.FetchRecordsOfGood failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Product,
			Minor: product.FetchRecordsOfProductRsp_,
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
