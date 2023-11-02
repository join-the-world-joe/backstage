package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/macro/abbreviation"
	role2 "backstage/common/macro/role"
	"backstage/common/macro/timestamp"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"golang.org/x/exp/slices"
	"sort"
	"strings"
)

type UserListOutput struct {
	Length   int           `json:"length"`
	UserList []*UserOutput `json:"user_list"`
}

type UserOutput struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Account     string `json:"account"`
	Email       string `json:"email"`
	Visible     string `json:"visible"`
	Department  string `json:"department"`
	Password    string `json:"password"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func transform(fieldList []string, input []*user.Model) *UserListOutput {
	userListOutput := &UserListOutput{
		UserList: []*UserOutput{},
		Length:   0,
	}
	for _, v := range input {
		o := &UserOutput{
			Id:          abbreviation.NA,
			Name:        abbreviation.NA,
			Account:     abbreviation.NA,
			Email:       abbreviation.NA,
			Visible:     abbreviation.NA,
			Department:  abbreviation.NA,
			Password:    abbreviation.NA,
			CountryCode: abbreviation.NA,
			PhoneNumber: abbreviation.NA,
			Status:      abbreviation.NA,
			CreatedAt:   abbreviation.NA,
			UpdatedAt:   abbreviation.NA,
		}
		if slices.Contains(fieldList, "id") {
			o.Id = cast.ToString(v.Id)
		}
		if slices.Contains(fieldList, "name") {
			o.Name = v.Name
		}
		if slices.Contains(fieldList, "account") {
			o.Account = v.Account
		}
		if slices.Contains(fieldList, "email") {
			o.Email = v.Email
		}
		if slices.Contains(fieldList, "visible") {
			if v.Visible == 1 {
				o.Visible = "enable"
			} else {
				o.Visible = "disable"
			}
		}
		if slices.Contains(fieldList, "department") {
			o.Department = v.Department
		}
		if slices.Contains(fieldList, "country_code") {
			o.CountryCode = v.CountryCode
		}
		if slices.Contains(fieldList, "phone_number") {
			o.PhoneNumber = v.PhoneNumber
		}
		if slices.Contains(fieldList, "status") {
			o.Status = cast.ToString(v.Status)
		}
		if slices.Contains(fieldList, "created_at") {
			o.CreatedAt = v.CreatedAt.Format(timestamp.YYMDHMS)
		}
		if slices.Contains(fieldList, "updated_at") {
			o.UpdatedAt = v.UpdatedAt.Format(timestamp.YYMDHMS)
		}
		userListOutput.UserList = append(userListOutput.UserList, o)
	}
	return userListOutput
}

func FetchUserListOfCondition(ctx context.Context, req *admin.FetchUserListOfConditionReq, rsp *admin.FetchUserListOfConditionRsp) error {
	if req.Id <= 0 {
		log.Error("FetchUserListOfCondition failure, req.UserId <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	selfRoleList := user_role.GetRoleListByUserId(req.Id)

	if len(selfRoleList) <= 0 {
		log.Error("FetchUserListOfCondition failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range selfRoleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchUserListOfConditionReq_)) {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		rsp.Code = code.AccessDenied
		return nil
	}

	userIdList := []int64{}
	userModelMap := map[int64]*user.Model{}
	rank := rbac.GetTopRankOfRoleList(selfRoleList)
	if req.Behavior == 1 {
		roleList := []string{}
		if slices.Contains(selfRoleList, role2.Administrator) {
			roleList = rbac.GetRoleListLERank(rbac.GetTopRankOfRoleList(selfRoleList))
		} else {
			for _, role := range selfRoleList {
				if name, department, _, _, b := rbac.GetRole(role); b {
					roleList = rbac.GetRoleListLERankInDepartment(rbac.GetTopRankOfRoleList([]string{name}), department)
				}
			}
		}
		if uidList := user_role.GetUserIdListByRoleList(roleList); len(uidList) > 0 {
			for _, uid := range uidList {
				if uid == req.Id { // skip the record of myself
					continue
				}
				usr, err := user.GetModelById(uid)
				if err == nil {
					if !slices.Contains(userIdList, usr.Id) {
						//userModelList = append(userModelList, usr)
						userModelMap[usr.Id] = usr
						userIdList = append(userIdList, usr.Id)
					}
				}
			}
		}
		// search where user without role
		if uidList := user.GetUserIdListNotInUserIdList(user_role.GetUserIdList()); len(uidList) > 0 {
			for _, uid := range uidList {
				if uid == req.Id { // skip the record of myself
					continue
				}
				usr, err := user.GetModelById(uid)
				if err == nil {
					if uid == req.Id { // skip the record of myself
						continue
					}
					if !slices.Contains(userIdList, usr.Id) {
						//userModelList = append(userModelList, usr)
						userModelMap[usr.Id] = usr
						userIdList = append(userIdList, usr.Id)
					}
				}
			}
		}
	} else if req.Behavior == 2 {
		if len(req.Name) <= 0 && len(req.PhoneNumber) <= 0 && req.UserId <= 0 {
			rsp.Code = code.NoData
			return nil
		}
		if len(req.Name) > 0 {
			tempUserModelList, err := user.GetModelListByName(string(req.Name))
			if err == nil {
				for _, usr := range tempUserModelList {
					if usr.Id == req.UserId { // skip the record of myself
						continue
					}
					if !slices.Contains(userIdList, usr.Id) {
						if rbac.GetTopRankOfRoleList(user_role.GetRoleListByUserId(usr.Id)) <= rank {
							//userModelList = append(userModelList, usr)
							userModelMap[usr.Id] = usr
							userIdList = append(userIdList, usr.Id)
						}
					}
				}
			}
		}

		if len(req.PhoneNumber) > 0 {
			tempUserModelList, err := user.GetModelListByPhoneNumber(req.PhoneNumber)
			if err == nil {
				for _, usr := range tempUserModelList {
					if usr.Id == req.UserId { // skip the record of myself
						continue
					}
					if !slices.Contains(userIdList, usr.Id) {
						if rbac.GetTopRankOfRoleList(user_role.GetRoleListByUserId(usr.Id)) <= rank {
							//userModelList = append(userModelList, usr)
							userModelMap[usr.Id] = usr
							userIdList = append(userIdList, usr.Id)
						}
					}
				}
			}
		}

		if req.UserId > 0 {
			usr, err := user.GetModelById(req.UserId)
			if err == nil {
				if usr.Id != req.UserId && !slices.Contains(userIdList, usr.Id) { // skip the record of myself
					if rbac.GetTopRankOfRoleList(user_role.GetRoleListByUserId(usr.Id)) <= rank {
						//userModelList = append(userModelList, usr)
						userModelMap[usr.Id] = usr
						userIdList = append(userIdList, usr.Id)
					}

				}

			}
		}
	} else {
		rsp.Code = code.InvalidData
		return nil
	}

	tableList, fieldListSet, _ := rbac.GetFieldListOfRoleList(selfRoleList)
	if slices.Contains(tableList, user.GetTableName()) {
		index := 0
		for i, length := 0, len(tableList); i < length; i++ {
			if strings.Compare(user.GetTableName(), tableList[i]) == 0 {
				index = i
				break
			}
		}

		userModelList := []*user.Model{}
		sort.Slice(userIdList, func(i, j int) bool { return userIdList[i] < userIdList[j] })
		for _, id := range userIdList {
			if temp, exist := userModelMap[id]; exist {
				userModelList = append(userModelList, temp)
			}
		}
		bytes, err := json.Marshal(transform(fieldListSet[index], userModelList))
		if err != nil {
			rsp.Code = code.InternalError
			return nil
		}

		rsp.Code = code.Success
		rsp.Body = bytes
		return nil
	}

	rsp.Code = code.AccessDenied
	return nil
}

func testSlicesContains(userId int64) bool {
	userIdList := []int64{1, 2, 3, 4, 5}
	return slices.Contains(userIdList, userId)
}
