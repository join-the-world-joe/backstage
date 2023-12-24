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

func removeListOfObjectFile(packet *payload.PacketInternal) {
	req := &oss.RemoveListOfObjectFileReq{}
	rsp := &oss.RemoveListOfObjectFileRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("json.Unmarshal failure, err: ", err.Error())
		return
	}

	req.UserId = packet.GetSession().GetUserId()

	err = business.RemoveListOfObjectFile(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.RemoveListOfObjectFile failure, err: ", err.Error())
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
			Minor: oss.RemoveListOfObjectFileRsp_,
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
