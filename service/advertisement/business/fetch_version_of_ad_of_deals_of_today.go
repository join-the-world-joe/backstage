package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_deals_of_today"
	"backstage/common/protocol/advertisement"
	"context"
	"github.com/spf13/cast"
)

func FetchVersionOfADOfDealsOfToday(ctx context.Context, req *advertisement.FetchVersionOfADOfDealsOfTodayReq, rsp *advertisement.FetchVersionOfADOfDealsOfTodayRsp) error {
	id, err := version_of_ad_of_deals_of_today.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	rsp.Code = code.Success
	rsp.Version = cast.ToInt(id)
	return nil
}
