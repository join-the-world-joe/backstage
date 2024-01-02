package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"fmt"
	"github.com/spf13/cast"
)

func SoftDeleteRecordsOfProduct(ctx context.Context, req *admin.SoftDeleteRecordsOfProductReq, rsp *admin.SoftDeleteRecordsOfProductRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.SoftDeleteRecordsOfProductReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if len(req.ProductIdList) <= 0 {
		rsp.Code = code.InvalidData
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
