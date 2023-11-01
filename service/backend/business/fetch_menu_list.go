package business

import (
	"backstage/common/code"
	"backstage/common/service/backend"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func FetchMenuList(ctx context.Context, req *backend.FetchMenuListReq, rsp *backend.FetchMenuListRsp) error {
	if len(req.Role) <= 0 {
		rsp.Code = code.InvalidDataType
		return nil
	}
	hasPermission := rbac.HasPermission(req.Role, cast.ToInt(backend.FetchMenuListReq_))
	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}
	rsp.Body = rbac.GetMenuList()
	rsp.Code = code.Success
	return nil
}
