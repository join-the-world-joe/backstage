package backend

import "encoding/json"

type FetchMenuListOfRoleReq struct {
	Role  string `json:"role"`
	Query string `json:"query"`
}

type FetchMenuListOfRoleRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
