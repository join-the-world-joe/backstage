package rpc

import (
	"backstage/common/code"
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/gateway"
	"backstage/common/protocol/inform"
	"backstage/global/log"
	"backstage/global/routing"
	"backstage/service/frontend_gateway/runtime"
	"context"
	"encoding/json"
	"errors"
)

type Async struct {
}

func (p *Async) P2P(ctx context.Context, req *gateway.P2PReq, rsp *gateway.P2PRsp) error {
	channel, err := runtime.LoadChannel(req.Sequence)
	if err != nil {
		return err
	}
	if err := channel.Push(req.Packet); err != nil {
		return err
	}
	return nil
}

func (p *Async) Route(ctx context.Context, req *gateway.RouteReq, rsp *gateway.RouteRsp) error {
	if len(req.ServiceName) <= 0 {
		return errors.New("len(req.ServiceName) <= 0")
	}
	routing.Store(req.ServiceName, req.Service)
	return nil
}

func (p *Async) Broadcast(ctx context.Context, req *gateway.BroadcastReq, rsp *gateway.BroadcastRsp) error {
	list := runtime.LoadOnlineSequenceList()
	if len(list) <= 0 {
		rsp.Code = code.Empty
		return nil
	}
	for _, v := range list {
		channel, err := runtime.LoadChannel(v)
		if err != nil {
			continue
		}
		channel.Push(req.Packet)
	}
	return nil
}

func (p *Async) ForceOffline(ctx context.Context, req *gateway.ForceOfflineReq, rsp *gateway.ForceOfflineRsp) error {
	session, err := runtime.LoadSession(req.UserId)
	if err != nil {
		log.ErrorF("ForceOffline.runtime.LoadSession failure, err: ", err.Error())
		return nil
	}

	bytes, err := json.Marshal(req.Notification)
	if err != nil {
		log.ErrorF("ForceOffline failure, err: ", err.Error())
		return nil
	}

	packet := &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Inform,
			Minor: inform.Notification_,
		},
		Body: bytes,
	}

	channel, err := runtime.LoadChannel(session.GetSequence())
	if err != nil {
		return err
	}
	if err := channel.Push(packet); err != nil {
		return err
	}

	session.SetForceOffline(true)
	session.SetPacketClient(packet)
	return nil
}
