package user_role

import "time"

type Model struct {
	Id      int64  `gorm:"column:id;type:bigint not null;primary_key;comment:ID" json:"id,string"`
	UserId  int64  `gorm:"column:user_id;type:bigint not null;default:0;comment:用户ID" json:"user_id,string"`
	Role    string `gorm:"column:role;type:varchar(30) not null;default:'';comment:角色" json:"role"`
	Visible int    `gorm:"column:visible;type:tinyint(3) not null;default:1;comment:软删除标记 1-可见、0-不可见，已删除" json:"visible,string"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
