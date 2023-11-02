package admin

import "encoding/json"

type FetchMenuListOfRoleReq struct {
	RoleList []string `json:"role_list"`
	Role     string   `json:"role"`
}

type FetchMenuListOfRoleRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
