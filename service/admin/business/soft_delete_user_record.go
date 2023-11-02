package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func SoftDeleteUserRecord(ctx context.Context, req *admin.SoftDeleteUserRecordReq, rsp *admin.SoftDeleteUserRecordRsp) error {
	if req.Id <= 0 {
		log.Error("SoftDeleteUserRecord failure, req.UserId <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	if len(req.UserIdList) <= 0 {
		rsp.Code = code.NoData
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.Id)

	if len(roleList) <= 0 {
		log.Error("SoftDeleteUserRecord failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.InsertUserRecordReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	for _, userId := range req.UserIdList {
		err := user.UpdateVisibleById(userId, 0)
		if err != nil {
			log.Error("SoftDeleteUserRecord.user.UpdateVisibleById failure, err: ", err.Error())
		}
		err = user_role.UpdateVisibleByUserId(userId, 0)
		if err != nil {
			log.Error("SoftDeleteUserRecord.user_role.UpdateVisibleById failure, err: ", err.Error())
		}
	}

	rsp.Code = code.Success
	return nil
}
