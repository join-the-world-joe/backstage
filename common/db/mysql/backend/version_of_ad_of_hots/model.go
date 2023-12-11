package version_of_ad_of_hots

import "time"

type Model struct {
	Id        int64     `gorm:"column:id;type:bigint not null;primary_key;comment:ID" json:"id,string"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}
