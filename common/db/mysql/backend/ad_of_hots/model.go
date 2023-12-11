package ad_of_hots

import "time"

type Model struct {
	Id                  int64     `gorm:"column:id;type:bigint not null;primary_key;comment:用户ID" json:"id,string"`
	AdvertisementIdList string    `gorm:"column:advertisement_id_list;type:varchar(200) not null;default:'[]';comment:广告ID列表;" json:"advertisement_id_list"`
	Version             int64     `gorm:"column:version;type:bigint not null;default:0;comment:版本号" json:"version,string"`
	CreatedAt           time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	Description         string    `gorm:"column:description;type:varchar(200) not null;default:'';comment:注意、描述、记事" json:"description"`
}
