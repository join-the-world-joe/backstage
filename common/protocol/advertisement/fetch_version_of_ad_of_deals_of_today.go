package advertisement

type FetchVersionOfADOfDealsOfTodayReq struct {
	UserId int64
}

type FetchVersionOfADOfDealsOfTodayRsp struct {
	Code    int `json:"code"`
	Version int `json:"version"`
}
