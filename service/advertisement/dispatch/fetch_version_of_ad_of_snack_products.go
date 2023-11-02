package dispatch

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/advertisement"
	"backstage/common/route"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/advertisement/business"
	"context"
	"encoding/json"
)

func fetchVersionOfADOfSnackProducts(packet *payload.PacketInternal) {
	req := &advertisement.FetchVersionOfADOfSnackProductsReq{}
	rsp := &advertisement.FetchVersionOfADOfSnackProductsRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfSnackProducts.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.FetchVersionOfADOfSnackProducts(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfSnackProducts.business.FetchVersionOfADOfSnackProducts failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfSnackProducts.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Advertisement,
			Minor: advertisement.FetchVersionOfADOfSnackProductsRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfSnackProducts.route.Downstream failure, err: ", err.Error())
		return
	}
}
