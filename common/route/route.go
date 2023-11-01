package route

import (
	"backstage/common/payload"
	"context"
)

type Route struct {
	Chan *payload.PacketInternalChannel
}

func (p *Route) Forward(ctx context.Context, req *payload.PacketInternal, rsp *interface{}) error {
	return p.Chan.Push(req)
}
