package admin

type UpdateRecordOfAdvertisementReq struct {
	UserId        int64
	Title         []byte `json:"title"`
	ProductId     int64  `json:"product_id"`
	Name          []byte `json:"name"`
	SellingPrice  int    `json:"selling_price"`
	PlaceOfOrigin []byte `json:"place_of_origin"`
	SellingPoint  []byte `json:"selling_point"`
	Url           []byte `json:"url"`
	Status        int    `json:"status"`
	Stock         int    `json:"stock"`
}

type UpdateRecordOfAdvertisementRsp struct {
	Code int `json:"code"`
}
