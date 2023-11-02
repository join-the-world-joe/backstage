package advertisement

type FetchVersionOfADOfCarouselReq struct {
	UserId int64
}

type FetchVersionOfADOfCarouselRsp struct {
	Code    int `json:"code"`
	Version int `json:"version"`
}
