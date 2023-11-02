package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_bbq_products"
	"backstage/common/protocol/advertisement"
	"context"
	"github.com/spf13/cast"
)

func FetchVersionOfADOfBBQProducts(ctx context.Context, req *advertisement.FetchVersionOfADOfBBQProductsReq, rsp *advertisement.FetchVersionOfADOfBBQProductsRsp) error {
	id, err := version_of_ad_of_bbq_products.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	rsp.Code = code.Success
	rsp.Version = cast.ToInt(id)
	return nil
}
