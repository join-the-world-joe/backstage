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

func fetchVersionOfADOfDealsOfToday(packet *payload.PacketInternal) {
	req := &advertisement.FetchVersionOfADOfDealsOfTodayReq{}
	rsp := &advertisement.FetchVersionOfADOfDealsOfTodayRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfDealsOfToday.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.FetchVersionOfADOfDealsOfToday(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfDealsOfToday.business.FetchVersionOfADOfDealsOfToday failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfDealsOfToday.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Advertisement,
			Minor: advertisement.FetchVersionOfADOfDealsOfTodayRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchCarouselAdvertisement.route.Downstream failure, err: ", err.Error())
		return
	}
}
