package model

import "time"

type User struct {
	Id        uint64    `gorm:"type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;NOT NULL;comment:主键" json:"id"`
	Username  string    `gorm:"type:varchar(100);default:'';NOT NULL;comment:用户名" json:"username"`
	Nickname  string    `gorm:"type:varchar(100);default:'';NOT NULL;comment:昵称/姓名" json:"nickname"`
	Password  string    `gorm:"type:varchar(100);default:'';NOT NULL;comment:密码" json:"password"`
	Mobile    string    `gorm:"type:char(11);default:'';NOT NULL;comment:手机号" json:"mobile"`
	Email     string    `gorm:"type:varchar(100);default:'';NOT NULL;comment:邮箱" json:"email"`
	Status    int8      `gorm:"type:tinyint(4);default:1;NOT NULL;comment:状态 1.启用 2.禁用" json:"status"`
	CreatedAt time.Time `gorm:"type:datetime;default:NULL;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime;default:NULL;comment:更新时间" json:"updatedAt"`
	Avatar    string    `gorm:"type:varchar(500);default:'';NOT NULL;comment:头像" json:"avatar"`
}
