package advertisement

type FetchVersionOfADOfHotDealsReq struct {
	UserId int64
}

type FetchVersionOfADOfHotDealsRsp struct {
	Code    int `json:"code"`
	Version int `json:"version"`
}
