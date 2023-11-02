package business

import (
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/generic"
	"backstage/common/route"
	"backstage/global"
	"backstage/global/config"
	"backstage/global/log"
	"fmt"
)

func Pong(packet *payload.PacketInternal) {
	packet.Response = &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Generic,
			Minor: generic.Pong,
		},
		Body: []byte(fmt.Sprintf("Pong from %v.%v", global.ServiceName(), global.ServiceId())),
	}
	err := route.Downstream(
		config.DownstreamProtocol(),
		packet,
	)
	if err != nil {
		log.ErrorF(err.Error())
	}
}
