package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/macro/abbreviation"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"golang.org/x/exp/slices"
)

type Permission struct {
	Name        string `json:"name"`
	Major       string `json:"major"`
	Minor       string `json:"minor"`
	Description string `json:"description"`
}

type PermissionListOutput struct {
	PermissionList []*Permission `json:"permission_list"`
	Length         int           `json:"length"`
}

func FetchPermissionListOfCondition(ctx context.Context, req *admin.FetchPermissionListOfConditionReq, rsp *admin.FetchPermissionListOfConditionRsp) error {
	if req.Id <= 0 {
		log.Error("FetchPermissionListOfCondition failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.Id)

	if len(roleList) <= 0 {
		log.Error("FetchPermissionListOfCondition failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchPermissionListOfConditionReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	nameList, majorList, minorList, descList := []string{}, []string{}, []string{}, []string{}

	if req.Behavior == 1 {
	} else if req.Behavior == 2 {
		roleList = []string{}
		if req.UserId <= 0 && len(req.Name) <= 0 && len(req.RoleList) <= 0 {
			rsp.Code = code.NoData
			return nil
		}
		if req.UserId > 0 {
			temp := user_role.GetRoleListByUserId(req.UserId)
			for _, v := range temp {
				if !slices.Contains(roleList, v) {
					roleList = append(roleList, v)
				}
			}
		}
		if len(req.RoleList) > 0 {
			for _, v := range req.RoleList {
				if !slices.Contains(roleList, string(v)) {
					roleList = append(roleList, string(v))
				}
			}
		}
		if len(req.Name) > 0 {
			name, major, minor, desc := rbac.GetPermissionByName(req.Name)
			if len(name) > 0 && !slices.Contains(nameList, name) {
				nameList = append(nameList, name)
				majorList = append(majorList, major)
				minorList = append(minorList, minor)
				descList = append(descList, desc)
			}
		}
	} else {
		rsp.Code = code.InvalidData
		return nil
	}

	plo := &PermissionListOutput{
		PermissionList: []*Permission{},
		Length:         0,
	}

	tempNameList, tempMajorList, tempMinorList, tempDescList := rbac.GetPermissionListOfRoleList(roleList)
	for k, v := range tempNameList {
		if !slices.Contains(nameList, v) {
			nameList = append(nameList, tempNameList[k])
			desc := abbreviation.NA
			if len(tempDescList[k]) > 0 {
				desc = tempDescList[k]
			}
			permission := &Permission{
				Name:        tempNameList[k],
				Major:       tempMajorList[k],
				Minor:       tempMinorList[k],
				Description: desc,
			}
			plo.PermissionList = append(plo.PermissionList, permission)
			plo.Length += 1
		}
	}

	bytes, err := json.Marshal(plo)
	if err != nil {
		log.Error("FetchPermissionListOfCondition failure, err: ", err.Error())
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
