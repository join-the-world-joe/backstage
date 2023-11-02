package admin

type FetchUserListReq struct {
	RoleList    string `json:"role_list"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	PhoneNumber string `josn:"phone_number"`
}

type FetchUserListRsp struct {
	Code     int      `json:"code"`
	RoleList []string `json:"role_list"`
}
