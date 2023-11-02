package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_hot_deals"
	"backstage/common/protocol/advertisement"
	"context"
	"github.com/spf13/cast"
)

func FetchVersionOfADOfHotDeals(ctx context.Context, req *advertisement.FetchVersionOfADOfHotDealsReq, rsp *advertisement.FetchVersionOfADOfHotDealsRsp) error {
	id, err := version_of_ad_of_hot_deals.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	rsp.Code = code.Success
	rsp.Version = cast.ToInt(id)
	return nil
}
