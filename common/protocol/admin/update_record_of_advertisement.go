package admin

type UpdateRecordOfAdvertisementReq struct {
	UserId        int64
	Id            int64    `json:"id"`
	Title         []byte   `json:"title"`
	ProductId     int64    `json:"product_id"`
	Name          []byte   `json:"name"`
	SellingPrice  int      `json:"selling_price"`
	PlaceOfOrigin []byte   `json:"place_of_origin"`
	SellingPoints [][]byte `json:"selling_points"`
	Image         []byte   `json:"image"`
	Thumbnail     []byte   `json:"thumbnail"`
	Status        int      `json:"status"`
	Stock         int      `json:"stock"`
}

type UpdateRecordOfAdvertisementRsp struct {
	Code int `json:"code"`
}
