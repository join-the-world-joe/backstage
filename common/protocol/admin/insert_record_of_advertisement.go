package admin

type InsertRecordOfAdvertisementReq struct {
	UserId        int64
	Name          []byte   `json:"name"`
	Title         []byte   `json:"title"`
	SellingPrice  int      `json:"selling_price"`
	SellingPoints [][]byte `json:"selling_points"`
	PlaceOfOrigin []byte   `json:"place_of_origin"`
	Image         []byte   `json:"Image"`
	Thumbnail     []byte   `json:"thumbnail"`
	Stock         int      `json:"stock"`
	ProductId     int64    `json:"product_id"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type InsertRecordOfAdvertisementRsp struct {
	Code int `json:"code"`
}
