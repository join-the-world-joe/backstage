package dispatch

import (
	"backstage/common/payload"
	"backstage/common/protocol/account"
	"backstage/global/log"
)

func Dispatch(packet *payload.PacketInternal) {
	switch packet.GetRequest().GetHeader().GetMinor() {
	case account.RegisterReq_:
		register(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
