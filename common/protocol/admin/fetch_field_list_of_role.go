package admin

import "encoding/json"

type FetchFieldListOfRoleReq struct {
	RoleList []string `json:"role_list"`
	Role     string   `json:"role"`
}

type FetchFieldListOfRoleRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
