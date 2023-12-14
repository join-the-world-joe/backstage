package business

type Item struct {
	Title             string   `json:"title"`
	Stock             int      `json:"stock"`
	Status            int      `json:"status"`
	SellingPrice      int      `json:"selling_price"`
	ProductId         int64    `json:"product_id"`
	ProductName       string   `json:"product_name"`
	SellingPoints     [][]byte `json:"selling_points"`
	Image             string   `json:"image"`
	PlaceOfOrigin     string   `json:"place_of_origin"`
	AdvertisementId   int64    `json:"advertisement_id"`
	AdvertisementName string   `json:"advertisement_name"`
}
