package admin

type UpdateRecordOfGoodReq struct {
	UserId      int64
	ProductId   int64  `json:"product_id"`
	Name        []byte `json:"name"`
	BuyingPrice int    `json:"buying_price"`
	Vendor      []byte `json:"vendor"`
	Contact     []byte ` json:"contact"`
}

type UpdateRecordOfGoodRsp struct {
	Code int `json:"code"`
}
