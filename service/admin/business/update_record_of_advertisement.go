package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func UpdateRecordOfAdvertisement(ctx context.Context, req *admin.UpdateRecordOfAdvertisementReq, rsp *admin.UpdateRecordOfAdvertisementRsp) error {
	if req.UserId <= 0 || req.ProductId <= 0 {
		log.Error("UpdateRecordOfAdvertisement failure, req.UserId <= 0 || req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("UpdateRecordOfAdvertisement failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.UpdateRecordOfAdvertisementReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	fieldList := map[string]interface{}{}

	if len(req.Url) > 0 {
		fieldList["url"] = string(req.Url)
	}

	if len(req.Name) > 0 {
		fieldList["name"] = string(req.Name)
	}

	if len(req.Title) > 0 {
		fieldList["title"] = string(req.Title)
	}

	if req.Stock > 0 {
		fieldList["stock"] = req.Stock
	}

	if req.ProductId > 0 {
		fieldList["product_id"] = req.ProductId
	}

	if req.SellingPrice > 0 {
		fieldList["selling_price"] = req.SellingPrice
	}

	if len(req.SellingPoint) > 0 {
		fieldList["selling_point"] = string(req.SellingPoint)
	}

	if len(req.PlaceOfOrigin) > 0 {
		fieldList["place_of_origin"] = string(req.PlaceOfOrigin)
	}

	fieldList["status"] = req.Status

	err := advertisement.UpdateFieldListById(req.ProductId, fieldList)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Code = code.Success
	return nil
}
