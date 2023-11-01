package backend

type FetchPermissionListReq struct {
	Role string `json:"role"`
}

type FetchPermissionListRsp struct {
	Code           int   `json:"code"`
	PermissionList []int `json:"permission_list"`
}
