package admin

type UpdateRecordOfGoodReq struct {
	UserId      int64
	ProductId   int64  `json:"product_id"`
	Name        []byte `json:"name"`
	BuyingPrice int    `json:"buying_price"`
	Status      int    `json:"status"`
	Vendor      []byte `json:"vendor"`
	Contact     []byte ` json:"contact"`
	Description []byte `json:"description"`
}

type UpdateRecordOfGoodRsp struct {
	Code int `json:"code"`
}
