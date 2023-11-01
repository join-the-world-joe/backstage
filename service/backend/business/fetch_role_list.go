package business

import (
	"backstage/common/code"
	"backstage/common/service/backend"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func FetchRoleList(ctx context.Context, req *backend.FetchRoleListReq, rsp *backend.FetchRoleListRsp) error {
	if len(req.Role) <= 0 {
		rsp.Code = code.InvalidDataType
		return nil
	}
	hasPermission := rbac.HasPermission(req.Role, cast.ToInt(backend.FetchRoleListReq_))
	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}
	rsp.RoleList = rbac.GetRoleList()
	rsp.Code = code.Success
	return nil
}
