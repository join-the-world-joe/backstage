package product

import "time"

type Model struct {
	Id          int64  `gorm:"column:id;type:bigint not null;primary_key;comment:用户ID" json:"id,string"`
	Name        string `gorm:"column:name;type:varchar(50) not null;default:'';comment:姓名" json:"name"`
	BuyingPrice int    `gorm:"column:buying_price;type:int(20) not null; default:0;comment:进货价格" json:"buying_price,string"`
	Visible     int    `gorm:"column:visible;type:tinyint(3) not null;default:1;comment:软删除标记 1-可见、0-不可见，已删除" json:"visible,string"`
	Status      int    `gorm:"column:status;type:tinyint(3) not null;default:1;comment:状态 启/停 用, 1-启用(active), 0-停用(inactive)" json:"status"`
	Vendor      string `gorm:"column:vendor;type:varchar(50) not null;default:'';comment:供应商" json:"vendor"`
	Contact     string `gorm:"column:contact;type:varchar(50) not null;default:'';comment:供应商联系方式" json:"contact"`
	Description string `gorm:"column:description;type:varchar(200) not null;default:'';comment:商品描述" json:"description"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
