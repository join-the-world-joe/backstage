package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"context"
	"github.com/spf13/cast"
)

func InsertRecordOfGood(ctx context.Context, req *admin.InsertRecordOfGoodReq, rsp *admin.InsertRecordOfGoodRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfGoodReq_),
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
		Status:      req.Status,
		Contact:     string(req.Contact),
		BuyingPrice: req.BuyingPrice,
		Description: string(req.Description),
	})
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	rsp.Code = code.Success
	return nil
}
