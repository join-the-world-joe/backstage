package user

import (
	"time"
)

type Model struct {
	Id          int64  `gorm:"column:id;type:bigint not null;primary_key;comment:用户ID" json:"id,string"`
	Name        string `gorm:"column:name;type:varchar(50) not null;default:'';comment:姓名" json:"name"`
	Account     string `gorm:"column:account;type:varchar(60) not null;default:'';comment:帐号;" json:"account"`
	Email       string `gorm:"column:email;type:varchar(60) not null;default:'';comment:邮箱" json:"email"`
	Visible     int    `gorm:"column:visible;type:tinyint(3) not null;default:1;comment:软删除标记 1-可见、0-不可见，已删除" json:"visible,string"`
	Department  string `gorm:"column:department;type:varchar(30) not null;default:'';comment:部门" json:"department"`
	Password    string `gorm:"column:password;type:varchar(60) not null;default:'';comment:密码" json:"password"`
	CountryCode string `gorm:"column:country_code;type:varchar(10) not null;index:idx_mobile,priority:1;default:'';comment:国家地区码, 如: 86" json:"country_code"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20) not null;index:idx_mobile,priority:2;default:'';comment:电话号码" json:"phone_number"`
	Status      int    `gorm:"column:status;type:tinyint(3) not null;default:1;comment:状态 启/禁/停 用, 1-启用(active), 2-停用(inactive) 3-禁用(存在帐号问题)" json:"status"`
	MemberId    string `gorm:"column:member_id;type:varchar(40) not null;index:idx_member_id,unique;default:'';comment:会员编号" json:"member_id"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
