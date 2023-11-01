package user

import "time"

type Model struct {
	Id          int64  `gorm:"column:id;type:bigint not null;primary_key;comment:用户ID" json:"id,string"`
	Account     string `gorm:"column:account;type:varchar(60) not null;default:'';comment:帐号;" json:"account"`
	Email       string `gorm:"column:email;type:varchar(60) not null;default:'';comment:邮箱" json:"email"`
	Password    string `gorm:"column:password;type:varchar(250) not null;default:'';comment:密码" json:"password"`
	CountryCode string `gorm:"column:country_code;type:varchar(20) not null;index:idx_mobile,priority:1;default:'';comment:国家地区码, 如: 86" json:"country_code"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(30) not null;index:idx_mobile,priority:2;default:'';comment:电话号码" json:"phone_number"`
	Status      int    `gorm:"column:status;type:tinyint not null;default:0;comment:启/禁/停 用, 1-启用(active), 2-停用(inactive) 3-禁用(需要重置密码)" json:"status"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
