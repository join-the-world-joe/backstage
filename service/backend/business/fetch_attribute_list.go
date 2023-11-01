package business

import (
	"backstage/common/code"
	"backstage/common/service/backend"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func FetchAttributeList(ctx context.Context, req *backend.FetchAttributeListReq, rsp *backend.FetchAttributeListRsp) error {
	if len(req.Role) <= 0 {
		rsp.Code = code.InvalidDataType
		return nil
	}
	hasPermission := rbac.HasPermission(req.Role, cast.ToInt(backend.FetchAttributeListReq_))
	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}
	rsp.Body = rbac.GetAttributeList()
	rsp.Code = code.Success
	return nil
}
