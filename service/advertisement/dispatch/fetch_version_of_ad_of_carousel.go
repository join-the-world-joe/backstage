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

func fetchVersionOfADOfCarousel(packet *payload.PacketInternal) {
	req := &advertisement.FetchVersionOfADOfCarouselReq{}
	rsp := &advertisement.FetchVersionOfADOfCarouselRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfCarousel.json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.FetchVersionOfADOfCarousel(context.Background(), req, rsp)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfCarousel.business.FetchVersionOfCarouselAdvertisement failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfCarousel.json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Advertisement,
			Minor: advertisement.FetchVersionOfADOfCarouselRsp_,
		},
		Body: bytes,
	}

	err = route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.Error("Dispatch.fetchVersionOfADOfCarousel.route.Downstream failure, err: ", err.Error())
		return
	}
}
