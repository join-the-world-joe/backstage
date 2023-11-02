package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/selling_point_of_advertisement"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/macro/timestamp"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

type RecordOfAdvertisement struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	Title         string   `json:"title"`
	Status        int      `json:"status"`
	SellingPrice  int      `json:"selling_price"`
	PlaceOfOrigin string   `json:"place_of_origin"`
	SellingPoint  []string `json:"selling_point"`
	Url           string   `json:"url"`
	Stock         int      `json:"stock"`
	ProductId     int64    `json:"product_id"`
	BuyingPrice   int      `json:"buying_price"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type OutputOfRecordsOfAdvertisement struct {
	RecordsOfAdvertisement map[int64]*RecordOfAdvertisement `json:"records_of_advertisement"`
}

func FetchRecordsOfAdvertisement(ctx context.Context, req *admin.FetchRecordsOfAdvertisementReq, rsp *admin.FetchRecordsOfAdvertisementRsp) error {
	if len(req.AdvertisementIdList) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	if req.UserId <= 0 {
		log.Error("FetchRecordsOfAdvertisement failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("FetchRecordsOfAdvertisement failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchRecordsOfAdvertisementReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	output := &OutputOfRecordsOfAdvertisement{
		RecordsOfAdvertisement: map[int64]*RecordOfAdvertisement{},
	}

	ml, err := advertisement.GetModelListByIdList(req.AdvertisementIdList)
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	for _, m := range ml {
		pml, err := selling_point_of_advertisement.GetModelListByProductId(m.ProductId)
		if err != nil {
			log.Error("business.selling_point_of_advertisement.GetModelListByProductId failure, err: ", err.Error())
			continue
		}
		points := []string{}
		for _, v := range pml {
			points = append(points, v.SellingPoint)
		}
		output.RecordsOfAdvertisement[m.Id] = &RecordOfAdvertisement{
			Id:            m.Id,
			Name:          m.Name,
			Title:         m.Title,
			Status:        m.Status,
			SellingPrice:  m.SellingPrice,
			PlaceOfOrigin: m.PlaceOFOrigin,
			SellingPoint:  points,
			Url:           m.Url,
			Stock:         m.Stock,
			ProductId:     m.ProductId,
			CreatedAt:     m.CreatedAt.Format(timestamp.YYMDHMS),
			UpdatedAt:     m.UpdatedAt.Format(timestamp.YYMDHMS),
		}
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
