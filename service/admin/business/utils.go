package business

import (
	"backstage/common/db/mysql/backend/user_role"
	"backstage/global/log"
	"backstage/global/rbac"
)

func hasPermission(major, minor int, userId int64) bool {
	if userId <= 0 {
		log.Error("hasPermission failure, userId <= 0")
		return false
	}

	roleList := user_role.GetRoleListByUserId(userId)
	if len(roleList) <= 0 {
		log.Error("hasPermission failure, len(roleList) <= 0")
		return false
	}

	// check if role_list has permission
	has := false
	for _, v := range roleList {
		if rbac.HasPermission(v, minor) {
			has = true
			break
		}
	}

	if !has {
		return false
	}
	return true
}
