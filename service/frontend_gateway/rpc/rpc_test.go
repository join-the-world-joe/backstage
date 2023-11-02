package rpc

import (
	"backstage/common/macro/service"
	"backstage/common/payload"
	"backstage/global/rpc"
	"testing"
)

func TestP2P(t *testing.T) {
	sequence := uint64(1)
	downstreamServerId := "1"
	downstreamServerName := service.FrontendGateway
	downstreamServerHost := "172.20.10.6"
	downstreamServerPort := "11001"
	major := "test"
	minor := "test"
	body := "hello, world"
	packet := &payload.PacketInternal{
		Session: &payload.Session{
			Sequence:           sequence,
			LoginServerId:      downstreamServerId,
			LoginServerName:    downstreamServerName,
			LoginServerHost:    downstreamServerHost,
			LoginServerRPCPort: downstreamServerPort,
		},
		Response: &payload.PacketClient{
			Header: &payload.Header{
				Major: major,
				Minor: minor,
			},
			Body: []byte(body),
		},
	}

	err := rpc.P2P(packet)
	if err != nil {
		t.Fatal(err)
	}
}

func TestForceOffline(t *testing.T) {

}
