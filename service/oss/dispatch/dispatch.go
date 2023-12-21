package dispatch

import (
	"backstage/common/payload"
	"backstage/global/log"
)

func Dispatch(packet *payload.PacketInternal) {
	minor := packet.GetRequest().GetHeader().GetMinor()
	log.Info("Dispatch.Minor: ", minor)
	switch minor {
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
