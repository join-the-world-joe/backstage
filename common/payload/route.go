package payload

import (
	"backstage/common/macro/service"
	"backstage/common/major"
)

func GetUpstreamServiceName(packet *PacketInternal) string {
	switch packet.GetRequest().GetHeader().GetMajor() {
	case major.Generic:
		return service.Generic
	case major.FrontendGateway:
		return service.FrontendGateway
	case major.BackendGateway:
		return service.BackendGateway
	case major.Account:
		return service.Account
	case major.Admin:
		return service.Admin
	default:
		return "Unknown"
	}
}
