package admin

type InsertRecordOfProductReq struct {
	UserId      int64
	Name        []byte `json:"name"`
	BuyingPrice int    `json:"buying_price"`
	Vendor      []byte `json:"vendor"`
	Contact     []byte ` json:"contact"`
}

type InsertRecordOfProductRsp struct {
	Code int `json:"code"`
}
