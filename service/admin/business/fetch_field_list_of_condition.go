package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/user_role"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"backstage/global/rbac"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

type Field struct {
	Name        string `json:"name"`
	Table       string `json:"table"`
	Description string `json:"description"`
}

type FieldListOutput struct {
	FieldList []*Field `json:"field_list"`
	Length    int      `json:"length"`
}

func FetchFieldListOfCondition(ctx context.Context, req *admin.FetchFieldListOfConditionReq, rsp *admin.FetchFieldListOfConditionRsp) error {
	if req.Id <= 0 {
		log.Error("FetchFieldListOfCondition failure, req.Id <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	roleList := user_role.GetRoleListByUserId(req.Id)

	if len(roleList) <= 0 {
		log.Error("FetchFieldListOfCondition failure, len(roleList) <= 0")
		rsp.Code = code.InternalError
		return nil
	}

	// check if role_list has permission
	hasPermission := false
	for _, v := range roleList {
		if rbac.HasPermission(v, cast.ToInt(admin.FetchFieldListOfConditionReq_)) {
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
		roleList = []string{} // reset role list
		if len(req.Role) <= 0 {
			rsp.Code = code.NoData
			return nil
		}
		if len(req.Role) > 0 {
			roleList = append(roleList, req.Role)
		}
	} else {
		rsp.Code = code.InvalidData
		return nil
	}

	tableList, fieldList, descList := rbac.GetFieldListOfRoleList(roleList)
	fmt.Println("Table List: ", tableList)
	fmt.Println("Field List: ", fieldList)
	fmt.Println("Description List: ", descList)

	flo := &FieldListOutput{
		FieldList: []*Field{},
		Length:    0,
	}

	for k, tblName := range tableList {
		for index, _ := range fieldList[k] {
			flo.FieldList = append(flo.FieldList, &Field{
				Table:       "table." + tblName,
				Name:        tblName + "." + fieldList[k][index],
				Description: descList[k][index],
			})
			flo.Length += 1
		}
	}

	bytes, err := json.Marshal(flo)
	if err != nil {
		log.Error("FetchFieldListOfCondition.json.Marshal failure, err: ", err.Error())
		rsp.Code = code.InternalError
		return nil
	}

	fmt.Println("bytes: ", string(bytes))

	//js := json.New()
	//for i, length := 0, len(tableList); i < length; i++ {
	//	js.SetPath([]string{"table_list", tableList[i], "Field"}, fieldList[i])
	//	js.SetPath([]string{"table_list", tableList[i], "Description"}, descList[i])
	//}
	//
	//bytes, err := js.Encode()
	//if err != nil {
	//	log.Error("FetchFieldListOfCondition failure, err: ", err.Error())
	//	rsp.Code = code.InternalError
	//	return nil
	//}

	rsp.Code = code.Success
	rsp.Body = bytes
	return nil
}
