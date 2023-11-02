package business

import (
	"backstage/common/protocol/gateway"
	"context"
)

func Echo(ctx context.Context, req *gateway.PingReq, rsp *gateway.PongRsp) error {
	rsp.Message = req.Message
	return nil
}
