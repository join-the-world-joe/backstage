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
	"strings"
)

func UpdateRecordOfAdvertisement(ctx context.Context, req *admin.UpdateRecordOfAdvertisementReq, rsp *admin.UpdateRecordOfAdvertisementRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfAdvertisementReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	if req.UserId <= 0 || req.ProductId <= 0 {
		log.Error("UpdateRecordOfAdvertisement failure, req.UserId <= 0 || req.Id <= 0")
		rsp.Code = code.InternalError
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

	if len(req.PlaceOfOrigin) > 0 {
		fieldList["place_of_origin"] = string(req.PlaceOfOrigin)
	}

	fieldList["status"] = req.Status

	if len(req.Description) > 0 {
		fieldList["description"] = string(req.Description)
	}

	err := advertisement.UpdateFieldListById(req.Id, fieldList)
	if err != nil {
		log.Error("UpdateRecordOfAdvertisement.advertisement.UpdateFieldListById failure, err: ", err.Error())
		rsp.Code = code.InternalError
		return nil
	}

	pointsToBeAdded := []string{}
	idListToBeDeleted := []int64{}
	requestedSellingPoints := []string{}
	oriSellingPoints := map[string]int64{}
	fmt.Println("req: ", req)
	ml, err := selling_point_of_advertisement.GetModelListByAdvertisementId(req.Id)
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	for _, v := range ml {
		oriSellingPoints[v.SellingPoint] = v.Id
	}
	if len(req.SellingPoints) > 0 {
		for _, v := range req.SellingPoints {
			requestedSellingPoints = append(requestedSellingPoints, string(v))
		}
	}
	{
		// new point
		for _, v := range requestedSellingPoints {
			if _, exist := oriSellingPoints[v]; !exist {
				pointsToBeAdded = append(pointsToBeAdded, v)
			}
		}
		// to be deleted
		for oriPoint, oriId := range oriSellingPoints {
			found := false
			for _, reqPoint := range requestedSellingPoints {
				if strings.Compare(reqPoint, oriPoint) == 0 {
					found = true
				}
			}
			if !found {
				idListToBeDeleted = append(idListToBeDeleted, oriId)
			}
		}
		if len(idListToBeDeleted) > 0 {
			err = selling_point_of_advertisement.UpdateVisibleByIdList(idListToBeDeleted, 0)
			if err != nil {
				rsp.Code = code.DatabaseFailure
				return nil
			}
		}
		for _, v := range pointsToBeAdded {
			_, err = selling_point_of_advertisement.InsertModel(&selling_point_of_advertisement.Model{
				AdvertisementId: req.Id,
				SellingPoint:    v,
			})
			if err != nil {
				rsp.Code = code.DatabaseFailure
				return nil
			}
		}
	}

	rsp.Code = code.Success
	return nil
}
