package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"fmt"
	"github.com/spf13/cast"
)

func SoftDeleteRecordsOfAdvertisement(ctx context.Context, req *admin.SoftDeleteRecordsOfAdvertisementReq, rsp *admin.SoftDeleteRecordsOfAdvertisementRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.SoftDeleteRecordsOfAdvertisementReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if len(req.AdvertisementIdList) <= 0 {
		rsp.Code = code.InvalidData
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
