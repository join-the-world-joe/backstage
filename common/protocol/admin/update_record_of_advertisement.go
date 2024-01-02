package admin

type UpdateRecordOfAdvertisementReq struct {
	UserId        int64
	Id            int64    `json:"id"`
	Title         []byte   `json:"title"`
	ProductId     int64    `json:"product_id"`
	Name          []byte   `json:"name"`
	CoverImage    []byte   `json:"cover_image"`
	FirstImage    []byte   `json:"first_image"`
	SecondImage   []byte   `json:"second_image"`
	ThirdImage    []byte   `json:"third_image"`
	FourthImage   []byte   `json:"fourth_image"`
	FifthImage    []byte   `json:"fifth_image"`
	SellingPrice  int      `json:"selling_price"`
	PlaceOfOrigin []byte   `json:"place_of_origin"`
	SellingPoints [][]byte `json:"selling_points"`
	Status        int      `json:"status"`
	Stock         int      `json:"stock"`
}

type UpdateRecordOfAdvertisementRsp struct {
	Code int `json:"code"`
}
