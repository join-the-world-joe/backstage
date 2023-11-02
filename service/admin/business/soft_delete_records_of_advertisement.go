package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"fmt"
	"github.com/spf13/cast"
)

func SoftDeleteRecordsOfAdvertisement(ctx context.Context, req *admin.SoftDeleteRecordsOfAdvertisementReq, rsp *admin.SoftDeleteRecordsOfAdvertisementRsp) error {
	if len(req.AdvertisementIdList) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}
	if req.UserId <= 0 {
		log.Error("SoftDeleteRecordsOfAdvertisement failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("SoftDeleteRecordsOfAdvertisement failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.SoftDeleteRecordsOfAdvertisementReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	err := advertisement.UpdateVisibleByIdList(req.AdvertisementIdList, 0)
	if err != nil {
		log.Error(fmt.Sprintf("SoftDeleteRecordsOfAdvertisement failure, err: %v", err.Error()))
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Code = code.Success
	return nil
}
