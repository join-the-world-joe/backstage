package ad_of_carousel

import "time"

type Model struct {
	Id            int64     `gorm:"column:id;type:bigint not null;primary_key;comment:用户ID" json:"id,string"`
	ImagePathList string    `gorm:"column:image_path_list;type:varchar(500) not null;default:'[]';comment:图片路径列表;" json:"image_path_list"`
	CreatedAt     time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	Description   string    `gorm:"column:description;type:varchar(200) not null;default:'';comment:注意、描述、记事" json:"description"`
}
