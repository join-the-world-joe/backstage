package business

import (
	"backstage/common/code"
	"backstage/common/service/backend"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func FetchPermissionList(ctx context.Context, req *backend.FetchPermissionListReq, rsp *backend.FetchPermissionListRsp) error {
	if len(req.Role) <= 0 {
		rsp.Code = code.InvalidDataType
		return nil
	}
	hasPermission := rbac.HasPermission(req.Role, cast.ToInt(backend.FetchPermissionListReq_))
	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}
	rsp.PermissionList = rbac.GetPermissionList()
	rsp.Code = code.Success
	return nil
}
