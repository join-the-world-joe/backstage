package global

import "backstage/common/payload"

var _forward *payload.PacketInternalChannel

func SetForward(forward *payload.PacketInternalChannel) {
	_forward = forward
}

func Forward() *payload.PacketInternalChannel {
	return _forward
}
