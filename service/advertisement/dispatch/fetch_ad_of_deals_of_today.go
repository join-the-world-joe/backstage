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

func fetchADOfDealsOfToday(packet *payload.PacketInternal) {
	req := &advertisement.FetchADOfDealsOfTodayReq{}
	rsp := &advertisement.FetchADOfDealsOfTodayRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchADOfDealsOfToday.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.FetchADOfDealsOfToday(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchADOfDealsOfToday.business.FetchADOfCarousel failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchADOfDealsOfToday.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Advertisement,
			Minor: advertisement.FetchADOfDealsOfTodayRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchADOfDealsOfToday.route.Downstream failure, err: ", err.Error())
		return
	}
}
