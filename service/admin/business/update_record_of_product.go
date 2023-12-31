package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"github.com/spf13/cast"
)

func UpdateRecordOfProduct(ctx context.Context, req *admin.UpdateRecordOfProductReq, rsp *admin.UpdateRecordOfProductRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.UpdateRecordOfProductReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if req.UserId <= 0 || req.ProductId <= 0 {
		log.Error("UpdateRecordOfGood failure, req.UserId <= 0 || req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	fieldList := map[string]interface{}{}

	if len(req.Name) > 0 {
		fieldList["name"] = string(req.Name)
	}

	if req.BuyingPrice > 0 {
		fieldList["buying_price"] = req.BuyingPrice
	}

	if len(req.Vendor) > 0 {
		fieldList["vendor"] = string(req.Vendor)
	}

	if len(req.Contact) > 0 {
		fieldList["contact"] = string(req.Contact)
	}

	err := product.UpdateFieldListById(req.ProductId, fieldList)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Code = code.Success
	return nil
}
