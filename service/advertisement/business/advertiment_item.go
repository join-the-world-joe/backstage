package business

type Item struct {
	Title             string   `json:"title"`
	Stock             int      `json:"stock"`
	SellingPrice      int      `json:"selling_price"`
	ProductId         int64    `json:"product_id"`
	ProductName       string   `json:"product_name"`
	SellingPoints     [][]byte `json:"selling_points"`
	Image             string   `json:"image"`
	Description       string   `json:"description"` // from ad_of_xxx.description
	PlaceOfOrigin     string   `json:"place_of_origin"`
	AdvertisementId   int64    `json:"advertisement_id"`
	AdvertisementName string   `json:"advertisement_name"`
}
