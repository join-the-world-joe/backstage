package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/selling_point_of_advertisement"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"fmt"
	"github.com/spf13/cast"
)

func InsertRecordOfAdvertisement(ctx context.Context, req *admin.InsertRecordOfAdvertisementReq, rsp *admin.InsertRecordOfAdvertisementRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfAdvertisementReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if req.ProductId <= 0 || len(req.Name) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	m, err := advertisement.InsertModel(&advertisement.Model{
		Image:         string(req.Image),
		Title:         string(req.Title),
		Name:          string(req.Name),
		PlaceOFOrigin: string(req.PlaceOfOrigin),
		SellingPrice:  req.SellingPrice,
		Stock:         req.Stock,
		ProductId:     req.ProductId,
		Thumbnail:     string(req.Thumbnail),
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
