package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_snacks"
	"backstage/common/db/mysql/backend/version_of_ad_of_snacks"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"github.com/spf13/cast"
)

func RemoveOutdatedRecordsOfADOfSnacks(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfSnacksReq, rsp *admin.RemoveOutdatedRecordsOfADOfSnacksRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.RemoveOutdatedRecordsOfADOfSnacksReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	version, err := version_of_ad_of_snacks.GetMaxId()
	if err != nil {
		log.Error(" version_of_ad_of_snacks.GetMaxId failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	ad_of_snacks.RemoveOutdatedRecordsOfADOfSnacks(version)

	rsp.Code = code.Success
	return nil
}
