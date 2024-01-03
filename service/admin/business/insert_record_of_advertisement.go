package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/selling_point_of_advertisement"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

type OutputOfRecordOfAdvertisement struct {
	AdvertisementId int64 `json:"advertisement_id"`
}

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
		CoverImage:    string(req.CoverImage),
		FirstImage:    string(req.FirstImage),
		SecondImage:   string(req.SecondImage),
		ThirdImage:    string(req.ThirdImage),
		FourthImage:   string(req.FourthImage),
		FifthImage:    string(req.FifthImage),
		OSSFolder:     string(req.OSSFolder),
		OSSPath:       string(req.OSSPath),
		Title:         string(req.Title),
		Name:          string(req.Name),
		PlaceOFOrigin: string(req.PlaceOfOrigin),
		SellingPrice:  req.SellingPrice,
		Stock:         req.Stock,
		Status:        req.Status,
		ProductId:     req.ProductId,
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

	output := &OutputOfRecordOfAdvertisement{
		AdvertisementId: m.Id,
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
