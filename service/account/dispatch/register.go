package dispatch

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/account"
	"backstage/common/route"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/service/account/business"

	"context"
	"encoding/json"
)

func register(packet *payload.PacketInternal) {
	req := &account.RegisterReq{}
	rsp := &account.RegisterRsp{}

	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
	if err != nil {
		log.Error("json.Unmarshal failure, err: ", err.Error())
		return
	}

	err = business.Register(context.Background(), req, rsp)
	if err != nil {
		log.Error("business.Register failure, err: ", err.Error())
		return
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Error("json.Marshal failure, err: ", err.Error())
		return
	}

	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Account,
			Minor: account.RegisterRsp_,
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
