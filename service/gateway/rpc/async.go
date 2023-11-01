package rpc

import (
	"backstage/common/code"
	"backstage/common/service/gateway"
	"backstage/global/routing"
	"backstage/service/gateway/runtime"
	"context"
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
