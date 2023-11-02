package route

import (
	"backstage/common/payload"
	protocol "backstage/common/server"
	"backstage/global/broker"
	"backstage/global/rpc"
	"fmt"
)

func Upstream(protoc string, packet *payload.PacketInternal) error {
	switch protoc {
	case protocol.RPC:
		return rpc.Forward(packet)
	case protocol.NATS:
		return broker.Forward(packet)
	default:
		return fmt.Errorf("route.Upstream unknown error")
	}
}
