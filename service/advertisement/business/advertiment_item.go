package business

type Item struct {
	Title             string   `json:"title"`
	Stock             int      `json:"stock"`
	Status            int      `json:"status"`
	OSSPath           string   `json:"oss_path"`
	SellingPrice      int      `json:"selling_price"`
	ProductId         int64    `json:"product_id"`
	ProductName       string   `json:"product_name"`
	SellingPoints     [][]byte `json:"selling_points"`
	CoverImage        string   `json:"cover_image"`
	FirstImage        string   `json:"first_image"`
	SecondImage       string   `json:"second_image"`
	ThirdImage        string   `json:"third_image"`
	FourthImage       string   `json:"fourth_image"`
	FifthImage        string   `json:"fifth_image"`
	PlaceOfOrigin     string   `json:"place_of_origin"`
	AdvertisementId   int64    `json:"advertisement_id"`
	AdvertisementName string   `json:"advertisement_name"`
}
