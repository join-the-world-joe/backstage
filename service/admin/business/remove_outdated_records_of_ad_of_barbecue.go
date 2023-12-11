package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_barbecue"
	"backstage/common/db/mysql/backend/version_of_ad_of_barbecue"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"github.com/spf13/cast"
)

func RemoveOutdatedRecordsOfADOfBarbecue(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfBarbecueReq, rsp *admin.RemoveOutdatedRecordsOfADOfBarbecueRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.RemoveOutdatedRecordsOfADOfBarbecueReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	version, err := version_of_ad_of_barbecue.GetMaxId()
	if err != nil {
		log.Error(" version_of_ad_of_carousel.GetMaxId failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	ad_of_barbecue.RemoveOutdatedRecordsOfADOfBarbecue(version)

	rsp.Code = code.Success
	return nil
}
