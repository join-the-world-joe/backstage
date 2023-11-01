package dispatch

import (
	"backstage/common/payload"
	"backstage/common/service/generic"
	"backstage/global/log"
	"backstage/service/generic/business"
)

func Dispatch(packet *payload.PacketInternal) {
	switch packet.GetRequest().GetHeader().GetMinor() {
	case generic.Ping:
		business.Pong(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
