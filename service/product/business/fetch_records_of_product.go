package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/macro/timestamp"
	product2 "backstage/common/protocol/product"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

type RecordOfProduct struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Vendor      string `json:"vendor"`
	Contact     string `json:"contact"`
	BuyingPrice int    `json:"buying_price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type OutputOfRecordsOfProduct struct {
	Behavior         int                        `json:"behavior"`
	RecordsOfProduct map[int64]*RecordOfProduct `json:"records_of_product"`
}

func FetchRecordsOfProduct(ctx context.Context, req *product2.FetchRecordsOfProductReq, rsp *product2.FetchRecordsOfProductRsp) error {
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
		if rbac.HasPermission(v, cast.ToInt(product2.FetchRecordsOfProductReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	output := &OutputOfRecordsOfProduct{
		RecordsOfProduct: map[int64]*RecordOfProduct{},
	}

	ml, err := product.GetModelListByIdList(req.ProductIdList)
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	for _, m := range ml {
		output.RecordsOfProduct[m.Id] = &RecordOfProduct{
			Id:          m.Id,
			Name:        m.Name,
			Vendor:      m.Vendor,
			Contact:     m.Contact,
			BuyingPrice: m.BuyingPrice,
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
