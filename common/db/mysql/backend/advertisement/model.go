package advertisement

import "time"

type Model struct {
	Id            int64  `gorm:"column:id;type:bigint not null;primary_key;comment:广告ID" json:"id,string"`
	Name          string `gorm:"column:name;type:varchar(80) not null;default:'';comment:广告名称" json:"name"`
	Title         string `gorm:"column:title;type:varchar(50) not null;default:'';comment:销售标题" json:"title"`
	CoverImage    string `gorm:"column:cover_image;type:varchar(200) not null;default:'';comment:封面图" json:"cover_image"`
	FirstImage    string `gorm:"column:first_image;type:varchar(200) not null;default:'';comment:广告首图" json:"first_image"`
	SecondImage   string `gorm:"column:second_image;type:varchar(200) not null;default:'';comment:第二张广告图" json:"second_image"`
	ThirdImage    string `gorm:"column:third_image;type:varchar(200) not null;default:'';comment:第三张广告图" json:"third_image"`
	FourthImage   string `gorm:"column:fourth_image;type:varchar(200) not null;default:'';comment:第四张广告图" json:"fourth_image"`
	FifthImage    string `gorm:"column:fifth_image;type:varchar(200) not null;default:'';comment:第五张广告图" json:"fifth_image"`
	Stock         int    `gorm:"column:stock;type:int(20) not null; default:0;comment:库存数量" json:"stock,string"`
	Status        int    `gorm:"column:status;type:tinyint(3) not null;default:1;comment:状态, 1-启用(active), 0-停用(inactive)" json:"status"`
	Visible       int    `gorm:"column:visible;type:tinyint(3) not null;default:1;comment:软删除标记 1-可见、0-不可见，已删除" json:"visible,string"`
	ProductId     int64  `gorm:"column:product_id;type:int(20) not null;comment:商品ID" json:"product_id,string"`
	SellingPrice  int    `gorm:"column:selling_price;type:int(20) not null; default:0;comment:销售价格" json:"selling_price,string"`
	PlaceOFOrigin string `gorm:"column:place_of_origin;type:varchar(50) not null;default:'';comment:产地" json:"place_of_origin"`
	OSSPath       string `gorm:"column:oss_path;type:varchar(100) not null;default:'';comment:Image的OSS路径" json:"oss_path"`
	OSSFolder     string `gorm:"column:oss_folder;type:varchar(50) not null;default:'';comment:Image的OSS目录" json:"oss_folder"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
