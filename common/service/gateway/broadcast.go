package gateway

import "backstage/common/payload"

type BroadcastReq struct {
	Packet *payload.PacketClient
}

type BroadcastRsp struct {
	Code int
}
