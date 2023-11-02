package ad_of_snack_products

import "time"

type Model struct {
	Id          int64     `gorm:"column:id;type:bigint not null;primary_key;comment:用户ID" json:"id,string"`
	ADIdList    string    `gorm:"column:ad_id_list;type:varchar(200) not null;default:'[]';comment:特价广告ID列表;" json:"ad_id_list"`
	CreatedAt   time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	Description string    `gorm:"column:description;type:varchar(200) not null;default:'';comment:注意、描述、记事" json:"description"`
}
