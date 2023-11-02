package admin

import "encoding/json"

type FetchMenuListReq struct {
	RoleList            []string `json:"role_list"`
	ConditionOfRoleList []string `json:"condition_of_role_list"`
}

type FetchMenuListRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
