package dispatch

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/oss"
	"backstage/common/route"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/oss/business"
	"context"
	"encoding/json"
)

func fetchHeaderListOfObjectFileListOfAdvertisement(packet *payload.PacketInternal) {
	req := &oss.FetchHeaderListOfObjectFileListOfAdvertisementReq{}
	rsp := &oss.FetchHeaderListOfObjectFileListOfAdvertisementRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.FetchHeaderListOfObjectFileListOfAdvertisement(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.FetchHeaderListOfObjectFileListOfAdvertisement failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.OSS,
			Minor: oss.FetchHeaderListOfObjectFileListOfAdvertisementRsp_,
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
