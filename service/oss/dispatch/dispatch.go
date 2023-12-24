package dispatch

import (
	"backstage/common/payload"
	"backstage/common/protocol/oss"
	"backstage/global/log"
)

func Dispatch(packet *payload.PacketInternal) {
	minor := packet.GetRequest().GetHeader().GetMinor()
	log.Info("Dispatch.Minor: ", minor)
	switch minor {
	case oss.FetchHeaderListOfObjectFileListOfAdvertisementReq_:
		fetchHeaderListOfObjectFileListOfAdvertisement(packet)
	case oss.VerifyObjectFileListOfAdvertisementReq_:
		verifyObjectFileListOfAdvertisement(packet)
	case oss.RemoveListOfObjectFileReq_:
		removeListOfObjectFile(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
