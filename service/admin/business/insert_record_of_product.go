package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"context"
	"github.com/spf13/cast"
)

func InsertRecordOfProduct(ctx context.Context, req *admin.InsertRecordOfProductReq, rsp *admin.InsertRecordOfProductRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfProductReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if req.BuyingPrice <= 0 || len(req.Name) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	_, err := product.InsertModel(&product.Model{
		Name:        string(req.Name),
		Vendor:      string(req.Vendor),
		Contact:     string(req.Contact),
		BuyingPrice: req.BuyingPrice,
	})
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	rsp.Code = code.Success
	return nil
}
