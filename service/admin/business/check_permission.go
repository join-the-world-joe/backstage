package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
)

func CheckPermission(ctx context.Context, req *admin.CheckPermissionReq, rsp *admin.CheckPermissionRsp) error {
	rsp.Major = req.Major
	rsp.Minor = req.Minor

	if req.Id <= 0 {
		log.Error("CheckPermission failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.Id)
	if len(roleList) <= 0 {
		log.Error("CheckPermission failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, req.Minor) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	rsp.Code = code.Success
	return nil
}
