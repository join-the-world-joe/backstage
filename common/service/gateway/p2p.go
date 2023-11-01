package gateway

import (
	"backstage/common/payload"
)

type P2PReq struct {
	Sequence uint64
	Packet   *payload.PacketClient
}

type P2PRsp struct {
}
