package dispatch

import (
	"backstage/common/payload"
	"backstage/common/service/backend"
	"backstage/global/log"
)

func Dispatch(packet *payload.PacketInternal) {
	switch packet.GetRequest().GetHeader().GetMinor() {
	case backend.FetchMenuListReq_:
		fetchMenuList(packet)
	case backend.FetchRoleListReq_:
		fetchRoleList(packet)
	case backend.FetchPermissionListReq_:
		fetchPermissionList(packet)
	case backend.FetchAttributeListReq_:
		fetchAttributeList(packet)
	case backend.FetchMenuListOfRoleReq_:
		fetchMenuListOfRole(packet)
	case backend.FetchPermissionListOfRoleReq_:
		fetchPermissionListOfRole(packet)
	case backend.FetchAttributeListOfRoleReq_:
		fetchAttributeListOfRole(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
