package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"github.com/spf13/cast"
)

func SoftDeleteRecordOfUser(ctx context.Context, req *admin.SoftDeleteRecordOfUserReq, rsp *admin.SoftDeleteRecordOfUserRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.SoftDeleteRecordOfUserReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if len(req.UserIdList) <= 0 {
		rsp.Code = code.NoData
		return nil
	}

	for _, userId := range req.UserIdList {
		err := user.UpdateVisibleById(userId, 0)
		if err != nil {
			log.Error("user.UpdateVisibleById failure, err: ", err.Error())
		}
		err = user_role.UpdateVisibleByUserId(userId, 0)
		if err != nil {
			log.Error("user_role.UpdateVisibleById failure, err: ", err.Error())
		}
	}

	rsp.Code = code.Success
	return nil
}
