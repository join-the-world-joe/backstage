package advertisement

import "time"

type Model struct {
	Id            int64  `gorm:"column:id;type:bigint not null;primary_key;comment:广告ID" json:"id,string"`
	Name          string `gorm:"column:name;type:varchar(80) not null;default:'';comment:广告名称" json:"name"`
	Title         string `gorm:"column:title;type:varchar(50) not null;default:'';comment:销售标题" json:"title"`
	Image         string `gorm:"column:image;type:varchar(500) not null;default:'';comment:广告图片信息" json:"image"` // key: 0-N, 0-thumbnail; 1-N, image; value: url of image
	Stock         int    `gorm:"column:stock;type:int(20) not null; default:0;comment:库存数量" json:"stock,string"`
	Status        int    `gorm:"column:status;type:tinyint(3) not null;default:1;comment:状态, 1-启用(active), 0-停用(inactive)" json:"status"`
	Visible       int    `gorm:"column:visible;type:tinyint(3) not null;default:1;comment:软删除标记 1-可见、0-不可见，已删除" json:"visible,string"`
	ProductId     int64  `gorm:"column:product_id;type:int(20) not null;comment:商品ID" json:"product_id,string"`
	SellingPrice  int    `gorm:"column:selling_price;type:int(20) not null; default:0;comment:销售价格" json:"selling_price,string"`
	PlaceOFOrigin string `gorm:"column:place_of_origin;type:varchar(50) not null;default:'';comment:产地" json:"place_of_origin"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
