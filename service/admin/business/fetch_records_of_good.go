package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/macro/timestamp"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

type RecordOfGood struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	Vendor      string `json:"vendor"`
	Contact     string `json:"contact"`
	Description string `json:"description"`
	BuyingPrice int    `json:"buying_price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type OutputOfRecordsOfGood struct {
	Behavior      int                     `json:"behavior"`
	RecordsOfGood map[int64]*RecordOfGood `json:"records_of_good"`
}

func FetchRecordsOfGood(ctx context.Context, req *admin.FetchRecordsOfGoodReq, rsp *admin.FetchRecordsOfGoodRsp) error {
	if len(req.ProductIdList) <= 0 {
		rsp.Code = code.InvalidData
		return nil
	}

	if req.UserId <= 0 {
		log.Error("FetchRecordsOfGood failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("FetchRecordsOfGood failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchRecordsOfGoodReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	output := &OutputOfRecordsOfGood{
		RecordsOfGood: map[int64]*RecordOfGood{},
	}

	ml, err := product.GetModelListByIdList(req.ProductIdList)
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	for _, m := range ml {
		output.RecordsOfGood[m.Id] = &RecordOfGood{
			Id:          m.Id,
			Name:        m.Name,
			Status:      m.Status,
			Vendor:      m.Vendor,
			Contact:     m.Contact,
			BuyingPrice: m.BuyingPrice,
			Description: m.Description,
			CreatedAt:   m.CreatedAt.Format(timestamp.YYMDHMS),
			UpdatedAt:   m.UpdatedAt.Format(timestamp.YYMDHMS),
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
