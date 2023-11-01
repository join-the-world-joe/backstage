package business

import (
	"backstage/common/code"
	"backstage/common/service/backend"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func FetchAttributeListOfRole(ctx context.Context, req *backend.FetchAttributeListOfRoleReq, rsp *backend.FetchAttributeListOfRoleRsp) error {
	if len(req.Role) <= 0 || len(req.Query) <= 0 {
		rsp.Code = code.InvalidDataType
		return nil
	}
	hasPermission := rbac.HasPermission(req.Role, cast.ToInt(backend.FetchAttributeListOfRoleReq_))
	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}
	rsp.Body = rbac.GetAttributeListOfRole(req.Query)
	rsp.Code = code.Success
	return nil
}
