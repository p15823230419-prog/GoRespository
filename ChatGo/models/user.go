package models

import "time"

type User struct {
	Id        uint64    `gorm:"type:bigint(20);unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Username  string    `gorm:"type:varchar(20);uniqueIndex;comment:用户名;NOT NULL" json:"username" validate:"required"`
	Nickname  string    `gorm:"type:varchar(20);default:'';comment:昵称;NOT NULL" json:"nickname"`
	Password  string    `gorm:"type:varchar(20);comment:密码;NOT NULL" json:"password" validate:"required,min=6,max=20"`
	Avatar    string    `gorm:"type:varchar(255);default:'';comment:头像;NOT NULL" json:"avatar"`
	Email     string    `gorm:"type:varchar(20);default:'';comment:邮箱;NOT NULL" json:"email"`
	Phone     string    `gorm:"type:varchar(20);default:'';comment:手机号;NOT NULL" json:"phone"`
	CreatedAt time.Time `gorm:"type:datetime;comment:创建时间;column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime;comment:更新时间;column:updatedAt" json:"updatedAt"`
}
