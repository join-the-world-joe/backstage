package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"github.com/spf13/cast"
)

func InsertRecordOfGood(ctx context.Context, req *admin.InsertRecordOfGoodReq, rsp *admin.InsertRecordOfGoodRsp) error {
	if req.BuyingPrice <= 0 || len(req.Name) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	if req.UserId <= 0 {
		log.Error("InsertRecordOfGood failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("InsertRecordOfGood failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.InsertRecordOfGoodReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
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
