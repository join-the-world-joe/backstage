package admin

type SoftDeleteRecordsOfGoodReq struct {
	UserId        int64
	ProductIdList []int64 `json:"product_id_list"`
}

type SoftDeleteRecordsOfGoodRsp struct {
	Code int `json:"code"`
}
