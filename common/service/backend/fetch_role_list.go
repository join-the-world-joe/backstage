package backend

type FetchRoleListReq struct {
	Role string `json:"role"`
}

type FetchRoleListRsp struct {
	Code     int      `json:"code"`
	RoleList []string `json:"role_list"`
}
