package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

type OutputOfIdListOfGood struct {
	Behavior     int     `json:"behavior"`
	IdListOfGood []int64 `json:"id_list_of_good"`
}

func FetchIdListOfGood(ctx context.Context, req *admin.FetchIdListOfGoodReq, rsp *admin.FetchIdListOfGoodRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.FetchIdListOfGoodReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}
	if len(req.ProductName) <= 0 && req.CategoryId <= 0 && (req.Behavior != 1 && req.Behavior != 2) {
		rsp.Code = code.InvalidData
		return nil
	}

	if req.UserId <= 0 {
		log.Error("FetchIdListOfGood failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("FetchIdListOfGood failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchIdListOfGoodReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	output := &OutputOfIdListOfGood{
		Behavior:     req.Behavior,
		IdListOfGood: []int64{},
	}

	if req.Behavior == 1 {
		idList, err := product.GetIdListInTable()
		if err != nil {
			rsp.Code = code.DatabaseFailure
			return nil
		}
		for _, v := range idList {
			output.IdListOfGood = append(output.IdListOfGood, v)
		}
	} else if req.Behavior == 2 {
		if len(req.ProductName) > 0 {
			ml, err := product.GetModelListByName(string(req.ProductName))
			if err != nil {
				log.Error(fmt.Sprintf("FetchIdListOfGood.product.GetModelListByName failure, err: %v", err.Error()))
			} else {
				for _, v := range ml {
					output.IdListOfGood = append(output.IdListOfGood, v.Id)
				}
			}
		}
	} else {
		rsp.Code = code.InvalidData
		return nil
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
