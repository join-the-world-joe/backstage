package product

import "encoding/json"

type FetchRecordsOfProductReq struct {
	UserId        int64
	ProductIdList []int64 `json:"product_id_list"`
}

type FetchRecordsOfProductRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
