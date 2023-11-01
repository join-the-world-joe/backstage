package payload

import (
	"backstage/common/macro/service"
	"backstage/common/major"
)

func GetUpstreamServiceName(packet *PacketInternal) string {
	switch packet.GetRequest().GetHeader().GetMajor() {
	case major.Generic:
		return service.Generic
	case major.Gateway:
		return service.Gateway
	case major.Account:
		return service.Account
	case major.Backend:
		return service.Backend
	default:
		return "Unknown"
	}
}
