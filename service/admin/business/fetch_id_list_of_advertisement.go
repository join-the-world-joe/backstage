package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

type OutputOfIdListOfAdvertisement struct {
	Behavior              int     `json:"behavior"`
	IdListOfAdvertisement []int64 `json:"id_list_of_advertisement"`
}

func FetchIdListOfAdvertisement(ctx context.Context, req *admin.FetchIdListOfAdvertisementReq, rsp *admin.FetchIdListOfAdvertisementRsp) error {
	if len(req.AdvertisementName) <= 0 && (req.Behavior != 1 && req.Behavior != 2) {
		rsp.Code = code.InvalidData
		return nil
	}

	if req.UserId <= 0 {
		log.Error("FetchIdListOfAdvertisement failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.UserId)

	if len(roleList) <= 0 {
		log.Error("FetchIdListOfAdvertisement failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchIdListOfAdvertisementReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	output := &OutputOfIdListOfAdvertisement{
		Behavior:              req.Behavior,
		IdListOfAdvertisement: []int64{},
	}

	if req.Behavior == 1 {
		idList, err := advertisement.GetIdListInTable()
		if err != nil {
			rsp.Code = code.DatabaseFailure
			return nil
		}
		for _, v := range idList {
			output.IdListOfAdvertisement = append(output.IdListOfAdvertisement, v)
		}
	} else if req.Behavior == 2 {
		if len(req.AdvertisementName) > 0 {
			ml, err := advertisement.GetModelListByName(string(req.AdvertisementName))
			if err != nil {
				log.Error(fmt.Sprintf("FetchIdListOfAdvertisement.product.GetModelListByName failure, err: %v", err.Error()))
			} else {
				for _, v := range ml {
					output.IdListOfAdvertisement = append(output.IdListOfAdvertisement, v.Id)
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
