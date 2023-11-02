package category

import "time"

type Model struct {
	Id          int64  `gorm:"column:id;type:bigint not null;primary_key;comment:类目ID" json:"id,string"`
	Name        string `gorm:"column:name;type:varchar(50) not null;default:'';comment:类目名" json:"name"`
	Description string `gorm:"column:description;type:varchar(200) not null;default:'';comment:注意、描述、记事" json:"description"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
