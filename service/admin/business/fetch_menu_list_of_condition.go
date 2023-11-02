package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"backstage/utils/json"
	"context"
	"github.com/spf13/cast"
	"golang.org/x/exp/slices"
)

type Menu struct {
}

func FetchMenuListOfCondition(ctx context.Context, req *admin.FetchMenuListOfConditionReq, rsp *admin.FetchMenuListOfConditionRsp) error {
	if req.Id <= 0 {
		log.Error("FetchMenuListOfCondition failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.Id)

	if len(roleList) <= 0 {
		log.Error("FetchMenuListOfCondition failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchMenuListOfConditionReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	if req.Behavior == 1 {
	} else if req.Behavior == 2 {
		if req.UserId <= 0 && len(req.RoleList) <= 0 && len(req.Menu) <= 0 && len(req.Parent) <= 0 { // no conditions
			rsp.Code = code.InvalidData
			return nil
		}
		roleList = []string{} // reset roleList
		if req.UserId > 0 {   // user_id condition
			tempRoleList := user_role.GetRoleListByUserId(req.UserId)
			for _, v := range tempRoleList {
				if !slices.Contains(roleList, v) {
					roleList = append(roleList, v)
				}
			}
		}
		if len(req.RoleList) > 0 { // role_list condition
			for _, v := range req.RoleList {
				if !slices.Contains(roleList, v) {
					roleList = append(roleList, v)
				}
			}
		}
		if len(req.Menu) > 0 {

		}
		if len(req.Parent) > 0 {

		}
	} else {
		rsp.Code = code.InvalidData
		return nil
	}

	menuList, itemList, descList := rbac.GetMenuListOfRoleList(roleList)

	js := json.New()
	for i, length := 0, len(menuList); i < length; i++ {
		js.SetPath([]string{menuList[i], "item_list"}, itemList[i])
		js.SetPath([]string{menuList[i], "description_list"}, descList[i])
	}

	bytes, err := js.Encode()
	if err != nil {
		rsp.Code = code.InternalError
		log.Error("FetchMenuListOfCondition.js.Encode failure, err: ", err.Error())
		return nil
	}

	rsp.Code = code.Success
	rsp.Body = bytes
	return nil
}
