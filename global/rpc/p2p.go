package rpc

import (
	"backstage/common/payload"
	"backstage/common/protocol/gateway"
	"context"
)

func P2P(packet *payload.PacketInternal) error {
	return P2P_(
		context.Background(),
		packet.GetSession().GetLoginServerName(),
		packet.GetSession().GetLoginServerId(),
		packet.GetSession().GetLoginServerHost(),
		packet.GetSession().GetLoginServerRPCPort(),
		packet.GetSession().GetSequence(),
		packet.GetResponse(),
	)
}

func P2P_(ctx context.Context, srvName, srvId, host, port string, sequence uint64, packet *payload.PacketClient) error {
	xClient, err := GetXClient(srvName, srvId, host, port)
	if err != nil {
		return err
	}
	req := &gateway.P2PReq{
		Sequence: sequence,
		Packet:   packet,
	}
	return xClient.Call(ctx, "P2P", req, &gateway.P2PRsp{})
}
