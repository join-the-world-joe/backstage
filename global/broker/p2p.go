package broker

import (
	"backstage/common/broker"
	"backstage/common/payload"
	"backstage/common/protocol/gateway"
	"encoding/json"
	"fmt"
)

func P2P(packet *payload.PacketInternal) error {
	which, err := Select()
	if err != nil {
		return err
	}
	req := &gateway.P2PReq{
		Sequence: packet.GetSession().GetSequence(),
		Packet:   packet.GetResponse(),
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return Publish(which,
		fmt.Sprintf(broker.P2P, packet.GetSession().GetLoginServerName(), packet.GetSession().GetLoginServerId()),
		bytes,
	)
}
