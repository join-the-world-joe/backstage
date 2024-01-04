package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user"
	"backstage/common/db/mysql/backend/user_role"
	role2 "backstage/common/macro/role"
	"backstage/common/protocol/admin"
	"backstage/global/crypto"
	"backstage/global/log"
	"backstage/global/rbac"
	"backstage/utils/bcrypt"
	"context"
	"github.com/spf13/cast"
	"golang.org/x/exp/slices"
)

func UpdateRecordOfUser(ctx context.Context, req *admin.UpdateRecordOfUserReq, rsp *admin.UpdateRecordOfUserRsp) error {
	if req.Id <= 0 {
		log.Error("UpdateRecordOfUser failure, req.UserId <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	if len(req.Name) <= 0 ||
		len(req.PhoneNumber) <= 0 ||
		len(req.CountryCode) <= 0 ||
		(req.Status != 1 && req.Status != 2) {
		rsp.Code = code.NoData
		return nil
	}

	selfRoleList := user_role.GetRoleListByUserId(req.Id)

	if len(selfRoleList) <= 0 {
		log.Error("UpdateRecordOfUser failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range selfRoleList {
		if rbac.HasPermission(v, cast.ToInt(admin.UpdateRecordOfUserReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	fieldList := map[string]interface{}{}
	if len(req.Name) > 0 {
		fieldList["name"] = req.Name
	}

	if len(req.CountryCode) > 0 {
		fieldList["country_code"] = req.CountryCode
	}

	if len(req.PhoneNumber) > 0 {
		fieldList["phone_number"] = req.PhoneNumber
	}

	if req.Status > 0 {
		fieldList["status"] = req.Status
	}

	if len(req.Password) > 0 {
		passwordBytes, err := crypto.RSADecrypt(req.Password)
		if err != nil {
			rsp.Code = code.UnsupportedType
			return nil
		}

		password, err := bcrypt.PasswordHash(string(passwordBytes))
		if err != nil {
			rsp.Code = code.UnsupportedType
			return nil
		}
		fieldList["password"] = password
	}

	err := user.UpdateFieldListById(req.UserId, fieldList)
	if err != nil {
		log.Error("UpdateUserRecord failure, err: ", err.Error())
		rsp.Code = code.InternalError
		return nil
	}

	/*
				in roles of user	<(rank)		in roles of manager 	in roles selected		operation
		role1			0								1					1					new role
		role2			1								1					0					remove role
	*/

	if slices.Contains(selfRoleList, role2.Administrator) {
		selfRoleList = rbac.GetRoleListLERank(rbac.GetTopRankOfRoleList(selfRoleList))
	} else {
		for _, role := range selfRoleList {
			if name, department, _, _, b := rbac.GetRole(role); b {
				selfRoleList = rbac.GetRoleListLERankInDepartment(rbac.GetTopRankOfRoleList([]string{name}), department)
			}
		}
	}
	roleList := user_role.GetRoleListByUserId(req.UserId)
	//fmt.Println("roleList: ", roleList, ", len: = ", len(roleList))
	//fmt.Println("req.RoleList: ", req.RoleList, ", len: = ", len(req.RoleList))
	//fmt.Println("selfRoleList: ", selfRoleList, ", len: = ", len(selfRoleList))
	for _, roleName := range req.RoleList {
		if slices.Contains(selfRoleList, roleName) {
			if !slices.Contains(roleList, roleName) {
				if id := user_role.GetIdByRoleAndVisible(roleName, 0); id > 0 {
					err = user_role.UpdateVisibleByRole(roleName, 1)
					if err != nil {
						log.Error("UpdateUserRecord.user_role.UpdateVisibleByRole failure, err: ", err.Error())
						rsp.Code = code.InternalError
						return nil
					}
				} else {
					_, err = user_role.InsertModel(
						&user_role.Model{
							UserId: req.UserId,
							Role:   roleName,
						},
					)
					if err != nil {
						log.Error("UpdateUserRecord.user_role.InsertModel failure, err: ", err.Error())
						rsp.Code = code.InternalError
						return nil
					}
				}
			}
		}
	}
	for _, roleName := range roleList {
		if !slices.Contains(req.RoleList, roleName) &&
			slices.Contains(selfRoleList, roleName) &&
			slices.Contains(roleList, roleName) {
			err = user_role.UpdateVisibleByRole(roleName, 0)
			if err != nil {
				log.Error("UpdateUserRecord.user_role.UpdateVisibleByRole failure, err: ", err.Error())
				rsp.Code = code.InternalError
				return nil
			}
		}
	}

	rsp.Code = code.Success
	return nil
}
