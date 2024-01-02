package admin

type SoftDeleteRecordsOfProductReq struct {
	UserId        int64
	ProductIdList []int64 `json:"product_id_list"`
}

type SoftDeleteRecordsOfProductRsp struct {
	Code int `json:"code"`
}
