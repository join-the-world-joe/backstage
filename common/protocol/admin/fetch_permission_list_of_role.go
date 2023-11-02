package admin

type FetchPermissionListOfRoleReq struct {
	RoleList []string `json:"role_list"`
	Role     string   `json:"role"`
}

type FetchPermissionListOfRoleRsp struct {
	Code           int   `json:"code"`
	PermissionList []int `json:"permission_list"`
}
