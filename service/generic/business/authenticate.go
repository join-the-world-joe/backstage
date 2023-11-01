package business

import (
	"backstage/common/service/generic"
	"backstage/global/log"
	"context"
)

var _userId = int64(0)

func Authenticate(ctx context.Context, req *generic.AuthenticateReq, rsp *generic.AuthenticateRsp) error {
	log.Debug("Authenticate has been called!")
	_userId++
	rsp.UserId = _userId
	rsp.Code = 0
	return nil
}
