package admin

type InsertRecordOfGoodReq struct {
	UserId      int64
	Name        []byte `json:"name"`
	BuyingPrice int    `json:"buying_price"`
	Vendor      []byte `json:"vendor"`
	Contact     []byte ` json:"contact"`
}

type InsertRecordOfGoodRsp struct {
	Code int `json:"code"`
}
