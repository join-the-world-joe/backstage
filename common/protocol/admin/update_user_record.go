package admin

type UpdateUserRecordReq struct {
	Id          int64    `json:"id"`
	Name        []byte   `json:"name"`
	UserId      int64    `json:"user_id"`
	Status      int      `json:"status"`
	Password    []byte   `json:"password"`
	CountryCode string   `json:"country_code"`
	PhoneNumber string   `json:"phone_number"`
	RoleList    []string `json:"role_list"`
}

type UpdateUserRecordRsp struct {
	Code int `json:"code"`
}
