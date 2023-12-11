package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_deals"
	"backstage/common/db/mysql/backend/version_of_ad_of_deals"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"github.com/spf13/cast"
)

func RemoveOutdatedRecordsOfADOfDeals(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfDealsReq, rsp *admin.RemoveOutdatedRecordsOfADOfDealsRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.RemoveOutdatedRecordsOfADOfDealsReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	version, err := version_of_ad_of_deals.GetMaxId()
	if err != nil {
		log.Error(" version_of_ad_of_deals.GetMaxId failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	ad_of_deals.RemoveOutdatedRecordsOfADOfDeals(version)

	rsp.Code = code.Success
	return nil
}
