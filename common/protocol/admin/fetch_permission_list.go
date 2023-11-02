package admin

type FetchPermissionListReq struct {
	RoleList []string `json:"role_list"`
}

type FetchPermissionListRsp struct {
	Code           int   `json:"code"`
	PermissionList []int `json:"permission_list"`
}
