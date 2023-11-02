package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"fmt"
	"github.com/spf13/cast"
)

func SoftDeleteRecordsOfGood(ctx context.Context, req *admin.SoftDeleteRecordsOfGoodReq, rsp *admin.SoftDeleteRecordsOfGoodRsp) error {
	if len(req.ProductIdList) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}
	if req.UserId <= 0 {
		log.Error("SoftDeleteRecordOfGood failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("SoftDeleteRecordOfGood failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.SoftDeleteRecordsOfGoodReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	err := product.UpdateVisibleByIdList(req.ProductIdList, 0)
	if err != nil {
		log.Error(fmt.Sprintf("SoftDeleteRecordOfGood failure, err: %v", err.Error()))
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Code = code.Success
	return nil
}
