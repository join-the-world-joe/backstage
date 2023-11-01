package business

import (
	"backstage/common/code"
	"backstage/common/service/backend"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func FetchMenuListOfRole(ctx context.Context, req *backend.FetchMenuListOfRoleReq, rsp *backend.FetchMenuListOfRoleRsp) error {
	if len(req.Role) <= 0 || len(req.Query) <= 0 {
		rsp.Code = code.InvalidDataType
		return nil
	}
	hasPermission := rbac.HasPermission(req.Role, cast.ToInt(backend.FetchMenuListOfRoleReq_))
	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}
	rsp.Body = rbac.GetMenuListOfRole(req.Query)
	rsp.Code = code.Success
	return nil
}
