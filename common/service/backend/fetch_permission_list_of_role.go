package backend

type FetchPermissionListOfRoleReq struct {
	Role  string `json:"role"`
	Query string `json:"query"`
}

type FetchPermissionListOfRoleRsp struct {
	Code           int   `json:"code"`
	PermissionList []int `json:"permission_list"`
}
