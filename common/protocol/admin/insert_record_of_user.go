package admin

type InsertRecordOfUserReq struct {
	Id          int64    `json:"id"`
	UserId      int64    `json:"user_id"`
	Name        []byte   `json:"name"`
	CountryCode string   `json:"country_code"`
	PhoneNumber string   `json:"phone_number"`
	Status      int      `json:"status"`
	Password    []byte   `json:"password"`
	RoleList    []string `json:"role_list"`
}

type InsertRecordOfUserRsp struct {
	Code int `json:"code"`
}
