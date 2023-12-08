package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/selling_point_of_advertisement"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"fmt"
	"github.com/spf13/cast"
)

func InsertRecordOfAdvertisement(ctx context.Context, req *admin.InsertRecordOfAdvertisementReq, rsp *admin.InsertRecordOfAdvertisementRsp) error {
	if req.ProductId <= 0 || len(req.Name) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	if req.UserId <= 0 {
		log.Error("InsertRecordOfAdvertisement failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("InsertRecordOfAdvertisement failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.InsertRecordOfAdvertisementReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	m, err := advertisement.InsertModel(&advertisement.Model{
		Url:           string(req.Url),
		Title:         string(req.Title),
		Name:          string(req.Name),
		PlaceOFOrigin: string(req.PlaceOfOrigin),
		SellingPrice:  req.SellingPrice,
		Stock:         req.Stock,
		ProductId:     req.ProductId,
		Status:        req.Status,
		Description:   string(req.Description),
	})
	if err != nil {
		rsp.Code = code.DatabaseFailure
		log.Error(fmt.Sprintf("InsertRecordOfAdvertisement failure, err: %v", err.Error()))
		return nil
	}

	if len(req.SellingPoints) > 0 {
		for _, v := range req.SellingPoints {
			_, err := selling_point_of_advertisement.InsertModel(&selling_point_of_advertisement.Model{
				SellingPoint:    string(v),
				AdvertisementId: m.Id,
			})
			if err != nil {
				log.Error("InsertRecordOfAdvertisement.selling_point_of_advertisement.InsertModel failure, err: ", err.Error())
				continue
			}
		}
	}

	rsp.Code = code.Success
	return nil
}
