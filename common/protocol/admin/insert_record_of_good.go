package admin

type InsertRecordOfGoodReq struct {
	UserId      int64
	Name        []byte `json:"name"`
	BuyingPrice int    `json:"buying_price"`
	Status      int    `json:"status"`
	Vendor      []byte `json:"vendor"`
	Contact     []byte ` json:"contact"`
	Description []byte `json:"description"`
}

type InsertRecordOfGoodRsp struct {
	Code int `json:"code"`
}
