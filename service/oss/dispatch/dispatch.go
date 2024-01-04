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
	case oss.FetchHeaderListOfObjectFileListReq_:
		fetchHeaderListOfObjectFileList(packet)
	case oss.VerifyObjectFileListReq_:
		verifyObjectFileList(packet)
	case oss.RemoveListOfObjectFileReq_:
		removeListOfObjectFile(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
