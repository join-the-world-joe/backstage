package business

type Item struct {
	Title         string `json:"title"`
	Stock         int    `json:"stock"`
	Price         int    `json:"price"`
	ProductId     int64  `json:"product_id"`
	ImagePath     string `json:"image_path"`
	Description   string `json:"description"`
	SellingPoint  string `json:"selling_point"`
	PlaceOfOrigin string `json:"place_of_origin"`
}
