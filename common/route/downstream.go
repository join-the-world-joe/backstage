package route

import (
	"backstage/common/payload"
	"backstage/common/protocol"
	"backstage/global/broker"
	"backstage/global/rpc"
	"fmt"
)

func Downstream(protoc string, packet *payload.PacketInternal) error {
	switch protoc {
	case protocol.RPC:
		return rpc.P2P(packet)
	case protocol.NATS:
		return broker.P2P(packet)
	default:
		return fmt.Errorf("rpc.Downstream unknown error")
	}
}
