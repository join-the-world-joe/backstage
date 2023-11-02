package admin

type InsertRecordOfAdvertisementReq struct {
	UserId        int64
	Name          []byte `json:"name"`
	Title         []byte `json:"title"`
	SellingPrice  int    `json:"selling_price"`
	SellingPoint  string `json:"selling_point"`
	PlaceOfOrigin []byte `json:"place_of_origin"`
	Url           []byte `json:"url"`
	Stock         int    `json:"stock"`
	ProductId     int64  `json:"product_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Status        int    `json:"status"` // 0-inactive; 1-active
	Description   []byte `json:"description"`
}

type InsertRecordOfAdvertisementRsp struct {
	Code int `json:"code"`
}
