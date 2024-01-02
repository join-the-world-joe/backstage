package admin

import "encoding/json"

type InsertRecordOfAdvertisementReq struct {
	UserId        int64
	Name          []byte   `json:"name"`
	Title         []byte   `json:"title"`
	SellingPrice  int      `json:"selling_price"`
	SellingPoints [][]byte `json:"selling_points"`
	PlaceOfOrigin []byte   `json:"place_of_origin"`
	CoverImage    []byte   `json:"cover_image"`
	FirstImage    []byte   `json:"first_image"`
	SecondImage   []byte   `json:"second_image"`
	ThirdImage    []byte   `json:"third_image"`
	FourthImage   []byte   `json:"fourth_image"`
	FifthImage    []byte   `json:"fifth_image"`
	Stock         int      `json:"stock"`
	Status        int      `json:"status"`
	ProductId     int64    `json:"product_id"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type InsertRecordOfAdvertisementRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
