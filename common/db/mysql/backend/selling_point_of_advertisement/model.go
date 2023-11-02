package selling_point_of_advertisement

import "time"

type Model struct {
	Id           int64  `gorm:"column:id;type:bigint not null;primary_key;comment:ID" json:"id,string"`
	ProductId    int64  `gorm:"column:product_id;type:bigint not null;default:0;comment:商品ID" json:"product_id,string"`
	SellingPoint string `gorm:"column:selling_point;type:varchar(30) not null;default:'';comment:卖点" json:"selling_point"`
	Visible      int    `gorm:"column:visible;type:tinyint(3) not null;default:1;comment:软删除标记 1-可见、0-不可见，已删除" json:"visible,string"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
