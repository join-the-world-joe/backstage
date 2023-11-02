package admin

type FetchRoleListReq struct {
	RoleList []string `json:"role_list"`
}

type FetchRoleListRsp struct {
	Code     int      `json:"code"`
	RoleList []string `json:"role_list"`
}
