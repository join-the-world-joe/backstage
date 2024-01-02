package product

import "encoding/json"

type FetchIdListOfProductReq struct {
	UserId      int64
	Behavior    int    `json:"behavior"`
	ProductName []byte `json:"product_name"`
	CategoryId  int64  `json:"category_id"`
}

type FetchIdListOfProductRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
