package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/crypto"
	"backstage/global/log"
	"backstage/global/rbac"
	"backstage/utils/bcrypt"
	"context"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func InsertUserRecord(ctx context.Context, req *admin.InsertUserRecordReq, rsp *admin.InsertUserRecordRsp) error {
	if req.Id <= 0 {
		log.Error("InsertUserRecord failure, req.UserId <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	if len(req.Name) <= 0 ||
		len(req.Password) <= 0 ||
		len(req.PhoneNumber) <= 0 ||
		len(req.CountryCode) <= 0 ||
		(req.Status != 1 && req.Status != 2) {
		rsp.Code = code.NoData
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.Id)

	if len(roleList) <= 0 {
		log.Error("InsertUserRecord failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.InsertUserRecordReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	//todo: check if user valid
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
	usr, err := user.InsertModel(
		&user.Model{
			Name:        string(req.Name),
			Status:      req.Status,
			CountryCode: req.CountryCode,
			PhoneNumber: req.PhoneNumber,
			Password:    password,
			MemberId:    uuid.New().String(),
		},
	)
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	rank := rbac.GetTopRankOfRoleList(roleList)
	for _, role := range req.RoleList {
		if rank >= rbac.GetRankOfRole(role) {
			_, err := user_role.InsertModel(
				&user_role.Model{
					UserId: usr.Id,
					Role:   role,
				},
			)
			if err != nil {
				rsp.Code = code.DatabaseFailure
				return nil
			}
		}
	}

	rsp.Code = code.Success
	return nil
}
